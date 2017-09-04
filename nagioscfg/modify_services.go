package nagioscfg

import (
	"github.com/kromg/go-nagios-config/nagioscfg/object"
)

// Loop on all services and, when a service is found with more than one host defined,
// then one service is created for eache single host, and the original is removed from
// the configuration, leaving no service defined on multiple hosts at once
func (n *nagioscfg) ExpandHostsInServiceDefinitions() {

	for i, s := range n.objects[object.Service] {
		if hosts, ok := s.GetList(object.HostName); ok {
			if len(hosts) > 1 {
				// Remove the service from the list
				n.removeObject(object.Service, i)
				for _, host := range hosts {
					// Make a copy of the original
					sc := s.Copy()
					// Overwrite original hosts definition with our own
					sc.SetList(object.HostName, []string{host})
					// Append this definition to the services list
					n.objects[object.Service] = append(n.objects[object.Service], sc)
				}
			}
		}
	}
}

// Loop on all services and, when a service is found which is defined on a hostgroup,
// substitute that definition with N equivalent service definitions tied to one single host.
func (n *nagioscfg) ExpandHostgroupsInServiceDefinitions() {
	for i, s := range n.objects[object.Service] {
		if hostgroups, ok := s.GetList(object.Hostgroup); ok {
			if len(hostgroups) > 0 {
				// Remove the service from the list
				n.removeObject(object.Service, i)

				// Expand all hostgroups to hosts (recursively)
				hosts := make(map[string]int)
				for _, hg := range hostgroups {
					for host, _ := range n.expandHostgroup(hg) {
						hosts[host] = 1 // Using a map in place of a set
					}
				}

				// Define one service per host
				for host, _ := range hosts {
					// Make a copy of the original
					sc := s.Copy()
					// Overwrite original hosts definition with our own
					sc.SetList(object.HostName, []string{host})
					// Clear original hostgroup names
					sc.SetList(object.HostgroupName, []string{})
					// Append this definition to the services list
					n.objects[object.Service] = append(n.objects[object.Service], sc)
				}
			}
		}
	}
}
