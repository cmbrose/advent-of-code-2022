package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"main/util"
	"os"
)

func main() {
	target := 14

	f, err := os.Open("./input.txt")
	util.Check(err)
	defer f.Close()

	br := bufio.NewReader(f)

	// buff is treated as a ring buffer
	buff := make([]byte, target)
	// cached count of each byte in buff
	cnts := make(map[byte]int)
	idx := 0

	for b, err := br.ReadByte(); err == nil; b, err = br.ReadByte() {
		idx += 1

		out := buff[idx%target]
		buff[idx%target] = b

		cnts[b] += 1

		if idx <= target {
			continue
		}

		// Decrement the removed byte - remove it instead if it hits 0
		cnt := cnts[out] - 1
		if cnt == 0 {
			delete(cnts, out)
		} else {
			cnts[out] = cnt
		}

		if len(cnts) == target {
			break
		}
	}

	if !errors.Is(err, io.EOF) {
		util.Check(err)
	}

	fmt.Println(idx)
}
