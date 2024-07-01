package day7 
import (
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)




func TestPartB(t *testing.T) {
    input:= challenge.FromLiteral(example)
    require.Equal(t, 5905, partB(input))
}
