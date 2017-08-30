package object

import ()


var serviceescalationRegularProperties = []string{
    "contact_groups",
    "escalation_period",
    "first_notification",
    "host_name",
    "hostgroup_name",
    "last_notification",
    "name",
    "notification_interval",
    "register",
    "service_description",
}

var serviceescalationListProperties = []string{
    "contacts",
    "use",
}

var serviceescalationEnumProperties = map[string]container{
    "escalation_options": enumProperty{"w": 1, "u": 1, "c": 1, "r": 1},
}


func NewServiceescalation(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("serviceescalation", properties, serviceescalationRegularProperties, serviceescalationListProperties, serviceescalationEnumProperties)
	return s
}

