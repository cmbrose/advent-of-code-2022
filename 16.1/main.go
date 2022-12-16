package main

import (
	"encoding/hex"
	"fmt"
	"main/util"
	"regexp"
	"strings"

	"github.com/dropbox/godropbox/container/bitvector"
)

type valve struct {
	name      string
	flowRate  int
	neighbors []*valve

	isClosed func() bool
	close    func()
	open     func()
}

func main() {
	re := regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

	valves := make(map[string]*valve)
	closedStates := bitvector.NewBitVector(nil, 0)

	for _, line := range util.ReadInputLines("input.txt") {
		match := re.FindStringSubmatch(line)

		var (
			name         = match[1]
			rate         = util.AssertInt(match[2])
			neighborsStr = match[3]
			neighborsArr = strings.Split(neighborsStr, ", ")
		)

		id := closedStates.Length()

		// If 0 rate, just treat is as open
		if rate == 0 {
			closedStates.Append(0)
		} else {
			closedStates.Append(1)
		}

		v := &valve{
			name:      name,
			flowRate:  rate,
			neighbors: nil,

			isClosed: func() bool {
				return closedStates.Element(id) == 1
			},
			close: func() {
				closedStates.Set(1, id)
			},
			open: func() {
				closedStates.Set(0, id)
			},
		}

		valves[name] = v

		for _, name := range neighborsArr {
			if n, ok := valves[name]; ok {
				n.neighbors = append(n.neighbors, v)
				v.neighbors = append(v.neighbors, n)
			}
		}
	}

	maxPressureByState := make(map[string]int)

	getStateKey := func(node string, minutes int) string {
		return fmt.Sprintf("%s_%s_%d", node, hex.EncodeToString(closedStates.Bytes()), minutes)
	}

	start := valves["AA"]
	solution := getMaxPressureRelief(start, 30, "", maxPressureByState, getStateKey, "AA")

	fmt.Println(solution)
}

func getMaxPressureRelief(v *valve, minutesLeft int, fromValve string, cache map[string]int, getStateKey func(string, int) string, path string) int {
	if minutesLeft <= 0 {
		//fmt.Printf("%s | %d\n", path, 0)
		return 0
	}

	stateKey := getStateKey(v.name, minutesLeft)

	if val, ok := cache[stateKey]; ok {
		//fmt.Printf("%s | %d (cache %s)\n", path, val, stateKey)
		return val
	}

	maxPressure := 0

	for _, n := range v.neighbors {
		if n.name == fromValve {
			continue // Don't just go back the way you came immediately
		}

		testPressure := getMaxPressureRelief(n, minutesLeft-1, v.name, cache, getStateKey, path+" => "+n.name)
		maxPressure = util.MaxInt(maxPressure, testPressure)
	}

	if v.isClosed() {
		v.open()
		minutesLeft -= 1

		openingRelief := minutesLeft * v.flowRate

		testPressure := getMaxPressureRelief(v, minutesLeft, "", cache, getStateKey, path+" OP")
		maxPressure = util.MaxInt(maxPressure, testPressure+openingRelief)

		// Reset state before returning to the caller
		minutesLeft += 1
		v.close()
	}

	//fmt.Printf("%s | %d (%d, %s)\n", path, maxPressure, minutesLeft, stateKey)
	cache[stateKey] = maxPressure
	return maxPressure
}
