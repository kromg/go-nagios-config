package nagioscfg

import (
	"fmt"
	"sort"
)

// Prints the Nagios configurations as currently in memory in a friendly format.
func (n *nagioscfg) Print() {
	// Config file location, as passed in the initialization
	fmt.Printf("Location: %s\n", n.nagiosCfgLocation)

	// Dump of nagios.cfg directives
	fmt.Println("nagios.cfg:")
	keys := sortKeys(n.nagioscfg)
	for _, k := range keys {
		fmt.Printf("    %s => %s\n", k, n.nagioscfg[k])
	}

	// Dump of resource.cfg content
	if res, ok := n.nagioscfg[resourceFileDirective]; ok {
		fmt.Printf("%s: %s\n", resourceFileDirective, res)
		keys = sortKeys(n.resourcecfg)
		for _, k := range keys {
			fmt.Printf("    %s => %s\n", k, n.resourcecfg[k])
		}
	}

	// List of cfg_dir found in nagios.cfg
	fmt.Printf("%s:\n", cfgDirDirective)
	for _, v := range n.configDir {
		fmt.Printf("    %s\n", v)
	}

	// List of cfg_file found in nagios.cfg
	fmt.Printf("%s:\n", cfgFileDirective)
	for _, v := range n.configFile {
		fmt.Printf("    %s\n", v)
	}

	// List of broker_module found in nagios.cfg
	fmt.Printf("%s:\n", brokerModuleDirective)
	for _, v := range n.brokerModule {
		fmt.Printf("    %s\n", v)
	}

	// Configured objects
	fmt.Println("Objects:")
	for _, list := range n.objects {
		for _, o := range list {
			o.Dump()
		}
	}

}

// Sort the keys of maps in order to have a slightly better output
func sortKeys(m map[string]string) (mk []string) {
	mk = make([]string, 0, len(m))
	for k := range m {
		mk = append(mk, k)
	}
	sort.Strings(mk)
	return
}
