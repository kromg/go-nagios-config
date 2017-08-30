package nagioscfg

import (
	"log"
)

// Parse() reads nagios.cfg and parses it, plus it parses all configuration files
// references by nagios.cfg, as Nagios would, building up the complete representation
// of the configuration. It may take a while.
func (n *nagioscfg) Parse() (err error) {

	// Parse main configuration file
	log.Println("Parsing main configuration file...")
	if err = n.parseMainFile(); err != nil {
		return err
	}

	// Now parse resource.cfg (if configured)
	log.Println("Parsing resource_file (if configured)...")
	if err = n.parseResourceFile(); err != nil {
		return err
	}

	// Now parse all other configuration files
	log.Println("Parsing object definition files...")
	if err = n.parseObjectFiles(); err != nil {
		return err
	}

	// No errors reported
	return
}
