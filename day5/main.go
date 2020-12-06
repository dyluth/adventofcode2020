package main

import (
	"fmt"

	"github.com/dyluth/adventofcode2020/day5/plane"
	"github.com/dyluth/adventofcode2020/selectionbox"
)

func main() {
	input := selectionbox.ReadInput()
	maxSeat := plane.Seat{}
	p := plane.NewPlane()
	for i := range input {
		seat := plane.NewSeat(input[i])
		if seat.ID > maxSeat.ID {
			maxSeat = seat
		}
		p.FillSeat(&seat)
	}
	fmt.Printf("MaxID = %v (%v,%v)\n", maxSeat.ID, maxSeat.Row, maxSeat.Column)
	p.PrintRowGaps()
}
