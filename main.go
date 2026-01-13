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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	list := make([]int, size)
	if size == 0 {
		return nil
	}

	for i := 0; i < size; i++ {
		list[i] = rand.Int()
	}
	return list
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	maxNum := data[0]
	for _, v := range data {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	var wg sync.WaitGroup

	size := len(data) / CHUNKS
	eightMax := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		st := i * size
		end := st + size
		if end > len(data) || i == CHUNKS-1 {
			end = len(data)
		}

		go func(st, end int) {
			defer wg.Done()
			eightMax[i] = maximum(data[st:end])
		}(st, end)
	}
	wg.Wait()

	return maximum(eightMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	list := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	// ваш код здесь
	start := time.Now()
	maxNumber := maximum(list)
	elapsed := time.Now().Sub(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumber, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	maxNumber = maxChunks(list)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumber, elapsed)
}
