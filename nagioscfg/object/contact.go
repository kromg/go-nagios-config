package object

import ()

var contactRegularProperties = []string{
	"alias",
	"can_submit_commands",
	"contact_name",
	"host_notification_commands",
	"host_notification_period",
	"host_notifications_enabled",
	"name",
	"pager",
	"register",
	"retain_nonstatus_information",
	"retain_status_information",
	"service_notification_commands",
	"service_notification_period",
	"service_notifications_enabled",
}

var contactListProperties = []string{
	"addressx",
	"contactgroups",
	"email",
	"use",
}

var contactEnumProperties = map[string]container{
	"host_notification_options":    enumProperty{"d": 1, "u": 1, "r": 1, "f": 1, "s": 1, "n": 1},
	"service_notification_options": enumProperty{"w": 1, "u": 1, "c": 1, "r": 1, "f": 1, "s": 1, "n": 1},
}

func NewContact(properties map[string]string) Object {
	c := new(object)
	// Initialize structure
	c.init("contact", properties, contactRegularProperties, contactListProperties, contactEnumProperties)
	return c
}
