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
		maxi := data[0]
		for _, v := range data {
			if v > maxi {
				maxi = v
			}
		}
		return maxi
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	size := len(data) / CHUNKS
	var eightMax []int

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		st := i * size
		end := st + size
		if end > len(data) || i == CHUNKS-1 {
			end = len(data)
		}

		go func(st, end int) {
			defer wg.Done()
			maxNum := data[st]
			for _, v := range data[st:end] {
				if v > maxNum {
					maxNum = v
				}
			}
			mu.Lock()
			eightMax = append(eightMax, maxNum)
			mu.Unlock()
		}(st, end)
	}
	wg.Wait()

	maxi := eightMax[0]
	for _, v := range eightMax {
		if v > maxi {
			maxi = v
		}
	}
	return maxi
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
	elapsed = time.Now().Sub(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxNumber, elapsed)
}
