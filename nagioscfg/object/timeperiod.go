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
	t.init("timeperiod")
	t.fill(properties, timeperiodRegularProperties, timeperiodListProperties, timeperiodEnumProperties)
	return t
}
