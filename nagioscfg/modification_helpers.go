package nagioscfg

import (
	"github.com/kromg/go-nagios-config/nagioscfg/object"
)

// Remove from the configuration the object of type objType at index i.
func (n *nagioscfg) removeObject(objType string, index int) {
	// Remove the object from the list
	n.objects[objType] = append(n.objects[objType][:index], n.objects[objType][index+1:]...)
}

// Return a set of the hosts included in a given hostgroup
func (n *nagioscfg) expandHostgroup(hg string) (hosts map[string]int) {

	hosts = make(map[string]int)

	for _, hgDef := range n.objects[object.Hostgroup] {
		if hgName, hasName := hgDef.GetProperty(object.HostgroupName); !hasName || hgName != hg {
			continue // Next hostgroup, please!
		}

		// Get all host members
		if members, hasMembers := hgDef.GetList(object.HostgroupHostMembers); hasMembers {
			for _, host := range members {
				hosts[host] = 1 // Using map as a set
			}
		}

		// Get all hostgroups members, and expand them (recursively)
		if hgMembers, hasMembers := hgDef.GetList(object.HostgroupHostMembers); hasMembers {
			for _, mhg := range hgMembers {
				memberHosts := n.expandHostgroup(mhg)
				for host, _ := range memberHosts {
					hosts[host] = 1
				}
			}
		}
	}

	return hosts
}
