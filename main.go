package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pkg/errors"
)

func defaultNotesLocation() string {
	return path.Join(os.Getenv("HOME"), "notes-debug")
}

func parseAndFormatNote(dir fs.FS, notePath string, in io.WriteCloser) error {

	//ldebug.Printf("Checking note : %s", notePath)
	note, err := func() (*Note, error) {
		f, err := dir.Open(notePath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		return ParseNote(f, NoteFilter{ExcludeBody: true})
	}()
	if err != nil {
		return errors.Errorf("Failed to parse %s (skipping): %s", notePath, err)
	}

	_, err = fmt.Fprintf(in, "%s|%s|%s\n", note.Title, strings.Join(note.Tags, ","), path.Base(notePath))
	if err != nil {
		// TODO
		lerror.Fatal(err)
	}
	return nil

}

func sinkFzf(in io.WriteCloser) {
	dir := os.DirFS(notesLocation)
	entries, err := fs.ReadDir(dir, "store")
	if err != nil {
		// TODO
		lerror.Fatalf("Failed to read dir entries: %s", err)
	}
	mru := getMru()
	mruLookup := make(map[string]bool)
	for _, e := range mru.entries {
		mruLookup[e] = true
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".note") {
			continue
		}
		if _, prs := mruLookup[entry.Name()]; prs {
			continue
		}
		notePath := path.Join("store", entry.Name())
		err = parseAndFormatNote(dir, notePath, in)
		if err != nil {
			lerror.Print(err)
		}
	}
	for _, entry := range mru.entries {
		notePath := path.Join("store", entry)
		parseAndFormatNote(dir, notePath, in)
		if err != nil {
			lerror.Print(err)
		}
	}
}

var notesLocation string

func edit(notePath string) {
	editCmd := "vim"
	if os.Getenv("EDITOR") != "" {
		editCmd = os.Getenv("EDITOR")
	}
	if os.Getenv("NOTES_EDITOR") != "" {
		editCmd = os.Getenv("NOTES_EDITOR")
	}
	fullPath := path.Join(notesLocation, "store", notePath)
	cmd := exec.Command(editCmd, fullPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	notesLocation = defaultNotesLocation()
	initLogging()
	fzfLister := &FzfLister{MultiSelect: false}
	notePaths, err := fzfLister.filterInput(sinkFzf)
	if err != nil {
		lerror.Fatal(err)
	}
	if len(notePaths) > 0 {
		fmt.Println(notePaths[0])
		parts := strings.Split(notePaths[0], "|")
		edit(parts[2])
	}
}
