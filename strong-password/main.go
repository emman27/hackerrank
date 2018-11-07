package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type password string

const (
	digits    = "0123456789"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special   = "!@#$%^&*()-+"
	minLength = 6
)

func (p *password) hasMinLength() bool {
	return len(*p) >= minLength
}

func isInCharSet(r rune, charset string) bool {
	for _, c := range charset {
		if c == r {
			return true
		}
	}
	return false
}

func (p *password) hasFromCharset(s string) bool {
	for _, c := range *p {
		if isInCharSet(c, s) {
			return true
		}
	}
	return false
}

// Complete the minimumNumber function below.
func minimumNumber(n int32, pass string) int32 {
	// Return the minimum number of characters to make the password strong
	p := password(pass)
	var addedChars int32
	if !p.hasFromCharset(digits) {
		p = password(fmt.Sprintf("%s%c", p, digits[0]))
		addedChars++
	}
	if !p.hasFromCharset(lowercase) {
		p = password(fmt.Sprintf("%s%c", p, lowercase[0]))
		addedChars++
	}
	if !p.hasFromCharset(uppercase) {
		p = password(fmt.Sprintf("%s%c", p, lowercase[0]))
		addedChars++
	}
	if !p.hasFromCharset(special) {
		p = password(fmt.Sprintf("%s%c", p, lowercase[0]))
		addedChars++
	}
	for !(p.hasMinLength()) {
		p = password(fmt.Sprintf("%s%c", p, 'a'))
		addedChars++
	}
	return addedChars
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

	pass := readLine(reader)

	answer := minimumNumber(n, pass)

	fmt.Fprintf(writer, "%d\n", answer)

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
