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

var servicedependencyEnumProperties = map[string]container{
	"execution_failure_criteria":    enumProperty{"o": 1, "w": 1, "u": 1, "c": 1, "p": 1, "n": 1},
	"notification_failure_criteria": enumProperty{"o": 1, "w": 1, "u": 1, "c": 1, "p": 1, "n": 1},
}

func NewServicedependency(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("servicedependency")
	s.fill(properties, servicedependencyRegularProperties, servicedependencyListProperties, servicedependencyEnumProperties)
	return s
}
