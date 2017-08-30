package nagioscfg

import (
	"bufio"
	"os"
)

func (n *nagioscfg) parseMainFile() (err error) {
	// Open the file for reding
	f, err := os.Open(n.nagiosCfgLocation)
	defer f.Close()

	// Return the error if open went bad
	if err != nil {
		return err
	}

	// Scan the file, line by line, and read in basic configuration

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Split line, skip if empty
		key, val, isEmpty := splitLine(scanner.Text(), equal)
		if isEmpty {
			continue
		}

		switch key {
		case cfgFileDirective:
			n.configFile = append(n.configFile, val)
		case cfgDirDirective:
			n.configDir = append(n.configDir, val)
		case brokerModuleDirective:
			n.brokerModule = append(n.brokerModule, val)
		default:
			n.nagioscfg[key] = val
		}
	}

	// Return an error if scanner exited for an error
	if err = scanner.Err(); err != nil {
		return err
	}

	// Everything ok
	return
}

func (n *nagioscfg) parseResourceFile() (err error) {
	res, configured := n.nagioscfg[resourceFileDirective]

	// No problem if not in configuration
	if !configured {
		return
	}

	f, err := os.Open(res)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Split line, skip if empty
		key, val, isEmpty := splitLine(scanner.Text(), equal)
		if isEmpty {
			n.resourcecfg[key] = val
		}
	}

	// Return an error if scanner exited for an error
	if err = scanner.Err(); err != nil {
		return err
	}

	// No errors reported
	return
}
