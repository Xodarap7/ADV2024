package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := "data.txt"

	lines, err1 := processFile(filename)
	result1 := getResult(lines)

	if err1 != nil {
		fmt.Println("Ошибка : ", err1)
	}

	fmt.Println("Cумма результатов (первая часть задачи):", result1)
}

func getResult(lines []string) int{
	result := 0
	filteredArr := filterStrings(lines)

	for _, arr := range filteredArr{
		result += arr[0] * arr[1]
	}

	return result
}

func filterStrings(lines []string)(list [][]int){
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	var result [][]int

	for _, line := range lines{
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result = append(result, []int{num1, num2})
		}
	}

	return result
}

func processFile(filename string) ([] string, error) {
	var result []string
	file, err := os.Open(filename)
	if err != nil {
		return result, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		result = append(result, line)
	}

	if err := s.Err(); err != nil {
		return result, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return result, nil
}
