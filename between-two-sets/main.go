package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func findMin(arr []int32) int32 {
	min := arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func isFactor(factor, number int32) bool {
	return math.Remainder(float64(number), float64(factor)) == 0
}

func findCommonFactors(arr []int32) []int32 {
	minimum := findMin(arr)
	factors := []int32{}
	var i int32
	for i = 1; i <= minimum; i++ {
		isFactorOfAll := true
		for _, elem := range arr {
			isFactorOfAll = isFactorOfAll && isFactor(i, elem)
		}
		if isFactorOfAll {
			factors = append(factors, i)
		}
	}
	return factors
}

func isMultipleOfAll(numbers []int32, query int32) bool {
	isMultiple := true
	for _, n := range numbers {
		isMultiple = isMultiple && isFactor(n, query)
	}
	return isMultiple
}

/*
 * Complete the getTotalX function below.
 */
func getTotalX(a []int32, b []int32) int32 {
	/*
	 * Write your code here.
	 */
	factors := findCommonFactors(b)
	qualified := []int32{}
	for _, factor := range factors {
		if isMultipleOfAll(a, factor) {
			qualified = append(qualified, factor)
		}
	}
	fmt.Println(qualified)
	return int32(len(qualified))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for aItr := 0; aItr < int(n); aItr++ {
		aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(readLine(reader), " ")

	var b []int32

	for bItr := 0; bItr < int(m); bItr++ {
		bItemTemp, err := strconv.ParseInt(bTemp[bItr], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	total := getTotalX(a, b)

	fmt.Fprintf(writer, "%d\n", total)

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
