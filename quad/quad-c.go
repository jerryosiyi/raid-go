package piscine

func QuadC(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	// Top
	top := CreateString("A", "B", "A", x)
	// MID
	mid := CreateString("B", " ", "B", x)
	// BOTTOM
	bot := CreateString("C", "B", "C", x)
	// FINAL PRINT
	FinalPrint(top, mid, bot, y)
}
