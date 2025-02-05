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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	encrypted := strings.Builder{}

	for _, c := range s {
		if isLetter(c) {
			c = encryptRune(c, k)
		}
		encrypted.WriteRune(c)
	}
	return encrypted.String()
}

func isLetter(r rune) bool {
	return isLowerCase(r) || isUpperCase(r)
}

func isLowerCase(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isUpperCase(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func encryptRune(r rune, k int32) rune {
	if isLowerCase(r) && !isLowerCase(r+k) {
		return (r-'a'+k)%26 + 'a'
	}

	if isUpperCase(r) && !isUpperCase(r+k) {
		return (r-'A'+k)%26 + 'A'
	}

	return r + k
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	_ = int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

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
