package nagioscfg

import (
	"github.com/kromg/go-nagios-config/nagioscfg/object"
)

// Loop on all services and, when a service is found with more than one host defined,
// then one service is created for eache single host, and the original is removed from
// the configuration, leaving no service defined on multiple hosts at once
func (n *nagioscfg) ExpandHostsInServiceDefinitions() {

	for i, s := range n.objects[object.SERVICE] {
		if hosts, ok := s.GetList(object.SVC_HOST_NAME); ok {
			if len(hosts) > 1 {
				// Remove the service from the list
				n.removeObject(object.SERVICE, i)
				for _, host := range hosts {
					// Make a copy of the original
					sc := s.Copy()
					// Overwrite original hosts definition with our own
					sc.SetList(object.SVC_HOST_NAME, []string{host})
					// Append this definition to the services list
					n.objects[object.SERVICE] = append(n.objects[object.SERVICE], sc)
				}
			}
		}
	}
}

// Loop on all services and, when a service is found which is defined on a hostgroup,
// substitute that definition with N equivalent service definitions tied to one single host.
func (n *nagioscfg) ExpandHostgroupsInServiceDefinitions() {
	for i, s := range n.objects[object.SERVICE] {
		if hostgroups, ok := s.GetList(object.HOSTGROUP); ok {
			if len(hostgroups) > 0 {
				// Remove the service from the list
				n.removeObject(object.SERVICE, i)

				// Expand all hostgroups to hosts (recursively)
				hosts := object.NewSet()
				for _, hg := range hostgroups {
					hosts.AddAll(n.expandHostgroup(hg))
				}

				// Define one service per host
				for host, _ := range hosts {
					// Make a copy of the original
					sc := s.Copy()
					// Overwrite original hosts definition with our own
					sc.SetList(object.SVC_HOST_NAME, []string{host.(string)})
					// Clear original hostgroup names
					sc.SetList(object.SVC_HOSTGROUP_NAME, []string{})
					// Append this definition to the services list
					n.objects[object.SERVICE] = append(n.objects[object.SERVICE], sc)
				}
			}
		}
	}
}
