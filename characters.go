package main

type fontInfo struct {
	width  int
	matrix [][]int
}

var font = map[rune]fontInfo{
	'L': {
		width: 4,
		matrix: [][]int{
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
		},
	},
	'U': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
	},
	'D': {
		width: 5,
		matrix: [][]int{
			{1, 1, 1, 0, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 1, 1, 0, 0, 0, 0},
		},
	},
	'I': {
		width: 1,
		matrix: [][]int{
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
		},
	},
	' ': {
		width: 1,
		matrix: [][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	},
	'R': {
		width: 5,
		matrix: [][]int{
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 1, 0, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'E': {
		width: 4,
		matrix: [][]int{
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
		},
	},
	'H': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'A': {
		width: 5,
		matrix: [][]int{
			{0, 0, 1, 0, 0, 0, 0},
			{0, 1, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'K': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0},
			{1, 0, 1, 0, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'O': {
		width: 5,
		matrix: [][]int{
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
	},
	'❤': {
		width: 7,
		matrix: [][]int{
			{0, 1, 1, 0, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
		},
	},
	'G': {
		width: 5,
		matrix: [][]int{
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 0},
			{1, 0, 1, 1, 1},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
	},
	'P': {
		width: 5,
		matrix: [][]int{
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{1, 0, 0, 0, 0},
		},
	},
	'S': {
		width: 5,
		matrix: [][]int{
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 1},
			{1, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 1},
			{1, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
	},
	'!': {
		width: 1,
		matrix: [][]int{
			{1},
			{1},
			{1},
			{1},
			{1},
			{0},
			{1},
		},
	},
	'.': {
		width: 1,
		matrix: [][]int{
			{0},
			{0},
			{0},
			{0},
			{0},
			{0},
			{1},
		},
	},
	'😀': {
		width: 7,
		matrix: [][]int{
			{0, 0, 1, 1, 1, 0, 0},
			{0, 1, 0, 0, 0, 1, 0},
			{1, 0, 1, 0, 1, 0, 1},
			{1, 1, 0, 0, 0, 1, 1},
			{1, 0, 1, 1, 1, 0, 1},
			{0, 1, 0, 0, 0, 1, 0},
			{0, 0, 1, 1, 1, 0, 0},
		},
	},
	'B': {
		width: 5,
		matrix: [][]int{
			{0, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
		},
	},
	'C': {
		width: 4,
		matrix: [][]int{
			{0, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0},
		},
	},
	'F': {
		width: 4,
		matrix: [][]int{
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
		},
	},
	'J': {
		width: 5,
		matrix: [][]int{
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{1, 0, 0, 1, 0, 0, 0},
			{0, 1, 1, 0, 0, 0, 0},
		},
	},
	'M': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 0, 1, 1, 0, 0},
			{1, 0, 1, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'N': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{1, 1, 0, 0, 1, 0, 0},
			{1, 0, 1, 0, 1, 0, 0},
			{1, 0, 0, 1, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 1},
		},
	},
	'Q': {
		width: 5,
		matrix: [][]int{
			{0, 1, 1, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 1, 0, 1, 0, 0},
			{0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 0},
		},
	},
	'T': {
		width: 5,
		matrix: [][]int{
			{1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
		},
	},
	'V': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{0, 1, 0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
		},
	},
	'W': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
			{1, 0, 1, 0, 1, 0, 0},
			{1, 1, 0, 1, 1, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'X': {
		width: 5,
		matrix: [][]int{
			{1, 0, 0, 0, 1, 0, 0},
			{0, 1, 0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 1, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 1, 0, 0},
		},
	},
	'Y': {
		width: 7,
		matrix: [][]int{
			{1, 0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 0, 1, 0},
			{0, 0, 1, 0, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
		},
	},
	'Z': {
		width: 7,
		matrix: [][]int{
			{1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 1, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1},
		},
	},
	'u': {
		width: 4,
		matrix: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{0, 1, 1, 0},
		},
	},
	'r': {
		width: 4,
		matrix: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 0, 1, 1},
			{1, 1, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
	},
	'e': {
		width: 4,
		matrix: [][]int{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{1, 0, 0, 1},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{0, 1, 1, 1},
		},
	},
	's': {
		width: 5,
		matrix: [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 1},
			{0, 1, 1, 1, 0},
		},
	},
	'h': {
		width: 4,
		matrix: [][]int{
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 1, 1, 0},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
		},
	},
}
