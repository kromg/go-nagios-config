// Package nagioscfg provides a parser for nagios.cfg able to represent Nagios
// configuration completely, and allows user to retrieve/modify objects and
// values from that representation and to dump a new (optionally re-formatted)
// copy of the configuration.
package nagioscfg

import (
	"nagioscfg/object"
)

// Interface NagiosCfg contains all the methods to handle Nagios configuration.
type NagiosCfg interface {
	Parse()
	Print()
}

var resourceFileDirective = "resource_file"
var cfgDirDirective = "cfg_dir"
var cfgFileDirective = "cfg_file"
var brokerModuleDirective = "broker_module"

type nagioscfg struct {
	// Main config file location
	nagiosCfgLocation string
	// Main config file, when parsed
	nagioscfg map[string]string
	// Things in main config file that may be specified more than once
	configFile   []string
	configDir    []string
	brokerModule []string
	// resource.cfg when parsed
	resourcecfg map[string]string
	objects     map[string][]object.Object
}

// New() initializes a new Nagios configuration container (NagiosCfg interface)
func New(nagiosCfgLocation string) *nagioscfg {
	cfg := new(nagioscfg)
	cfg.nagiosCfgLocation = nagiosCfgLocation
	cfg.nagioscfg = make(map[string]string)
	cfg.configFile = make([]string, 0)
	cfg.configDir = make([]string, 0)
	cfg.brokerModule = make([]string, 0)
	cfg.resourcecfg = make(map[string]string)
	// Prepare containers for Nagios objects
	cfg.objects = make(map[string][]object.Object)
	for ot, _ := range object.Type {
		cfg.objects[ot] = make([]object.Object, 0)
	}
	return cfg
}
