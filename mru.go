package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

type mru struct {
	entries  []string
	filePath string
}

func (m *mru) load() {
	f, err := openOrCreateMru(m.filePath)
	defer f.Close()
	if err != nil {
		// TODO
		lerror.Fatal(err)
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		// TODO
		lerror.Fatal(err)
	}
	m.entries = strings.Split(string(contents), "\n")
}

func (m *mru) touch(entry string) {
	newEntries := make([]string, 0, len(m.entries)+1)
	for _, e := range m.entries {
		if e != entry {
			newEntries = append(newEntries, e)
		}
	}
	newEntries = append(newEntries, entry)
}

func (m *mru) save(entry string) {
	f, err := os.Create(m.filePath)
	defer f.Close()
	if err != nil {
		// TODO
		lerror.Fatal(err)
	}
	for _, e := range m.entries {
		fmt.Fprintln(f, e)
	}
}

func getMru() *mru {
	mruPath := path.Join(notesLocation, "mru")
	mru := &mru{filePath: mruPath}
	mru.load()
	return mru
}
