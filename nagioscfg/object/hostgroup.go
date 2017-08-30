package object

import ()

var hostgroupRegularProperties = []string{
	"action_url",
	"alias",
	"hostgroup_name",
	"name",
	"notes",
	"notes_url",
	"register",
}

var hostgroupListProperties = []string{
	"hostgroup_members",
	"members",
	"use",
}

var hostgroupEnumProperties = map[string]container{}

func NewHostgroup(properties map[string]string) Object {
	h := new(object)
	// Initialize structure
	h.init("hostgroup")
	h.fill(properties, hostgroupRegularProperties, hostgroupListProperties, hostgroupEnumProperties)
	return h
}
