package day9
import (
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)


const example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
func TestPartA(t *testing.T) {
    input:= challenge.FromLiteral(example)
    require.Equal(t, 114, partA(input))
 }
