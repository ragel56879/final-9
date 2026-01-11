package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	// 0 элементов
	require.Nil(t, nil, generateRandomElements(0), "Ожидается: %v\nПолучено: %v", nil, generateRandomElements(0))
}

func TestMaximum(t *testing.T) {
	// пустой слайс
	var data []int
	expected := 0
	require.Equal(t, expected, maximum(data), "Ожидается: %d\nПолучено: %d", expected, maximum(data))

	// слайс с 1 элементом
	data = []int{10}
	expected = 10
	require.Equal(t, expected, maximum(data), "Ожидается: %d\nПолучено: %d", expected, maximum(data))

	// слайс с отрицательными числами
	data = []int{-1, -5, -3, -2, -10}
	expected = -1
	require.Equal(t, expected, maximum(data), "Ожидается: %d\nПолучено: %d", expected, maximum(data))

	// слайс с повторяющими значениями
	data = []int{3, 3, 3, 3, 3, 3}
	expected = 3
	require.Equal(t, expected, maximum(data), "Ожидается: %d\nПолучено: %d", expected, maximum(data))
}

func TestMaxChunks(t *testing.T) {
	// пустой слайс
	var data []int
	expected := 0
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))

	// слайс с 1 элементом
	data = []int{5}
	expected = 5
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))

	// слайс с кол-вом элементов <CHUNKS
	data = []int{1, 2, 3, 7, 5, 2}
	expected = 7
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))

	// слайс с кол-вом элементов ==CHUNKS
	data = []int{1, 2, 3, 40, 5, 6, 7, 10}
	expected = 40
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))

	// слайс с кол-вом элементов >CHUNKS % CHUNKS != 0
	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected = 10
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))

	// слайс с отрицательными числами
	data = []int{-1, -5, -3, -2, -10}
	expected = -1
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))

	// слайс с повторяющими значениями
	data = []int{3, 3, 3, 3, 3, 3}
	expected = 3
	require.Equal(t, expected, maxChunks(data), "Ожидается: %d\nПолучено: %d", expected, maxChunks(data))
}
