package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func minBribes(q []int32) string {
	total := 0
	swapped := map[int32]int{-1: -1}
	for len(swapped) != 0 {
		swapped = map[int32]int{}
		for i, sticker := range q {
			if i < len(q)-1 && sticker > q[i+1] {
				q[i], q[i+1] = q[i+1], q[i]
				swapped[q[i+1]]++
				total++
			}
		}
		for _, value := range swapped {
			if value > 2 {
				return "Too chaotic"
			}
		}
	}
	return fmt.Sprintf("%d", total)
}

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	fmt.Println(minBribes(q))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
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
