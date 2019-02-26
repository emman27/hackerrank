package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func stringToMap(s string) map[rune]int {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}
	return m
}

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	aMap := stringToMap(a)
	bMap := stringToMap(b)
	var count int32
	for key, value := range aMap {
		if value != bMap[key] {
			count += int32(math.Abs(float64(bMap[key] - value)))
			aMap[key] = bMap[key]
		}
	}
	for key, value := range bMap {
		if value != aMap[key] {
			count += int32(math.Abs(float64(aMap[key] - value)))
			aMap[key] = bMap[key]
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

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
