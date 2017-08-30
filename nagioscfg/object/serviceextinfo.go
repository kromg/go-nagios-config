package object

import ()

var serviceextinfoRegularProperties = []string{
	"action_url",
	"host_name",
	"icon_image",
	"icon_image_alt",
	"name",
	"notes",
	"notes_url",
	"register",
	"service_description",
}

var serviceextinfoListProperties = []string{
	"use",
}

var serviceextinfoEnumProperties = map[string]container{}

func NewServiceextinfo(properties map[string]string) Object {
	s := new(object)
	// Initialize structure
	s.init("serviceextinfo", properties, serviceextinfoRegularProperties, serviceextinfoListProperties, serviceextinfoEnumProperties)
	return s
}
