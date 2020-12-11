package xmas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {

	v := IsValid([]int{35, 20, 15, 25, 47}, 40)
	require.True(t, v)

	require.False(t, IsValid([]int{35, 20, 15, 25, 47}, 4))

}

func TestFindFirstInvalid(t *testing.T) {
	in := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	f, err := FindFirstInvalid(in, 5)
	require.NoError(t, err)
	require.Equal(t, 127, f)
}

func TestFindFirstInvalid2(t *testing.T) {
	in := []int{0, 1, 1, 2, 3, 5, 7, 2, 20}
	f, err := FindFirstInvalid(in, 6)
	require.NoError(t, err)
	require.Equal(t, 20, f)
}

func TestFindFirstInvalid3(t *testing.T) {
	in := []int{84, 85, 95, 87, 109, 132, 126, 96, 108, 101, 129, 233, 103, 106, 172, 153, 110, 128, 175, 149, 150, 155, 165, 161, 181, 222, 180, 242, 422}
	_, err := FindFirstInvalid(in, 25)
	require.Error(t, err)
}

func TestFindWeakness(t *testing.T) {
	r := FindWeakness([]int{50, 43, 55, 2, 3, 4, 5, 7, 60, 32}, 12)
	require.EqualValues(t, []int{3, 4, 5}, r)
}
