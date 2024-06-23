package day3 
import(
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
   // "github.com/Amr-Shams/aoc2024/util"
   "github.com/Amr-Shams/aoc2024/util/matrix"

)

func aCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "a",
        Short: "Problems for Day 3A",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
        },
    }
}

func partA(challenge *challenge.Input) int {
   mat:= matrix.CreateMatrixFrom[rune](challenge, matrix.ConvertToRune)
   sum:=0
 
   for i:=0; i<mat.Rows;i++{
    for j:=0; j<mat.Cols;j++{
        r:=mat.Get(i,j)
   
        if r<'0' || r>'9' || !adjSymbols(i,j,mat){continue}
        val,start,end:=extractNumber(i,j,mat)
        sum+=val
        for y:=start;y<=end;y++{
            mat.Set(i,y,'.')
        }
    }
    }
   return sum
}   

func adjSymbols(x,y int, mat *matrix.Matrix[rune])bool{
    var directions = [][]int{{1,0},{0,1},{-1,0},{0,-1},{1,1},{-1,-1},{1,-1},{-1,1}}
    valid:=false
    for _,dir:=range directions{
        r:=mat.Get(x+dir[0],y+dir[1])
        if r!='.' && !(r>='0' && r<='9'){
            valid=true
            break
        }
     }
    return valid 
}

func extractNumber(x,y int, mat *matrix.Matrix[rune])(val,start,end int){     

    for y>0 && mat.Get(x,y-1)>='0' && mat.Get(x,y-1)<='9'{y--}
    start=y
    for y<mat.Cols { 
        if mat.Get(x,y)<'0' || mat.Get(x,y)>'9'{break}
    val=val*10+int(mat.Get(x,y)-'0')
        y++
    }
    return val,start,y-1
}

