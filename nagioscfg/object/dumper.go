package object

import (
	"fmt"
	"sort"
)

func (o *object) Dump() {
	fmt.Printf("   %s\n", o.oType)

	// "Regular" properties
	keys := sortKeys(o.properties)
	for _, k := range keys {
		fmt.Printf("      %s => %s\n", k, o.properties[k])
	}

	// "List" properties
	for k, v := range o.listProperties {
		fmt.Printf("      %s => %v\n", k, v)
	}

	// "Enum" properties
	for k, v := range o.enumProperties {
		fmt.Printf("      %s => %v\n", k, v)
	}

}

// TODO: have a helper acting upon a "sortable" interface and have
// properties types implement sortable
// Sort the keys of maps in order to have a slightly better output
func sortKeys(m map[string]string) (mk []string) {
	mk = make([]string, 0, len(m))
	for k := range m {
		mk = append(mk, k)
	}
	sort.Strings(mk)
	return
}
