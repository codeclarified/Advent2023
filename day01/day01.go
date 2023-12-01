package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func quickReverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func main() {
	// numMap := map[string]string{
	// 	"one": "1",
	// 	"two": "2",
	// 	"three": "3",
	// 	"four": "4",
	// 	"five": "5",
	// 	"six": "6",
	// 	"seven": "7",
	// 	"eight": "8",
	// 	"nine": "9",
	// }
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	success := scanner.Scan()

	total := 0

	for success {
		currentLine := scanner.Text()

		// replace text numbers with ints

		reverseLine := quickReverse(currentLine)
		tens := 0
		ones := 0
		fmt.Println(currentLine)

		lineData := strings.Split(currentLine, "")
		reverseLineData := strings.Split(reverseLine, "")

		for _, v := range lineData {
			val, err := strconv.Atoi(v)
			if err == nil {
				tens = val
				break
			}
		}
		for _, v := range reverseLineData {
			val, err := strconv.Atoi(v)
			if err == nil {
				ones = val
				break
			}
		}

		result := (tens * 10) + ones
		total += result

		success = scanner.Scan()
	}

	fmt.Println("Result: ", total)
}
