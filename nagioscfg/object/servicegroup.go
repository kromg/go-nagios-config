package object

import ()


var servicegroupRegularProperties = []string{
    "action_url",
    "alias",
    "name",
    "notes",
    "notes_url",
    "register",
    "servicegroup_name",
}

var servicegroupListProperties = []string{
    "members",
    "servicegroup_members",
    "use",
}

var servicegroupEnumProperties = map[string]container{
}


func NewServicegroup(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("servicegroup", properties, servicegroupRegularProperties, servicegroupListProperties, servicegroupEnumProperties)
	return s
}

