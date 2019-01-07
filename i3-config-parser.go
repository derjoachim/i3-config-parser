package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	s "strings"
)

var PossiblePaths = [...]string{os.Getenv("HOME") + "/.config/i3/config", os.Getenv("HOME") + "/.i3/config", "/etc/xdg/i3/config", "/etc/i3/config"}

// TODO: Read from ~/.xmodmap file . I do not use one
var ModKeys = map[string]string{"Mod1": "Alt_L", "Mod2": "Num_Lock", "Mod3": "", "Mod4": "Win_L", "Mod5": "Mode_Switch"}

const DefaultModKey string = "Mod1"

type i3conf struct {
	path      string
	modifier  string
	variables map[string]string
	conf      map[string]string
}

// Add a configuration line to the i3conf struct
func (ic i3conf) AddToConfig(flds []string) {
	var key, value = "", ""
	if s.HasPrefix(flds[1], "--") {
		key = flds[2]
		value = s.Join(flds[3:], " ")
	} else {
		key = flds[1]
		value = s.Join(flds[2:], " ")
	}
	// TODO? Use a string builder?
	arK := s.Split(key, "+")
	for i := 0; i < cap(arK); i++ {
		if s.HasPrefix(arK[i], "$") {
			arK[i] = ic.variables[arK[i]]
		}
	}

	ic.conf[s.Join(arK, "+")] = value
}

// Read the found config file. make a map with a key-value pair of keypress
// and actions
func ParseConfig(f *os.File) i3conf {
	mod := ModKeys[DefaultModKey]

	v := make(map[string]string)
	m := make(map[string]string)

	ic := i3conf{path: f.Name(), modifier: mod, variables: v, conf: m}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		flds := s.Fields(scanner.Text())
		if len(flds) == 0 {
			continue
		}
		switch flds[0] {
		case "set":
			if s.Contains(line, "Mod") {
				mod = ModKeys[flds[2]]
				ic.modifier = mod
				ic.variables[flds[1]] = mod
			} else {
				ic.variables[flds[1]] = s.Join(flds[2:], " ")
			}
		case "bindsym":
			ic.conf[flds[1]] = s.Join(flds[2:], " ")
			ic.AddToConfig(flds)
		default:
			continue // Is this necessary?
		}
	}
	return ic
}

// Try to retrieve i3 config file from one of the possible locations. If it
// fails, error out
func FindConfigFile() *os.File {
	for i := 0; i < cap(PossiblePaths); i++ {
		if _, err := os.Stat(PossiblePaths[i]); os.IsNotExist(err) {
			continue
		}
		f, err := os.Open(PossiblePaths[i])
		if err != nil {
			log.Fatal(err)
		}

		return f
	}
	log.Fatal(errors.New("No config file was found at the usual locations."))
	return nil
}

func main() {
	f := FindConfigFile()
	defer f.Close()

	ic := ParseConfig(f)
	fmt.Println(ic)

	// Run the struct through a GTK3 thingumabob

	// profit
}
