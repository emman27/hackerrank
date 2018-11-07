package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func getHour(s string) int {
	hour, err := strconv.Atoi(s[0:2])
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	return hour
}

func clearPostfix(s string) string {
	postfix := s[len(s)-2:]
	newHour := getHour(s)
	if postfix == "PM" {
		newHour += 12
	}
	return fmt.Sprintf("%02d%s", newHour, s[2:len(s)-2])
}

// fixHour fixes the 12th hour offset.
func fixHour(s string) string {
	newHour := getHour(s)
	if s[0:2] == "12" {
		newHour -= 12
	}
	return fmt.Sprintf("%02d%s", newHour, s[2:])
}

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	return clearPostfix(fixHour(s))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
