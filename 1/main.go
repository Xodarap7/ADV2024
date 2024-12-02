package main

import "fmt"

func main() {
	filename := "day1.txt"

	totalDistance, similarity, err := start(filename)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Сумма расстояний между парами чисел:", totalDistance)
	fmt.Println("Оценка сходства:", similarity)
}
