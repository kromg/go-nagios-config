package nagioscfg

import (
	"regexp"
)

// General patterns
var leadingSpace = regexp.MustCompile("^\\s+")
var comment = regexp.MustCompile("\\s*(#|;).*$")
var emptyLine = regexp.MustCompile("^\\s*$")
var equal = regexp.MustCompile("\\s*=\\s*")

// Object-definitions-specific patterns
var space = regexp.MustCompile("\\s+")                            // Separator for block definitions
var continuedLine = regexp.MustCompile("\\\\$")                   // End in an escape (\)
var definition = regexp.MustCompile("define\\s+([^[:space:]{]+)") // define SOMEWHAT {
var blockEnd = regexp.MustCompile("^\\s*}\\s*$")                  // a lonely "}" on a line
