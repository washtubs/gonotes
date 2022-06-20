package main

import (
	"io/fs"

	"github.com/pkg/errors"
)

type Store struct {
	fs.GlobFS
}

func (s *Store) listNotes(filter NoteFilter) ([]*Note, error) {
	noteFiles, err := fs.Glob(s, "store/*.note")
	notes := make([]*Note, 0, len(noteFiles))

	if err != nil {
		lerror.Fatal(errors.Wrapf(err, "Failed to glob for notes files"))
	}
	for _, noteFile := range noteFiles {
		f, err := s.Open(noteFile)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to open %s", noteFile)
		}
		note, err := ParseNote(f, filter)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed to parse %s", noteFile)
		}
		notes = append(notes, note)
	}
	return notes, nil

}
