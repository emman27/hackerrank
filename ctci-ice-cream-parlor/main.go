package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func flavorIndexes(cost []int32, money int32) (int, int) {
	h := map[int32][]int{}
	for idx, c := range cost {
		h[c] = append(h[c], idx)
	}

	for idx, c := range cost {
		if money-c != c {
			if len(h[money-c]) > 0 {
				return idx, h[money-c][0]
			}
		} else {
			if len(h[money-c]) > 1 {
				return idx, h[money-c][1]
			}
		}
	}
	return 0, 0
}

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) {
	a, b := flavorIndexes(cost, money)
	fmt.Printf("%d %d\n", a+1, b+1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
