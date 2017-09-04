package object

import ()

var hostescalationRegularProperties = []string{
	"contact_groups",
	"escalation_period",
	"first_notification",
	"host_name",
	"hostgroup_name",
	"last_notification",
	"name",
	"notification_interval",
	"register",
}

var hostescalationListProperties = []string{
	"contacts",
	"use",
}

var hostescalationEnumProperties = map[string]Set{
	"escalation_options": Set{"d": 1, "u": 1, "r": 1},
}

func NewHostescalation(properties map[string]string) Object {
	h := new(object)
	// Initialize structure
	h.init("hostescalation")
	h.fill(properties, hostescalationRegularProperties, hostescalationListProperties, hostescalationEnumProperties)
	return h
}
