// Package nagioscfg/object provides representations for all Nagios object types
// (contacts, hosts, services, timeperiods, and so on)
//
// TODO:
//   - contact objects do not parse addressX lines - implement this
package object

import (
	"regexp"
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
	init(string, map[string]string, []string, []string, map[string]container) (err error)
	Dump()
}

type enumProperty map[string]int

type container interface {
	contains(string) bool
}

type object struct {
	oType          string
	properties     map[string]string
	listProperties map[string][]string
	enumProperties map[string][]string
}

func (o *object) init(oType string,
	properties map[string]string,
	propertiesDef []string,
	listPropertiesDef []string,
	enumPropertiesDef map[string]container) (err error) {
	// Initialize struct
	o.oType = oType
	o.properties = make(map[string]string)
	o.listProperties = make(map[string][]string)
	o.enumProperties = make(map[string][]string)

	// Get regular properties from the map
	for _, p := range propertiesDef {
		if val, ok := properties[p]; ok {
			o.properties[p] = val
		}
	}

	// Get list properties from the map
	for _, p := range listPropertiesDef {
		if val, ok := properties[p]; ok {
			o.listProperties[p] = listSeparator.Split(val, -1)
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
				if validChoices.contains(v) {
					o.enumProperties[p] = append(o.enumProperties[p], v)
				}
			}
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
