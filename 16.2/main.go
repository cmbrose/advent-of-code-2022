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

	getStateKey := func(you, ele string, cur, minutes int) string {
		if you > ele {
			you, ele = ele, you
		}
		return fmt.Sprintf("%s_%s_%d_%s_%d", you, ele, cur, hex.EncodeToString(closedStates.Bytes()), minutes)
	}

	bailEarly := func(minutes int) bool {
		return false
	}

	start := valves["AA"]
	solution := getMaxPressureRelief([]*valve{start, start}, 0, 26, []string{"", ""}, maxPressureByState, getStateKey, bailEarly)

	fmt.Println(solution)
}

func getMaxPressureRelief(valves []*valve, cur int, minutesLeft int, from []string, cache map[string]int, getStateKey func(string, string, int, int) string, bailEarly func(int) bool) int {
	if bailEarly(minutesLeft) {
		return 0
	}

	if minutesLeft <= 0 {
		//fmt.Printf("%s | %d\n", path, 0)
		return 0
	}

	stateKey := getStateKey(valves[0].name, valves[1].name, cur, minutesLeft)

	// if val, ok := cache[stateKey]; ok {
	// 	//fmt.Printf("%s | %d (cache %s)\n", path, val, stateKey)
	// 	return val
	// }

	maxPressure := 0

	v := valves[cur]
	f := from[cur]

	for _, n := range v.neighbors {
		if from[cur] == n.name {
			continue
		}

		valves[cur] = n
		from[cur] = v.name

		testPressure := getMaxPressureRelief(valves, 1-cur, minutesLeft-cur, from, cache, getStateKey, bailEarly)
		maxPressure = util.MaxInt(maxPressure, testPressure)
	}

	valves[cur] = v

	if v.isClosed() {
		v.open()

		from[cur] = ""

		testPressure := getMaxPressureRelief(valves, 1-cur, minutesLeft-cur, from, cache, getStateKey, bailEarly)

		openingRelief := (minutesLeft - 1) * v.flowRate
		maxPressure = util.MaxInt(maxPressure, testPressure+openingRelief)

		// Reset state before returning to the caller
		minutesLeft += cur
		v.close()
	}

	from[cur] = f

	//fmt.Printf("%s | %d (%d, %s)\n", path, maxPressure, minutesLeft, stateKey)
	cache[stateKey] = maxPressure
	return maxPressure
}
