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

var randInt = rand.Int

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{} // обработка крайних случаев
	}

	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = randInt()
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	maxValue := 0
	if data == nil || len(data) == 0 {
		return maxValue
	}
	for _, v := range data {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < 8 {
		return maximum(data)
	}
	chunks := [][]int{}
	chunkSize := len(data) / CHUNKS
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		chunks = append(chunks, data[start:end])
	}
	maxs := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		part := chunks[i]
		go func() {
			defer wg.Done()
			maxs[i] = maximum(part)
		}()
	}
	wg.Wait()
	return maximum(maxs)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	list := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxVal := maximum(list)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed.Milliseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxVal = maxChunks(list)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxVal, elapsed.Milliseconds())
}
