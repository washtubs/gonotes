package main

type Note struct {
	Title string
	Body  string
	Tags  []string
}

type NoteFilter struct {
	ExcludeBody bool
}
