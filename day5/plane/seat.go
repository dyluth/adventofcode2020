package plane

import "strings"

type Seat struct {
	Row    int
	Column int
	ID     int
}

var (
	conv map[rune]int
)

func init() {
	if conv == nil {
		conv = make(map[rune]int)
		conv['F'] = 0
		conv['B'] = 1
		conv['L'] = 0
		conv['R'] = 1
	}
}

func NewSeat(data string) (s Seat) {
	data = strings.TrimSpace(data)
	//data has: first 7 characters F or B (binary, F=0)
	r := []rune(data)

	for i := 0; i < 7; i++ {
		binary, _ := conv[r[i]]
		value := binary << (6 - i)
		s.Row = s.Row + value
	}
	r2 := r[7:]

	for i := 0; i < 3; i++ {
		binary, _ := conv[r2[i]]
		value := binary << (2 - i)
		s.Column = s.Column + value
	}

	s.ID = s.Row*8 + s.Column

	return s
}
