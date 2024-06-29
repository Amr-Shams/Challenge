package day5 
import (
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)

func TestB(t *testing.T) {
    input := challenge.FromLiteral(example)
    require.Equal(t, 46, partB(input))
}
