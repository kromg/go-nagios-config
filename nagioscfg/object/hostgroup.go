package object

import ()

const (
	HG_HOSTGROUP_NAME    = "hostgroup_name"
	HG_HOSTGROUP_MEMBERS = "hostgroup_members"
	HG_MEMBERS           = "members"
)

var hostgroupRegularProperties = []string{
	"action_url",
	"alias",
	HG_HOSTGROUP_NAME,
	"name",
	"notes",
	"notes_url",
	"register",
}

var hostgroupListProperties = []string{
	HG_HOSTGROUP_MEMBERS,
	HG_MEMBERS,
	"use",
}

var hostgroupEnumProperties = map[string]Set{}

func NewHostgroup(properties map[string]string) Object {
	h := new(object)
	// Initialize structure
	h.init("hostgroup")
	h.fill(properties, hostgroupRegularProperties, hostgroupListProperties, hostgroupEnumProperties)
	return h
}
