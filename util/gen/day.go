package gen

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/util"
)

func genCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "all",
        Short: "Problems for Day 2A",
        Run: func(_ *cobra.Command, _ []string) {
            n:=25
            for i:=1;i<=n;i++{
                GenDay(i)
            }
            fmt.Println("All input files generated successfully")
        },
    }
}

func GenDay(n int) {
    p, err := util.DaysPath(n)
    if err != nil {
        fmt.Println(err)
        return
    }
    if _, err := os.Stat(p); os.IsNotExist(err) {
        os.Mkdir(p, os.ModePerm)
    }

    inputFilePath := filepath.Join(p, "input.txt")
    if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
        fmt.Println("Generating input file", n)
        problemInput, err := getInput(n)
        if err != nil {
            fmt.Println(err)
            return
        }
        if err := os.WriteFile(inputFilePath, []byte(problemInput), os.ModePerm); err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Input file generated successfully")
    } else {
        fmt.Println("Input file already exists")
    }
}

func getInput(n int) (string, error) {
    url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", n)
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }
    session,_ := util.GetSession()
    req.Header.Add("Cookie", fmt.Sprintf("session=%s", session))
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}