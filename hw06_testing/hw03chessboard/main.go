package hw03chessboard

func Draw(size int) string {
	s := ""
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if (i+j)%2 == 0 {
				s += " "
			} else {
				s += "#"
			}
		}

		if i != size {
			s += "\n"
		}
	}

	return s
}
