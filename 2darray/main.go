package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	var offset []int32 = []int32{0, 1, 2}
	var tempMax int32 = -999
	var buffer int32 = 0
	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			buffer = buffer + arr[int32(i)+offset[0]][int32(j)+offset[0]]
			buffer = buffer + arr[int32(i)+offset[0]][int32(j)+offset[1]]
			buffer = buffer + arr[int32(i)+offset[0]][int32(j)+offset[2]]
			buffer = buffer + arr[int32(i)+offset[1]][int32(j)+offset[1]]
			buffer = buffer + arr[int32(i)+offset[2]][int32(j)+offset[0]]
			buffer = buffer + arr[int32(i)+offset[2]][int32(j)+offset[1]]
			buffer = buffer + arr[int32(i)+offset[2]][int32(j)+offset[2]]
			if buffer > tempMax {
				tempMax = buffer
			}
			buffer = 0
		}
	}
	return tempMax
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
