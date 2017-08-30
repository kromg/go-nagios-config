package object

import ()

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
	"host_name",
	"hostgroup_name",
	"contact_groups",
	"contacts",
	"servicegroups",
	"use",
}

var serviceEnumProperties = map[string]container{
	"initial_state":          enumProperty{"o": 1, "w": 1, "u": 1, "c": 1},
	"flap_detection_options": enumProperty{"o": 1, "w": 1, "c": 1, "u": 1},
	"stalking_options":       enumProperty{"o": 1, "w": 1, "u": 1, "c": 1},
	"notification_options":   enumProperty{"w": 1, "u": 1, "c": 1, "r": 1, "f": 1, "s": 1},
}

func NewService(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("service", properties, serviceRegularProperties, serviceListProperties, serviceEnumProperties)
	return s
}
