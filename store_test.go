package main

import (
	"testing"
	"testing/fstest"
)

func setupSimpleNoteStore() *Store {
	mapFs := make(fstest.MapFS)
	// Should include
	dogNote := `
Title

---
- pet
- dog
	`
	// Should *not* include
	aTxt := `
Title

---
- zzz
	`

	// Should include
	catNote := `
Title

---
- pet
- cat
	`

	// Should include
	carNote := `
Title

---
- car
- auto
	`
	mapFs["store/dog.note"] = &fstest.MapFile{Data: []byte(dogNote)}
	mapFs["store/a.txt"] = &fstest.MapFile{Data: []byte(aTxt)}
	mapFs["store/cat.note"] = &fstest.MapFile{Data: []byte(catNote)}
	mapFs["store/car.note"] = &fstest.MapFile{Data: []byte(carNote)}
	return &Store{mapFs}
}

func TestSimpleStore(t *testing.T) {
	s := setupSimpleNoteStore()
	notes, err := s.listNotes(NoteFilter{})
	if err != nil {
		t.Fatal(err)
	}
	if len(notes) != 3 {
		t.Errorf("Expected 3 notes, got %d", len(notes))
	}
}

func TestExcludeBody(t *testing.T) {
	s := setupSimpleNoteStore()
	notes, err := s.listNotes(NoteFilter{ExcludeBody: true})
	if err != nil {
		t.Fatal(err)
	}

	if notes[0].Body != "" {
		t.Error("Expected empty body")
	}
}
