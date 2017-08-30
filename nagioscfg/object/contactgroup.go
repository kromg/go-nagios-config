package object

import ()

var contactgroupRegularProperties = []string{
	"alias",
	"contactgroup_name",
	"name",
	"register",
}

var contactgroupListProperties = []string{
	"contactgroup_members",
	"members",
	"use",
}

var contactgroupEnumProperties = map[string]container{}

func NewContactgroup(properties map[string]string) Object {
	c := new(object)
	// Initialize structure
	c.init("contactgroup")
	c.fill(properties, contactgroupRegularProperties, contactgroupListProperties, contactgroupEnumProperties)
	return c
}
