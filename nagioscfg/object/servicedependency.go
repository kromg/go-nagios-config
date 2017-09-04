package object

import ()

var servicedependencyRegularProperties = []string{
	"dependency_period",
	"dependent_host_name",
	"dependent_hostgroup_name",
	"dependent_service_description",
	"dependent_servicegroup_name",
	"host_name",
	"hostgroup_name",
	"inherits_parent",
	"name",
	"register",
	"service_description",
	"servicegroup_name",
}

var servicedependencyListProperties = []string{
	"use",
}

var servicedependencyEnumProperties = map[string]Set{
	"execution_failure_criteria":    NewSet("o", "w", "u", "c", "p", "n"),
	"notification_failure_criteria": NewSet("o", "w", "u", "c", "p", "n"),
}

func NewServicedependency(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("servicedependency")
	s.fill(properties, servicedependencyRegularProperties, servicedependencyListProperties, servicedependencyEnumProperties)
	return s
}
