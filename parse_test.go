package main

import (
	"strings"
	"testing"
)

func TestSimpleParse(t *testing.T) {
	simpleTest := `
This is a title

This is a body
This is a body
This is a body

------
- tag1
- tag2
`
	note, err := ParseNote(strings.NewReader(simpleTest), NoteFilter{})
	if err != nil {
		t.Fatal(err)
	}
	if len(note.Tags) != 2 {
		t.Errorf("Expected there to be 2 tags")
	}
	if len(strings.Split(note.Body, "\n"))-1 != 3 {
		t.Errorf("Expected 3 lines in the body, got %d", len(strings.Split(note.Body, "\n"))-1)
	}
	if note.Title != "This is a title" {
		t.Errorf("Expected title to exactly match")
	}

}

func TestNoBody(t *testing.T) {

	// Should *not* include
	noBody := `
Title

---
- zzz`
	note, err := ParseNote(strings.NewReader(noBody), NoteFilter{})
	if err != nil {
		t.Fatal(err)
	}
	if len(note.Tags) != 1 {
		t.Errorf("Expected there to be 1 tag")
	}
	if len(strings.Split(note.Body, "\n"))-1 != 0 {
		t.Errorf("Expected 0 lines in the body, got %d", len(strings.Split(note.Body, "\n"))-1)
	}
	if note.Title != "Title" {
		t.Errorf("Expected title to exactly match")
	}
}
