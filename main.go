package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements создаёт слайс случайных положительных чисел.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}
	return data
}

// maximum возвращает максимальное значение слайса.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks находит максимальное значение слайса с помощью восьми горутин.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			// Для пустой части используем первый элемент исходного слайса:
			// он не изменит общий максимум.
			if len(chunk) == 0 {
				maxValues[index] = data[0]
			} else {
				maxValues[index] = maximum(chunk)
			}
		}(data[start:end], i)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxValue := maximum(data)
	elapsedMicroseconds := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsedMicroseconds)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	startTime = time.Now()
	maxValue = maxChunks(data)
	elapsedMicroseconds = time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsedMicroseconds)
}
