package piscine

func QuadD(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	// Top
	top := CreateString("A", "B", "C", x)
	// MID
	mid := CreateString("B", " ", "B", x)
	// FINAL PRINT
	FinalPrint(top, mid, top, y)
}
