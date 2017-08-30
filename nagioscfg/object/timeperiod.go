package object

import ()

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

var timeperiodEnumProperties = map[string]container{}

func NewTimeperiod(properties map[string]string) Object {
	t := new(object)
	// Initialize structure
	t.init("timeperiod", properties, timeperiodRegularProperties, timeperiodListProperties, timeperiodEnumProperties)

	// Timeperiod needs a special treatment for the so-called "exceptions".
	// init() removes the processed items from the properties map, so we can work on what is left

	// TODO: implement a different parsing for timeperiods and handle exceptions here (how?)

	return t
}
