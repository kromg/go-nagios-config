package object

import ()

var commandRegularProperties = []string{
	"command_line",
	"command_name",
	"name",
	"register",
}

var commandListProperties = []string{
	"use",
}

var commandEnumProperties = map[string]container{}

func NewCommand(properties map[string]string) Object {
	c := new(object)
	// Initialize structure
	c.init("command", properties, commandRegularProperties, commandListProperties, commandEnumProperties)
	return c
}
