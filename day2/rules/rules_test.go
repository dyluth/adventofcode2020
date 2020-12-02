package rules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseLine(t *testing.T) {
	r := ParseLine("17-19 p: pwpzpfbrcpppjppbmppp")
	require.Equal(t, "p", r.letter)
	require.Equal(t, 17, r.minCount)
	require.Equal(t, 19, r.maxCount)
	require.Equal(t, "pwpzpfbrcpppjppbmppp", r.password)
}

func TestIsValid(t *testing.T) {
	r := Rule{
		minCount: 2,
		maxCount: 2,
		letter:   "a",
		password: "aa",
	}
	require.True(t, r.IsValid())

	r = Rule{
		minCount: 2,
		maxCount: 2,
		letter:   "a",
		password: "ab",
	}
	require.False(t, r.IsValid())
	r = Rule{
		minCount: 2,
		maxCount: 2,
		letter:   "a",
		password: "abaa",
	}
	require.False(t, r.IsValid())

	r = Rule{
		minCount: 2,
		maxCount: 5,
		letter:   "a",
		password: "abaabcc",
	}
	require.True(t, r.IsValid())

	r = Rule{
		minCount: 2,
		maxCount: 5,
		letter:   "a",
		password: "abaabccaaaaaa",
	}
	require.False(t, r.IsValid())

}

func TestIsValidPt2(t *testing.T) {

	tests := []struct {
		def   string
		valid bool
	}{
		{
			def:   "1-3 a: abcde",
			valid: true,
		},
		{
			def:   "1-3 b: cdefg",
			valid: false,
		},
		{
			def:   "2-9 c: ccccccccc",
			valid: false,
		},
	}

	for _, tt := range tests {
		p := ParseLine(tt.def)
		require.Equal(t, tt.valid, p.IsValidPt2())
	}

}
