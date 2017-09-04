package object

import ()

const (
	HostName = "host_name"
)

var hostRegularProperties = []string{
	"2d_coords",
	"3d_coords",
	"action_url",
	"active_checks_enabled",
	"address",
	"alias",
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
	"host_name",
	"icon_image",
	"icon_image_alt",
	"low_flap_threshold",
	"max_check_attempts",
	"name",
	"notes",
	"notes_url",
	"notification_interval",
	"notification_period",
	"notifications_enabled",
	"obsess_over_host",
	"passive_checks_enabled",
	"process_perf_data",
	"register",
	"retain_nonstatus_information",
	"retain_status_information",
	"retry_interval",
	"statusmap_image",
	"vrml_image",
}

var hostListProperties = []string{
	"contact_groups",
	"contacts",
	"hostgroups",
	"parents",
	"use",
}

var hostEnumProperties = map[string]Set{
	"flap_detection_options": NewSet("o", "d", "u"),
	"initial_state":          NewSet("o", "d", "u"),
	"notification_options":   NewSet("d", "u", "r", "f", "s"),
	"stalking_options":       NewSet("o", "d", "u"),
}

func NewHost(properties map[string]string) Object {
	h := new(object)
	// Initialize structure
	h.init("host")
	h.fill(properties, hostRegularProperties, hostListProperties, hostEnumProperties)
	return h
}
