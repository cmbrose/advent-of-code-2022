package main

import (
	"encoding/hex"
	"fmt"
	"main/util"
	"math"
	"regexp"
	"strings"

	"github.com/dropbox/godropbox/container/bitvector"
)

type valve struct {
	name      string
	flowRate  int
	neighbors []struct {
		v *valve
		w int
	}

	isClosed func() bool
	close    func()
	open     func()
}

func main() {
	re := regexp.MustCompile(`Valve (.+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

	valves := make(map[string]*valve)
	valvesArr := []*valve{}

	for _, line := range util.ReadInputLines("input.txt") {
		match := re.FindStringSubmatch(line)

		var (
			name         = match[1]
			rate         = util.AssertInt(match[2])
			neighborsStr = match[3]
			neighborsArr = strings.Split(neighborsStr, ", ")
		)

		v := &valve{
			name:     name,
			flowRate: rate,
		}

		valvesArr = append(valvesArr, v)
		valves[name] = v

		for _, name := range neighborsArr {
			if n, ok := valves[name]; ok {
				n.neighbors = append(n.neighbors, struct {
					v *valve
					w int
				}{v, 1})
				v.neighbors = append(v.neighbors, struct {
					v *valve
					w int
				}{n, 1})
			}
		}
	}

	valves = minifyValves(valvesArr)
	closedStates := bitvector.NewBitVector(nil, 0)

	for _, v := range valves {
		id := closedStates.Length()
		closedStates.Append(0)

		v.open = func() {
			closedStates.Set(0, id)
		}
		v.close = func() {
			closedStates.Set(1, id)
		}
		v.isClosed = func() bool {
			return closedStates.Element(id) == 1
		}
	}

	solution := divideAndConquer(closedStates, valves)

	fmt.Println(solution)
}

func minifyValves(valves []*valve) map[string]*valve {
	valveToIndex := make(map[*valve]int)
	nonZeroValveToIndex := make(map[*valve]int)
	for i, v := range valves {
		valveToIndex[v] = i

		if v.flowRate > 0 || v.name == "AA" {
			nonZeroValveToIndex[v] = i
		}
	}

	grid := util.FillGrid(len(valves), len(valves), math.MaxInt)
	for i, v := range valves {
		grid[i][i] = 0
		for _, n := range v.neighbors {
			j := valveToIndex[n.v]
			grid[i][j] = n.w
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				grid[i][j] = util.MaxInt(grid[i][j], grid[i][k]+grid[k][j])
			}
		}
	}

	minified := make(map[string]*valve)

	for v, i := range nonZeroValveToIndex {
		clone := &valve{
			name:      v.name,
			flowRate:  v.flowRate,
			neighbors: nil,
		}

		for u, j := range nonZeroValveToIndex {
			clone.neighbors = append(clone.neighbors, struct {
				v *valve
				w int
			}{u, grid[i][j]})
		}

		minified[v.name] = v
	}

	return minified
}

func divideAndConquer(sourceBitVector *bitvector.BitVector, valves map[string]*valve) int {
	bitIdToValveWithFlow := make(map[int]*valve)

	for _, v := range valves {
		if v.flowRate > 0 {
			bitIdToValveWithFlow[len(bitIdToValveWithFlow)] = v
		}
	}

	totalWithFlow := len(bitIdToValveWithFlow)

	cloneVector := bitvector.NewBitVector(sourceBitVector.Bytes(), sourceBitVector.Length())
	reset := func() {
		for i := 0; i < sourceBitVector.Length(); i += 1 {
			sourceBitVector.Set(cloneVector.Element(i), i)
		}
	}

	max := 0

	for i := 0; i < (1 << totalWithFlow); i += 1 {
		for j := 0; j < totalWithFlow; j += 1 {
			bit := (i >> j) & 1

			v := bitIdToValveWithFlow[j]
			if bit == 0 {
				v.close()
			} else {
				v.open()
			}
		}

		you := getMaxPressure(sourceBitVector, valves)

		reset()

		for j := 0; j < totalWithFlow; j += 1 {
			bit := (i >> j) & 1

			v := bitIdToValveWithFlow[j]
			if bit == 0 {
				v.open()
			} else {
				v.close()
			}
		}

		ele := getMaxPressure(sourceBitVector, valves)

		reset()

		if you+ele > max {
			max = you + ele
			fmt.Println(i, max)
		}
	}

	return max
}

func getMaxPressure(closedStates *bitvector.BitVector, valves map[string]*valve) int {
	maxPressureByState := make(map[string]int)

	getStateKey := func(node string, minutes int) string {
		return fmt.Sprintf("%s_%s_%d", node, hex.EncodeToString(closedStates.Bytes()), minutes)
	}

	totalPressure := 0
	for _, v := range valves {
		totalPressure += v.flowRate
	}

	start := valves["AA"]
	return getMaxPressureReliefRec(start, 26, "", maxPressureByState, getStateKey, totalPressure)
}

func getMaxPressureReliefRec(v *valve, minutesLeft int, fromValve string, cache map[string]int, getStateKey func(string, int) string, remainingPressure int) int {
	if minutesLeft <= 0 || remainingPressure == 0 {
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
		if n.v.name == fromValve {
			continue // Don't just go back the way you came immediately
		}

		testPressure := getMaxPressureReliefRec(n.v, minutesLeft-n.w, v.name, cache, getStateKey, remainingPressure)
		maxPressure = util.MaxInt(maxPressure, testPressure)
	}

	if v.isClosed() {
		v.open()
		minutesLeft -= 1

		openingRelief := minutesLeft * v.flowRate

		testPressure := getMaxPressureReliefRec(v, minutesLeft, "", cache, getStateKey, remainingPressure-v.flowRate)
		maxPressure = util.MaxInt(maxPressure, testPressure+openingRelief)

		// Reset state before returning to the caller
		v.close()
	}

	//fmt.Printf("%s | %d (%d, %s)\n", path, maxPressure, minutesLeft, stateKey)
	cache[stateKey] = maxPressure
	return maxPressure
}
