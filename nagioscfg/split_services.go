package nagioscfg

import (
	"log"
)

// Loop on all services and, when a service is found with more than one host defined,
// then one service is created for eache single host, and the original is removed from
// the configuration, leaving no service defined on multiple hosts at once
func (n *nagioscfg) SplitServicesToSingleHosts() {

	for _, s := range n.objects["service"] {
		if hosts, ok := s.GetList("host_name"); ok {
			if len(hosts) > 1 {
				for _, host := range hosts {
					sc := s.Copy()
					sc.Dump()
					log.Printf("Host: %s\n", host)
				}
				log.Printf("Splitting: %v\n", hosts)
			}
		}
	}

}
