
package day6 

import (
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)

const example=`Time:      7  15   30
Distance:  9  40  200`


func TestPartA(t *testing.T) {
    input:=challenge.FromLiteral(example)
    require.Equal(t,288, partA(input))
}
