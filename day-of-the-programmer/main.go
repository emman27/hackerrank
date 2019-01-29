package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isLeapYearJulian(year int32) bool {
	return year%4 == 0
}

func isLeapYearGregorian(year int32) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func isGregorian(year int32) bool {
	return year > 1918
}

func isLeapYear(year int32) bool {
	if isGregorian(year) {
		return isLeapYearGregorian(year)
	}
	return isLeapYearJulian(year)
}

func isTransitionYear(year int32) bool {
	return year == 1918
}

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	if isTransitionYear(year) {
		return "26.09.1918"
	} else if isLeapYear(year) {
		return fmt.Sprintf("12.09.%d", year)
	} else if !isLeapYear(year) {
		return fmt.Sprintf("13.09.%d", year)
	}
	panic("Invalid state")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

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
