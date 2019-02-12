package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countSwaps function below.
func countSwaps(a []int32) {
	fmt.Printf("Array is sorted in %d swaps.\n", countSwapsInternal(a))
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d", a[len(a)-1])
}

func countSwapsInternal(a []int32) int {
	count := 0
	for range a {
		for i := 0; i < len(a)-1; i++ {
			j := i + 1
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				count++
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
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
