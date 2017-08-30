package object

import ()

var hostextinfoRegularProperties = []string{
	"2d_coords",
	"3d_coords",
	"action_url",
	"host_name",
	"icon_image",
	"icon_image_alt",
	"name",
	"notes",
	"notes_url",
	"register",
	"statusmap_image",
	"vrml_image",
}

var hostextinfoListProperties = []string{
	"use",
}

var hostextinfoEnumProperties = map[string]container{}

func NewHostextinfo(properties map[string]string) Object {
	h := new(object)
	// Initialize structure
	h.init("hostextinfo", properties, hostextinfoRegularProperties, hostextinfoListProperties, hostextinfoEnumProperties)
	return h
}
