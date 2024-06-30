package day6 
import(
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/Amr-Shams/aoc2024/challenge"
)

func TestB(t *testing.T) {
    input := challenge.FromLiteral(example)
    require.Equal(t, 71503, partB(input))
}
