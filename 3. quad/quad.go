package quad

import "fmt"

// drawQuad prints a rectangle of width x and height y using custom corner and border runes
func drawQuad(x, y int, topLeft, topRight, botLeft, botRight, horiz, vert rune) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			switch {
			case row == 0 && col == 0:
				fmt.Printf("%c", topLeft)
			case row == 0 && col == x-1:
				fmt.Printf("%c", topRight)
			case row == y-1 && col == 0:
				fmt.Printf("%c", botLeft)
			case row == y-1 && col == x-1:
				fmt.Printf("%c", botRight)
			case row == 0 || row == y-1:
				fmt.Printf("%c", horiz)
			case col == 0 || col == x-1:
				fmt.Printf("%c", vert)
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func QuadA(x, y int) {
	drawQuad(x, y, 'o', 'o', 'o', 'o', '-', '|')
}

func QuadB(x, y int) {
	drawQuad(x, y, '/', '\\', '\\', '/', '*', '*')
}

func QuadC(x, y int) {
	drawQuad(x, y, 'A', 'A', 'C', 'C', 'B', 'B')
}

func QuadD(x, y int) {
	drawQuad(x, y, 'A', 'C', 'A', 'C', 'B', 'B')
}

func QuadE(x, y int) {
	drawQuad(x, y, 'A', 'C', 'C', 'A', 'B', 'B')
}
