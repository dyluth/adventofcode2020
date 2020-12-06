package plane

import "fmt"

type Plane struct {
	rows map[int]*Row
}

func NewPlane() Plane {
	return Plane{
		rows: make(map[int]*Row),
	}
}

func (p *Plane) FillSeat(s *Seat) {
	row, ok := p.rows[s.Row]
	if !ok {
		row = &Row{row: make([]*Seat, 8)}
		p.rows[s.Row] = row
	}
	row.FillSeat(s)
}

type Row struct {
	row []*Seat
}

func (r *Row) FillSeat(s *Seat) {
	r.row[s.Column] = s
}

func (p *Plane) PrintRowGaps() {
	for RowID, row := range p.rows {

		for col := range row.row {
			if row.row[col] == nil {
				ID := RowID*8 + col
				fmt.Printf("Row %v Col %v enpty (ID: %v)\n", RowID, col, ID)
			}
		}
	}

}
