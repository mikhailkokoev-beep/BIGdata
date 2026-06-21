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

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = rand.Intn(1000000) + 1
	}
	return elements
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxVal {
			maxVal = data[i]
		}
	}
	return maxVal
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
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
			maxValues[index] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("\nИщем максимальное значение в один поток")
	startSeq := time.Now()
	maxSeq := maximum(data)
	elapsedSeq := time.Since(startSeq).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxSeq, elapsedSeq)

	fmt.Printf("\nИщем максимальное значение в %d потоков\n", CHUNKS)
	startConc := time.Now()
	maxConc := maxChunks(data)
	elapsedConc := time.Since(startConc).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxConc, elapsedConc)

	if maxSeq == maxConc {
		fmt.Println("\n✓ Оба метода нашли одинаковый максимум")
	} else {
		fmt.Println("\n✗ Ошибка: методы нашли разные максимумы")
	}
}