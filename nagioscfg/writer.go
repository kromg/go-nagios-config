package nagioscfg

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Write down a previously read Nagios configuration
func (n *nagioscfg) WriteTo(destDir string, createObjectsSubdir bool) (err error) {
	if err := os.Mkdir(destDir, 0755); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	// Create the subdir also?
	var objDestDir string
	if createObjectsSubdir {
		objDestDir = path.Join(destDir, "objects")
		if err := os.Mkdir(objDestDir, 0755); err != nil {
			if !os.IsExist(err) {
				return err
			}
		}
	} else {
		objDestDir = destDir
	}

	// Write down nagios.cfg
	if err := n.writeCfgFile(destDir, "nagios.cfg", n.nagioscfg); err != nil {
		return err
	}

	// Write brokers to nagios.cfg
	if len(n.brokerModule) > 0 {
		appendProperties(destDir, "nagios.cfg", brokerModuleDirective, n.brokerModule)
	}

	// Write down resource.cfg, if defined
	if resFile, isPresent := n.nagioscfg[resourceFileDirective]; isPresent {
		if err := n.writeCfgFile(destDir, filepath.Base(resFile), n.resourcecfg); err != nil {
			return err
		}
	}

	// Write down object definitions
	written, err := n.writeObjectFiles(objDestDir)
	if err != nil {
		return err
	}

	// Complete Nagios config with objects files definitions
	if createObjectsSubdir {
		appendProperties(destDir, "nagios.cfg", cfgDirDirective, []string{objDestDir})
	} else {
		appendProperties(destDir, "nagios.cfg", cfgFileDirective, written)
	}

	return

}

func (n *nagioscfg) writeCfgFile(destDir string, fileName string, originalConfig map[string]string) (err error) {
	src := filepath.Dir(n.nagiosCfgLocation)

	// Replace any occurrence of the original file path with the new destination
	fileContent := make(map[string]string)
	for k, v := range originalConfig {
		fileContent[k] = strings.Replace(v, src, destDir, 1)
	}

	filePath := path.Join(destDir, fileName)

	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	keys := sortKeys(fileContent)
	for _, k := range keys {
		if _, err := w.WriteString(fmt.Sprintf("%s = %s\n", k, fileContent[k])); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return
}

func appendProperties(destDir string, fileName string, property string, values []string) (err error) {
	filePath := path.Join(destDir, fileName)

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	for _, v := range values {
		if _, err := w.WriteString(fmt.Sprintf("%s = %s\n", property, v)); err != nil {
			return err
		}
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func (n *nagioscfg) writeObjectFiles(objDestDir string) (written []string, err error) {

	written = make([]string, 0)

	// Loop over object types
	for objType, objects := range n.objects {

		// Create a new file with the name of the type
		filePath := path.Join(objDestDir, fmt.Sprintf("%s.cfg", objType))
		f, err := os.Create(filePath)
		defer f.Close()
		if err != nil {
			return written, err
		}

		// Write to the file
		w := bufio.NewWriter(f)
		for _, o := range objects {
			// Open the block
			if _, err := w.WriteString(fmt.Sprintf("define %s {\n", objType)); err != nil {
				return written, err
			}

			// Write all the definition directives:
			o.Write(w)

			// Close the block
			if _, err := w.WriteString("}\n\n"); err != nil {
				return written, err
			}
		}

		if err := w.Flush(); err != nil {
			return written, err
		}
		written = append(written, filePath)
	}

	return written, nil
}
