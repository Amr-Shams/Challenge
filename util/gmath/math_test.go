package gmath

import (
"fmt"
	"testing"

//	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestIntMin(t *testing.T) {
    tests := []struct {
        a int 
        b int
        expected int
    }{
        {1, 2, 1},
        {2, 1, 1},
        {0, 0, 0},
        {0, 1, 0},
        {1, 0, 0},
        {-1, 1, -1}}
    for _, test := range tests {
        t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
            result := Min(test.a, test.b)
            assert.Equal(t, test.expected, result)
        })
    }
}

func TestIntMax(t *testing.T) {
    tests := []struct {
        a int 
        b int
        expected int
    }{
        {1, 2, 2},
        {2, 1, 2},
        {0, 0, 0},
        {0, 1, 1},
        {1, 0, 1},
        {-1, 1, 1}}
    for _, test := range tests {
        t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
            result := Max(test.a, test.b)
            assert.Equal(t, test.expected, result)
        })
    }
}
