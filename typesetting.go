package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func printMessage(message [][]int) {
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if message[i][j] > 0 {
				fmt.Print(" . ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func render(msg string) [][]int {
	matrix := make([][]int, numRows)
	for i := range matrix {
		matrix[i] = make([]int, numCols)
	}

	msgWidth := 0
	for _, char := range msg {
		msgWidth += font[char].width
	}

	numRunes := utf8.RuneCountInString(msg)
	spaceBetweenChars := (numCols - msgWidth) / (numRunes - 1)
	if spaceBetweenChars < 1 {
		fmt.Printf("Message %q is too long. Try a shorter message.", msg)
		os.Exit(1)
	}
	spaceBetweenChars = 1
	margins := numCols - msgWidth - spaceBetweenChars*(numRunes-1)

	xOffset := margins / 2
	for _, char := range msg {
		if charInfo, ok := font[char]; ok {
			charWidth := charInfo.width
			charMatrix := charInfo.matrix
			for r := 0; r < numRows; r++ {
				for c := 0; c < charWidth; c++ {
					matrix[r][c+xOffset] = charMatrix[r][c]
				}
			}
			xOffset += charWidth + spaceBetweenChars
		} else {
			fmt.Printf("Character %q is not recognized. Add it to characters.go and subimt a PR.\n", string(char))
			os.Exit(1)
		}
	}
	return matrix
}
