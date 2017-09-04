# go-nagios-config
A Nagios config parser in Go (work in progress).

The package is currently able to:
- parse a valid Nagios configuration;
- print the parsed configuration;
- write down the configuration;
- split services defined on multiple hosts (in a single definition block) into multiple definition blocks of services defined for a signle host (each);
- split services defined on one or more hostgroups into service definition blocks associated to a single host.

Before printing/writing configuration one can modify it slightly (the interface to the Nagios objects is still subject to changes).

### Roadmap:
- be able to actually resolve group members and templating hierarchies;
- be able to validate configuration;
- suggestions?


### Installation:

    # go get 'github.com/kromg/go-nagios-config/nagioscfg'



### Usage:

Something like:

    package main

    import (
    	"fmt"
    	"github.com/kromg/go-nagios-config/nagioscfg"
    	"log"
    )

    func main() {

    	// Init a configuration handler
    	cfg := nagioscfg.New("/etc/nagios/nagios.cfg")

    	// Read in all the configuration
    	err := cfg.Parse()

    	// Errors, anybody?
    	if err != nil {
    		log.Fatalf("Error parsing configuration: %v\n", err)
    	}

    	// Split the service definition blocks for services defined on many hosts
    	cfg.SplitServicesToSingleHosts()

    	// Split the service definition blocks for services defined on one or more hostgroups
		cfg.ExpandHostgroupsInServiceDefinitions()


    	// Finally print the read-in configuration.
    	cfg.Print()

    	// And maybe write it down somewhere
    	if err := cfg.WriteTo("/tmp/newNagiosCfg", true); err != nil {
		    log.Printf("Error: %v\n", err)
	    }

    }