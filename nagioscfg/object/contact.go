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

var contactEnumProperties = map[string]Set{
	"service_notification_options": NewSet("w", "u", "c", "r", "f", "s", "n"),
	"host_notification_options":    NewSet("d", "u", "r", "f", "s", "n"),
}

func NewContact(properties map[string]string) Object {
	c := new(object)
	// Initialize structure
	c.init("contact")
	c.fill(properties, contactRegularProperties, contactListProperties, contactEnumProperties)
	return c
}
