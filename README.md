# go-nagios-config
A Nagios config parser in Go (work in progress).

- The package is currently able to parse a valid nagios configuration and to print it

### Roadmap:
- be able to actually resolve group members and templating hierarchies;
- be able to validate configuration;
- parse timeperiod "execeptions"
- be able to write down configuration in a different format (for example: expand services defined on multiple hosts to single-host service definitions or - on the opposte - collapse multiple single-host definitions of a single service into a single, multi-host, definition);
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

    	// Finally print the read-in configuration.
    	cfg.Print()

    }