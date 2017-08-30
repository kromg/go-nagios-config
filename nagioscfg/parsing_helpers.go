package nagioscfg

import (
	"regexp"
)

// Remove leading spaces and comments.
func trim(line string) (string, bool) {
	// Trim leading spaces
	line = leadingSpace.ReplaceAllString(line, "")
	// Remove comments from the line
	line = comment.ReplaceAllString(line, "")

	return line, emptyLine.MatchString(line)
}

// Analyze a line and, if it contains useful values, split it and return key/value pair to the caller.
func splitLine(line string, separator *regexp.Regexp) (key, value string, isEmpty bool) {
	// Trim, and see if empty
	line, isEmpty = trim(line)
	if isEmpty {
		return
	}
	// Split on equal sign
	keyAndVal := separator.Split(line, 2)
	return keyAndVal[0], keyAndVal[1], false
}
