package piscine

func QuadE(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	// Top
	top := CreateString("A", "B", "C", x)
	// MID
	mid := CreateString("B", " ", "B", x)
	//BOT
	bot := CreateString("C", "B", "A", x)
	// FINAL PRINT
	FinalPrint(top, mid, bot, y)
}
