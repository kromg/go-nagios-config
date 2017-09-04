package object

import (
	"bufio"
	"fmt"
	"strings"
)

func (o *object) Write(w *bufio.Writer) (err error) {

	// "Regular" properties
	keys := sortKeys(o.properties)
	for _, k := range keys {
		if _, err := w.WriteString(fmt.Sprintf("\t%s\t\t%s\n", k, o.properties[k])); err != nil {
			return err
		}
	}

	// "List" properties
	for k, v := range o.listProperties {
		if _, err := w.WriteString(fmt.Sprintf("\t%s\t\t%s\n", k, strings.Join(v, ","))); err != nil {
			return err
		}
	}

	// "Enum" properties
	for k, v := range o.enumProperties {
		if _, err := w.WriteString(fmt.Sprintf("\t%s\t\t%s\n", k, strings.Join(v, ","))); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return
}
