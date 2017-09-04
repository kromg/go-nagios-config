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

var commandEnumProperties = map[string]Set{}

func NewCommand(properties map[string]string) Object {
	c := new(object)
	// Initialize structure
	c.init("command")
	c.fill(properties, commandRegularProperties, commandListProperties, commandEnumProperties)
	return c
}
