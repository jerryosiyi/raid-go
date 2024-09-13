package piscine

import "fmt"

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	// Top & Bottom
	top := CreateString("o", "-", "o", x)
	// Middle
	mid := CreateString("|", " ", "|", x)
	// FInal Print
	FinalPrint(top, mid, top, y)
}

func CreateString(a, b, c string, l int) string {
	res := ""
	if l == 1 {
		res = a
	} else {
		res += a
		for i := 0; i < l-2; i++ {
			res += b
		}
		res += c
	}
	return res
}

func FinalPrint(t, m, b string, l int) {
	if l == 1 {
		fmt.Println(t)
	} else {
		fmt.Println(t)
		for i := 0; i < l-2; i++ {
			fmt.Println(m)
		}
		fmt.Println(b)
	}
}
