// Package nagioscfg/object provides representations for all Nagios object types
// (contacts, hosts, services, timeperiods, and so on)
//
// TODO:
//   - contact objects do not parse addressX lines - implement this
package object

import (
	"bufio"
	"regexp"
)

const (
	COMMAND           = "command"
	CONTACT           = "contact"
	CONTACTGROUP      = "contactgroup"
	HOST              = "host"
	HOSTDEPENDENCY    = "hostdependency"
	HOSTESCALATION    = "hostescalation"
	HOSTEXTINFO       = "hostextinfo"
	HOSTGROUP         = "hostgroup"
	SERVICE           = "service"
	SERVICEDEPENDENCY = "servicedependency"
	SERVICEESCALATION = "serviceescalation"
	SERVICEEXTINFO    = "serviceextinfo"
	SERVICEGROUP      = "servicegroup"
	TIMEPERIOD        = "timeperiod"
)

var Type = map[string]interface{}{
	"contact":           NewContact,
	"command":           NewCommand,
	"contactgroup":      NewContactgroup,
	"host":              NewHost,
	"hostdependency":    NewHostdependency,
	"hostescalation":    NewHostescalation,
	"hostextinfo":       NewHostextinfo,
	"hostgroup":         NewHostgroup,
	"service":           NewService,
	"servicedependency": NewServicedependency,
	"serviceescalation": NewServiceescalation,
	"serviceextinfo":    NewServiceextinfo,
	"servicegroup":      NewServicegroup,
	"timeperiod":        NewTimeperiod,
}

var listSeparator = regexp.MustCompile("\\s*,\\s*")

type Object interface {
	init(string)
	fill(map[string]string, []string, []string, map[string]Set) (err error)
	Dump()
	GetProperty(string) (string, bool)
	GetList(string) ([]string, bool)
	GetEnum(string) ([]string, bool)
	SetProperty(string, string)
	SetList(string, []string)
	SetEnum(string, []string)
	Copy() Object
	Write(*bufio.Writer) error
}

type enumProperty map[string]int

type object struct {
	oType          string
	properties     map[string]string
	listProperties map[string][]string
	enumProperties map[string][]string
}

func (o *object) init(oType string) {
	// Initialize struct
	o.oType = oType
	o.properties = make(map[string]string)
	o.listProperties = make(map[string][]string)
	o.enumProperties = make(map[string][]string)
}

func (o *object) fill(properties map[string]string,
	propertiesDef []string,
	listPropertiesDef []string,
	enumPropertiesDef map[string]Set) (err error) {

	// Get regular properties from the map
	for _, p := range propertiesDef {
		if val, ok := properties[p]; ok {
			o.properties[p] = val
			// Remove the property from the map
			delete(properties, p)
		}
	}

	// Get list properties from the map
	for _, p := range listPropertiesDef {
		if val, ok := properties[p]; ok {
			o.listProperties[p] = listSeparator.Split(val, -1)
			// Remove the property from the map
			delete(properties, p)
		}
	}

	// Get "enumeration" properties from the map
	for p, validChoices := range enumPropertiesDef {
		if val, ok := properties[p]; ok {
			// Init the slice
			o.enumProperties[p] = make([]string, 0)
			// For each value read from the configuration, check if is a valid choice
			for _, v := range listSeparator.Split(val, -1) {
				// TODO: return an error if value from configuration is not allowed
				if validChoices.Contains(v) {
					o.enumProperties[p] = append(o.enumProperties[p], v)
				}
			}
			// Remove the property from the map
			delete(properties, p)
		}
	}

	// Everything ok
	return
}

// Lookup in a map is faster than running through a list
func (e enumProperty) contains(k string) bool {
	_, ok := e[k]
	return ok
}

// Retrieve a regular property by name
func (o *object) GetProperty(p string) (string, bool) {
	v, ok := o.properties[p]
	return v, ok
}

// Retrieve a list property by name
func (o *object) GetList(p string) ([]string, bool) {
	v, ok := o.listProperties[p]
	return v, ok
}

// Retrieve an enum property by name
func (o *object) GetEnum(p string) ([]string, bool) {
	v, ok := o.enumProperties[p]
	return v, ok
}

// Retrieve a regular property by name
func (o *object) SetProperty(p string, v string) {
	o.properties[p] = v
}

// Retrieve a list property by name
func (o *object) SetList(p string, s []string) {
	o.listProperties[p] = s
}

// Retrieve an enum property by name
func (o *object) SetEnum(p string, s []string) {
	o.enumProperties[p] = s
}

// Get a copy of the object
func (o *object) Copy() Object {
	no := new(object)
	no.init(o.oType)
	// Copy the properties
	for k, v := range o.properties {
		no.properties[k] = v
	}

	// Copy the list properties
	for k, s := range o.listProperties {
		no.listProperties[k] = copySlice(s)
	}

	// Copy the enum properties
	for k, e := range o.enumProperties {
		no.enumProperties[k] = copySlice(e)
	}

	return no
}

func copySlice(s []string) []string {
	ns := make([]string, len(s), len(s))
	for i, v := range s {
		ns[i] = v
	}
	return ns
}
