package day7 
import (
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)


const example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestPartA(t *testing.T) {
    input:= challenge.FromLiteral(example)
    require.Equal(t, 6440, partA(input))
}
