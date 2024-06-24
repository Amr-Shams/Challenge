package day3 
import(
    "fmt"
    "github.com/spf13/cobra"
    "github.com/Amr-Shams/aoc2024/challenge"
   "github.com/Amr-Shams/aoc2024/util/matrix"

)

func bCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "b",
        Short: "Day 3, Problem b",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n",partB( challenge.FromFile()))
        },
    }
}

func partB(challenge *challenge.Input) int {
    mat:=matrix.CreateMatrixFrom[rune](challenge,matrix.ConvertToRune)
    var sum int 
    for i:=0; i<mat.Rows; i++{
        for j:=0; j<mat.Cols; j++{
            r:=mat.Get(i,j)
            if r!='*'{continue}
              sum+=neighborNumbers(mat,i,j)
          }
    }
    return sum
}
func neighborNumbers(mat *matrix.Matrix[rune], i,j int) int{
   adj:=[]int{}
   seen:=matrix.NewMatrix[bool](mat.Rows,mat.Cols)
    for _,d := range [] struct{dx,dy int}{
    {0,1},
    {0,-1},
    {1,0},
    {-1,0},
    {1,1},
    {1,-1},
    {-1,1},
    {-1,-1},
    }{
        x,y:=i+d.dx,j+d.dy
        r:=mat.Get(x,y)
        if r>='0' && r<='9'&& seen.Get(x,y)==false{
            val,start,end:=extractNumber(x,y,mat) 
            for sy:=start; sy<=end; sy++{
                seen.Set(x,sy,true)
            }
            adj=append(adj,val)
        }
  }
  if len(adj)!=2{
      return 0
  }
  return adj[0]*adj[1]
}
