package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func processFile(filename string, checkSafety func([]string) bool) (int, error) {
	result := 0
	file, err := os.Open(filename)

	if err != nil {
		return result, fmt.Errorf("ошибка при открытии файла: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if checkSafety(parts) {
			result++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return result, nil
}

func isSafeReport(report []string) bool {
	if len(report) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i, currValue := range report[1:] {
		currValue, err1 := strconv.Atoi(currValue)
		prevValue, err2 := strconv.Atoi(report[i])

		if err1 != nil || err2 != nil {
			return false
		}

		diff := currValue - prevValue

		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}

		if diff < 0 {
			increasing = false
		} else if diff > 0 {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func isSafeWithDampener(report []string) bool {
	if isSafeReport(report) {
		return true
	}
	for i := range report{
		newReport := append([]string{}, report[:i]...)
		newReport = append(newReport, report[i+1:]...)
		if isSafeReport(newReport) {
			return true
		}
	}
	return false
}

func main() {
	filename := "data.txt"

	result1, err1 := processFile(filename, isSafeReport)
	result2, err2 := processFile(filename, isSafeWithDampener)

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка")
		return
	}

	fmt.Println("Количество безопасных отчетов (первая часть задачи):", result1)
	fmt.Println("Количество безопасных отчетов (вторая часть задачи):", result2)
}