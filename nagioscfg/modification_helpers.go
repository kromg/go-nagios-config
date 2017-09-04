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
func (n *nagioscfg) expandHostgroup(hg string) (hosts object.Set) {

	hosts = object.NewSet()

	for _, hgDef := range n.objects[object.HOSTGROUP] {
		if hgName, hasName := hgDef.GetProperty(object.HG_HOSTGROUP_NAME); !hasName || hgName != hg {
			continue // Next hostgroup, please!
		}

		// Get all host members
		if members, hasMembers := hgDef.GetList(object.HG_MEMBERS); hasMembers {
			for _, host := range members {
				hosts.Add(host)
			}
		}

		// Get all hostgroups members, and expand them (recursively)
		if hgMembers, hasMembers := hgDef.GetList(object.HG_HOSTGROUP_MEMBERS); hasMembers {
			for _, mhg := range hgMembers {
				memberHosts := n.expandHostgroup(mhg)
				for host, _ := range memberHosts {
					hosts.Add(host)
				}
			}
		}
	}

	return hosts
}
