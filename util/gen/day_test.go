package gen

import (
    "os"
    "testing"
)

func TestGenDay(t *testing.T) {
    // Arrange
    day := 1
    expectedFilePath := "/Users/amraly/Projects/aoc2024/challenge/day1/"

    // Act
    GenDay(day)

    // Assert
    if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
        t.Errorf("GenDay() failed, expected file %s to exist", expectedFilePath)
    }
    // print what was generated by open     
    _, err := os.Open(expectedFilePath)
    if err != nil {
        t.Errorf("GenDay() failed, expected file %s to exist", expectedFilePath)
    }
}