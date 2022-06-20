package main

//import "flag"

//const (
//fzfSinkOut  = "out"
//fzfSinkEdit = "edit"
//)

//type FzfOpts struct {
//enabled  bool
//sinkType string
//}

//func addFzfOpts(fs *flag.FlagSet, opts *FzfOpts) {
//fs.BoolVar(&opts.enabled, "fzf", false, "Use fzf to select")
//fs.StringVar(&opts.sinkType, "fzf", "", "What to do with fzf output. Possible values (out | edit)")
//}

//func getListFlags(fs *flag.FlagSet, opts *ListOpts) {
//}

type ListOpts struct {
	NoteFilter
	Fzf bool
}
