package main

import (
	"bufio"
	"errors"
	"io"
	"regexp"
)

var emptyLineRegex *regexp.Regexp = regexp.MustCompile(`^\s*$`)
var separatorRegex *regexp.Regexp = regexp.MustCompile(`^\s*[=-]+\s*$`)

// At least 3 - or =
var tagSeparatorRegex *regexp.Regexp = regexp.MustCompile(`^\s*[-=]{3}[-=]*\s*$`)
var tagRegex *regexp.Regexp = regexp.MustCompile(`^-\s+(?P<Tag>[a-zA-Z0-9]+[a-zA-Z0-9]+)$`)

const (
	parseExpectTitle = iota
	parseExpectBody
	parseExpectTags
)

func ParseNote(r io.Reader, filter NoteFilter) (*Note, error) {
	note := Note{}
	note.Tags = make([]string, 0)
	scanner := bufio.NewScanner(r)

	// Read all lines into memory
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		return nil, errors.New("Empty note")
	}

	state := parseExpectTags
	bodyEndIndex := 0
	// Count back from the end to get the tags and bodyEndIndex
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if state == parseExpectTags {
			if emptyLineRegex.MatchString(line) {
				continue
			}
			if tagSeparatorRegex.MatchString(line) {
				state = parseExpectBody
				continue
			}
			subMatches := tagRegex.FindStringSubmatch(line)
			if subMatches != nil {
				note.Tags = append(note.Tags, subMatches[tagRegex.SubexpIndex("Tag")])
			}
		} else { // parseExpectBody
			if emptyLineRegex.MatchString(line) {
				continue
			} else {
				bodyEndIndex = i + 1
				break
			}
		}
	}
	if len(note.Tags) == 0 {
		bodyEndIndex = len(lines)
	} else if bodyEndIndex == 0 {
		return nil, errors.New("No tag separator found")
	}

	state = parseExpectTitle
	bodyLineCount := 0
out:
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		switch state {
		case parseExpectTitle:
			if emptyLineRegex.MatchString(line) {
				continue
			}
			if separatorRegex.MatchString(line) {
				continue
			}
			note.Title = line
			state++
		case parseExpectBody:
			if i == bodyEndIndex {
				break out
			}
			if bodyLineCount == 0 {
				if emptyLineRegex.MatchString(line) {
					continue
				}
				if separatorRegex.MatchString(line) {
					continue
				}
			}
			if !filter.ExcludeBody {
				note.Body = note.Body + line + "\n"
			}
		}
	}

	return &note, nil
}
