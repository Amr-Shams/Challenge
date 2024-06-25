package day4 
import (
    "testing"
    "github.com/Amr-Shams/aoc2024/challenge"
    "github.com/stretchr/testify/require"
)

func TestPartB(t *testing.T){
    input:=challenge.FromLiteral(example)
    result:=partB(input)
    require.Equal(t,30,result)
}
