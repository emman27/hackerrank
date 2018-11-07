package main

// BUG: Test case for this problem is broken

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func swap(arr []int32, curr, next int) []int32 {
	tmp := arr[curr]
	arr[curr] = arr[next]
	arr[next] = tmp
	return arr
}

func find(arr []int32, val int32, startIndex int) int {
	for i := startIndex; i < len(arr); i++ {
		if val == arr[i] {
			return i
		}
	}
	return -1
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swaps int32
	for i, n := range arr {
		if n != int32(i+1) {
			arr = swap(arr, i, find(arr, int32(i+1), i+1))
			swaps++
		}
	}
	return swaps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

	writer.Flush()
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
