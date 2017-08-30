package object

import ()


var hostdependencyRegularProperties = []string{
    "dependency_period",
    "dependent_host_name",
    "dependent_hostgroup_name",
    "host_name",
    "hostgroup_name",
    "inherits_parent",
    "name",
    "register",
}

var hostdependencyListProperties = []string{
    "use",
}

var hostdependencyEnumProperties = map[string]container{
    "execution_failure_criteria": enumProperty{"o": 1, "d": 1, "u": 1, "p": 1, "n": 1},
    "notification_failure_criteria": enumProperty{"o": 1, "d": 1, "u": 1, "p": 1, "n": 1},
}


func NewHostdependency(properties map[string]string) Object {
	h := new(object)
	// Initialize structure
	h.init("hostdependency", properties, hostdependencyRegularProperties, hostdependencyListProperties, hostdependencyEnumProperties)
	return h
}

