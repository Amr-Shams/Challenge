package day8
import (
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)


const example = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func TestPartA(t *testing.T) {
    input:= challenge.FromLiteral(example)
    require.Equal(t, 6440, partA(input))
}
