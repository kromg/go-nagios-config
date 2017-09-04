package object

import (
	"fmt"
	"regexp"
)

var weekDayPattern = regexp.MustCompile("^\\s*((mon|tue|wednes|thurs|fri|satur|sun)day)\\s+(\\d+:\\d+.*)$")
var exceptionPattern = regexp.MustCompile("^\\s*(.*?)\\s+(\\d+:\\d+.*)$")
var spaceSeparator = regexp.MustCompile("\\s+")

var timeperiodRegularProperties = []string{
	"alias",
	"name",
	"register",
	"timeperiod_name",
}

var timeperiodListProperties = []string{
	"exclude",
	"friday",
	"monday",
	"saturday",
	"sunday",
	"thursday",
	"tuesday",
	"use",
	"wednesday",
}

var timeperiodEnumProperties = map[string]Set{}

func NewTimeperiod(properties map[string]string) Object {
	t := new(object)
	// Initialize structure
	t.init("timeperiod")

	// Timeperiods need a special treatment in order to parse "exceptions" also;
	// since there are "exceptions" which can confuse the normal parser, we have to
	// get them first
	for k, v := range properties {
		// Reassemble definition which was split on the first space
		def := fmt.Sprintf("%s %s", k, v)

		// See if it's a regular week-day definition
		wd := weekDayPattern.FindAllStringSubmatch(def, -1)
		if wd != nil {
			day := wd[0][1] // First capture group
			times := spaceSeparator.Split(wd[0][3], -1)
			t.listProperties[day] = times
			delete(properties, k)
			continue
		}

		// See if it's an exception
		xc := exceptionPattern.FindAllStringSubmatch(def, -1)
		if xc != nil {
			day := xc[0][1]
			times := spaceSeparator.Split(xc[0][2], -1)
			t.listProperties[day] = times
			delete(properties, k)
			continue
		}
	}

	// Now consider the remaining properties (if any)
	t.fill(properties, timeperiodRegularProperties, timeperiodListProperties, timeperiodEnumProperties)

	return t
}
