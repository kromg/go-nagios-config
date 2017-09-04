package object

import ()

const (
	HostgroupName             = "hostgroup_name"
	HostgroupHostgroupMembers = "hostgroup_members"
	HostgroupHostMembers      = "members"
)

var hostgroupRegularProperties = []string{
	"action_url",
	"alias",
	HostgroupName,
	"name",
	"notes",
	"notes_url",
	"register",
}

var hostgroupListProperties = []string{
	HostgroupHostgroupMembers,
	HostgroupHostMembers,
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
