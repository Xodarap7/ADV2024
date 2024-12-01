package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func processFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	var firstColumn []int
	var secondColumn []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("некорректная строка: %s", line)
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("ошибка парсинга чисел в строке: %s", line)
		}

		firstColumn = append(firstColumn, num1)
		secondColumn = append(secondColumn, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return firstColumn, secondColumn, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateTotalDistance(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	for i := 0; i < len(left); i++ {
		totalDistance += abs(left[i] - right[i])
	}

	return totalDistance
}

func calculateSimilarity(left, right []int) int {
	counts := make(map[int]int)
	for _, num := range right {
		counts[num]++
	}

	similarity := 0
	for _, num := range left {
		similarity += num * counts[num]
	}

	return similarity
}

func Start(filename string) (int, int, error) {
	firstColumn, secondColumn, err := processFile(filename)
	if err != nil {
		return 0, 0, err
	}

	totalDistance := calculateTotalDistance(firstColumn, secondColumn)

	similarity := calculateSimilarity(firstColumn, secondColumn)

	return totalDistance, similarity, nil
}
