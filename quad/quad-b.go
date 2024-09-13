package piscine

func QuadB(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	// Top
	top := CreateString("/", "*", "\\", x)
	// Middle
	mid := CreateString("*", " ", "*", x)
	// Bottom
	bot := CreateString("\\", "*", "/", x)
	// Final Print
	FinalPrint(top, mid, bot, y)
}
