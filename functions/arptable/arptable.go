package arptable

import (
	"os/exec"
	"strings"
	"sync"
)

// ArpTable ArpTable
// type ArpTable map[string]string

// IPTable IPTable
func IPTable() *sync.Map {
	var table sync.Map
	data, err := exec.Command("arp", "-a").Output()
	if err != nil {
		return nil
	}
	// var table = make(ArpTable)
	skipNext := false
	for _, line := range strings.Split(string(data), "\n") {
		// skip empty lines
		if len(line) <= 0 {
			continue
		}
		// skip Interface: lines
		if line[0] != ' ' {
			skipNext = true
			continue
		}
		// skip column headers
		if skipNext {
			skipNext = false
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		ip := fields[0]
		// Normalize MAC address to colon-separated format
		table.Store(strings.Replace(strings.ToUpper(fields[1]), "-", "", -1), ip)
	}
	return &table
}
