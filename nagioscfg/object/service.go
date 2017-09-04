package object

import ()

const (
	SVC_HOST_NAME      = "host_name"
	SVC_HOSTGROUP_NAME = "hostgroup_name"
)

var serviceRegularProperties = []string{
	"action_url",
	"active_checks_enabled",
	"check_command",
	"check_freshness",
	"check_interval",
	"check_period",
	"display_name",
	"event_handler",
	"event_handler_enabled",
	"first_notification_delay",
	"flap_detection_enabled",
	"freshness_threshold",
	"high_flap_threshold",
	"icon_image",
	"icon_image_alt",
	"is_volatile",
	"low_flap_threshold",
	"max_check_attempts",
	"name",
	"notes",
	"notes_url",
	"notification_interval",
	"notification_period",
	"notifications_enabled",
	"obsess_over_service",
	"passive_checks_enabled",
	"process_perf_data",
	"register",
	"retain_nonstatus_information",
	"retain_status_information",
	"retry_interval",
	"service_description",
}

var serviceListProperties = []string{
	"contact_groups",
	"contacts",
	"servicegroups",
	"use",
	SVC_HOST_NAME,
	SVC_HOSTGROUP_NAME,
}

var serviceEnumProperties = map[string]Set{
	"stalking_options":       NewSet("o", "w", "u", "c"),
	"flap_detection_options": NewSet("o", "w", "c", "u"),
	"notification_options":   NewSet("w", "u", "c", "r", "f", "s"),
	"initial_state":          NewSet("o", "w", "u", "c"),
}

func NewService(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("service")
	s.fill(properties, serviceRegularProperties, serviceListProperties, serviceEnumProperties)
	return s
}
