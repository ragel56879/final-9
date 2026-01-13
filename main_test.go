package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	// 0 элементов
	require.Nil(t, nil, generateRandomElements(0), "Ожидается: %v\nПолучено: %v", nil, generateRandomElements(0))
}

type maxNumTestSuite struct {
	suite.Suite
}

func TestMaxNumSuite(t *testing.T) {
	suite.Run(t, new(maxNumTestSuite))
}

func (suite *maxNumTestSuite) TestMaximum() {
	var tests = []struct {
		name string
		in   []int
		want int
	}{
		{
			name: "пустой слайс",
			in:   []int{},
			want: 0,
		},
		{
			name: "слайс с 1 элементом",
			in:   []int{10},
			want: 10,
		},
		{
			name: "слайс с отрицательными числами",
			in:   []int{-1, -5, -3, -2, -10},
			want: -1,
		},
		{
			name: "слайс с повторяющими значениями",
			in:   []int{3, 3, 3, 3, 3, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			res := maximum(tt.in)
			require.Equal(suite.T(), tt.want, res, "Ожидается: %d\nПолучено: %d", tt.want, res)
		})
	}
}

func (suite *maxNumTestSuite) TestMaxChunks() {
	var tests = []struct {
		name string
		in   []int
		want int
	}{
		{
			name: "пустой слайс",
			in:   []int{},
			want: 0,
		},
		{
			name: "слайс с 1 элементом",
			in:   []int{5},
			want: 5,
		},
		{
			name: "слайс с кол-вом элементов <CHUNKS",
			in:   []int{1, 2, 3, 7, 5, 2},
			want: 7,
		},
		{
			name: "слайс с кол-вом элементов ==CHUNKS",
			in:   []int{1, 2, 3, 40, 5, 6, 7, 10},
			want: 40,
		},
		{
			name: "слайс с кол-вом элементов >CHUNKS % CHUNKS != 0",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 10,
		},
		{
			name: "слайс с отрицательными числами",
			in:   []int{-1, -5, -3, -2, -10},
			want: -1,
		},
		{
			name: "слайс с повторяющими значениями",
			in:   []int{3, 3, 3, 3, 3, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			res := maxChunks(tt.in)
			require.Equal(suite.T(), tt.want, res, "Ожидается: %d\nПолучено: %d", tt.want, res)
		})
	}
}
