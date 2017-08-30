package nagioscfg

import (
	"bufio"
	"github.com/kromg/go-nagios-config/nagioscfg/object"
	"os"
	"path/filepath"
	"regexp"
)

var cfgFilePattern = regexp.MustCompile("[^.\\\\]+\\.cfg$")

func (n *nagioscfg) parseObjectFiles() (err error) {
	// The easy part: the files
	for _, f := range n.configFile {
		if err = n.parseObjectFile(f); err != nil {
			return err
		}
	}

	// Now the more complex part: traversing the filesystem for all cfg_dir
	// Using a closure here to be able to call parseObjectFile() on the files
	for _, d := range n.configDir {
		err = filepath.Walk(d, func(path string, _ os.FileInfo, e error) error {
			if cfgFilePattern.MatchString(path) {
				e = n.parseObjectFile(path)
				if e != nil {
					return e
				}
			}
			return e
		})
		if err != nil {
			return err
		}
	}

	// Evrything ok
	return
}

func (n *nagioscfg) parseObjectFile(fileLocation string) (err error) {

	f, err := os.Open(fileLocation)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentObject := ""
	currentLine := ""
	blockProperties := make(map[string]string)
	for scanner.Scan() {
		// Trim comments, skip empty lines
		line, isEmpty := trim(scanner.Text())
		if isEmpty {
			continue
		}

		// Se if this line ends with an escape, in that case defer processing until property end
		if continuedLine.MatchString(line) {
			currentLine += continuedLine.ReplaceAllString(line, " ")
			continue
		} else {
			currentLine += line
		}

		// Does this line start a new block?
		objType := definition.FindAllStringSubmatch(currentLine, -1)
		if objType != nil {
			currentObject = objType[0][1] // First capture group of (only?) match
			currentLine = ""              // Ready for next line
			continue
		}

		// Does this line END a block?
		if blockEnd.MatchString(currentLine) {
			// Make a new object out of these properties (if type is known)
			if _, ok := object.Type[currentObject]; ok {
				n.objects[currentObject] = append(
					n.objects[currentObject],
					object.Type[currentObject].(func(map[string]string) object.Object)(blockProperties))
			}
			// Get ready for next block
			blockProperties = make(map[string]string)
			currentObject = ""
			currentLine = ""
		}

		// Otherwise, this line must define some block property

		key, value, isEmpty := splitLine(currentLine, space)
		if !isEmpty {
			blockProperties[key] = value
		}
		currentLine = "" // Ready for next line
	}

	// Return an error if scanner exited for an error
	if err = scanner.Err(); err != nil {
		return err
	}

	// Everything ok
	return
}
