package main

import (
	"flag"
	"fmt"
	"log"
	"nagioscfg"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	// Init a configuration handler
	cfg := nagioscfg.New("/tmp/nagios/nagios.cfg")
	// Read in all the configuration
	err := cfg.Parse()

	// Errors, anybody?
	if err != nil {
		log.Fatalf("Error parsing configuration: %v\n", err)
	}

	// Finally print the read-in configuration.
	cfg.Print()

}
