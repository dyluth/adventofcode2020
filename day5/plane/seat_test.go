package plane

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSeat(t *testing.T) {
	s := NewSeat("BFFFFFFRLR")
	require.Equal(t, 64, s.Row)
	require.Equal(t, 5, s.Column)

	s = NewSeat("FBFBBFFRLR")
	require.Equal(t, 44, s.Row)
	require.Equal(t, 5, s.Column)

	s = NewSeat("BFFFBBFRRR")
	require.Equal(t, 70, s.Row)
	require.Equal(t, 7, s.Column)
	require.Equal(t, 567, s.ID)

	s = NewSeat("FFFBBBFRRR")
	require.Equal(t, 14, s.Row)
	require.Equal(t, 7, s.Column)
	require.Equal(t, 119, s.ID)

	s = NewSeat("BBFFBBFRLL")
	require.Equal(t, 102, s.Row)
	require.Equal(t, 4, s.Column)
	require.Equal(t, 820, s.ID)

}
