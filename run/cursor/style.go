package cursor

type Style struct {
	Vert                    rune
	TopHori, BottomHori     rune
	Topleft, Topright       rune
	Bottomleft, Bottomright rune
}

func Fancy() Style {
	return Style{
		Vert:        BORDER_HORI,
		TopHori:     BORDER_VERT,
		BottomHori:  BORDER_VERT,
		Topleft:     BORDER_TL,
		Topright:    BORDER_TR,
		Bottomleft:  BORDER_BL,
		Bottomright: BORDER_BR,
	}
}

func Default() Style {
	return Style{
		TopHori:     BORDER_HORI,
		BottomHori:  SPACE,
		Vert:        SPACE,
		Topleft:     CORNER_TL,
		Topright:    CORNER_TR,
		Bottomleft:  CORNER_BL,
		Bottomright: CORNER_BR,
	}
}

func Space() Style {
	return Style{
		Vert:        SPACE,
		TopHori:     SPACE,
		BottomHori:  SPACE,
		Topleft:     SPACE,
		Topright:    SPACE,
		Bottomleft:  SPACE,
		Bottomright: SPACE,
	}
}
