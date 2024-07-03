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
const example2 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`
func TestPartA(t *testing.T) {
    input:= challenge.FromLiteral(example)
    require.Equal(t, 6, partA(input))
    input = challenge.FromLiteral(example2)
    require.Equal(t, 2, partA(input))
}
