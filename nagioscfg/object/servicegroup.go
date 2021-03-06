package object

import ()

var servicegroupRegularProperties = []string{
	"action_url",
	"alias",
	"name",
	"notes",
	"notes_url",
	"register",
	"servicegroup_name",
}

var servicegroupListProperties = []string{
	"members",
	"servicegroup_members",
	"use",
}

var servicegroupEnumProperties = map[string]Set{}

func NewServicegroup(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("servicegroup")
	s.fill(properties, servicegroupRegularProperties, servicegroupListProperties, servicegroupEnumProperties)
	return s
}
