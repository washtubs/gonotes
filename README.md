
Search notes showing just the title

    notes ls -format '{{ .Title }}'

Search notes showing just the title, on select, simply write to stdout

    notes ls -format '{{ .Title }}' -out

List all note titles

    notes ls -r -format '{{ .Title }}'

Search notes by content (preview will have line highlighted)

    notes ls -content -format '{{ .Line }}' -q hello -tags dog,cat

Edit the most recently edited note

    notes edit -recent
