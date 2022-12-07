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
	f, err := os.Open("./input.txt")
	util.Check(err)
	defer f.Close()

	br := bufio.NewReader(f)

	buff := []byte{}
	cnts := make(map[byte]int)
	idx := 0

	for b, err := br.ReadByte(); err == nil; b, err = br.ReadByte() {
		idx += 1

		buff = append(buff, b)
		cnts[b] += 1

		if idx <= 4 {
			continue
		}

		out := buff[0]
		buff = buff[1:]

		cnts[out] -= 1
		if cnts[out] == 0 {
			delete(cnts, out)
		}

		if len(cnts) == 4 {
			break
		}
	}

	if !errors.Is(err, io.EOF) {
		util.Check(err)
	}

	fmt.Println(idx)
}
