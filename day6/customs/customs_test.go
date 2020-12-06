package customs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUnique(t *testing.T) {
	a := NewAnswers("abc")
	require.Equal(t, 3, a.GetUnique())
	a = NewAnswers("abc\nbc")
	require.Equal(t, 3, a.GetUnique())
	a = NewAnswers("a\nb\nc")
	require.Equal(t, 3, a.GetUnique())
	a = NewAnswers("ab\nac")
	require.Equal(t, 3, a.GetUnique())
	a = NewAnswers("a\na \na\na")
	require.Equal(t, 1, a.GetUnique())
	a = NewAnswers("b\n")
	require.Equal(t, 1, a.GetUnique())
}

func TestGetConsistent(t *testing.T) {
	a := NewAnswers("abc")
	require.Equal(t, 3, a.GetConsistent())
	a = NewAnswers("abc\nbc")
	require.Equal(t, 2, a.GetConsistent())
	a = NewAnswers("a\nb\nc")
	require.Equal(t, 0, a.GetConsistent())
	a = NewAnswers("ab\nac")
	require.Equal(t, 1, a.GetConsistent())
	a = NewAnswers("a\na \na\na")
	require.Equal(t, 1, a.GetConsistent())
	a = NewAnswers("b\n")
	require.Equal(t, 1, a.GetConsistent())
}
