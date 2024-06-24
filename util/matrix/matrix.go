package matrix 
import(
     "fmt"
     "golang.org/x/exp/constraints"
     "github.com/Amr-Shams/aoc2024/challenge"
)

type data interface {
    constraints.Integer | constraints.Float | rune | string | bool 
}

type Matrix[T data] struct {
    data [][]T
    Rows int
    Cols int    
}

func NewMatrix[T data](rows, cols int) *Matrix[T] {
    m := &Matrix[T]{Rows: rows, Cols: cols}
    m.data = make([][]T, rows)
    for i := range m.data {
        m.data[i] = make([]T, cols)
    }
    return m
}

func (m *Matrix[T]) Set(row, col int, val T) {
    m.data[row][col] = val
}

func (m *Matrix[T]) Get(row, col int) T {
    val,_ := m.getIndex(row, col)
    return val
}
func (m *Matrix[T]) getIndex(row, col int) (T,error) {
 if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
        return m.data[0][0],fmt.Errorf("Index out of range") 
    }
    return m.data[row][col],nil
}

func (m *Matrix[T]) String() string {
    var s string
    for i := 0; i < m.Rows; i++ {
      s += fmt.Sprintf("%v\n", m.data[i])
    }
    return s
}
func CreateMatrixFrom[T data](input *challenge.Input,convert func(rune) T) *Matrix[T] {
    lines := input.LineSlice()
    m := NewMatrix[T](len(lines), len(lines[0]))
    for i, line := range lines {
        for j, r := range line {
            m.Set(i, j, convert(r))
        }
    }
    return m
}

func (m *Matrix[T])Walkneighbors(row, col int, visit func(T)){
    var dir=[]struct{dx,dy int}{
        {0,1},
        {0,-1},
        {1,0},
        {-1,0},
        {1,1},
        {1,-1},
        {-1,1},
        {-1,-1},
    }
    for _,d:=range dir{
        r,c:=row+d.dx,col+d.dy
        if val,err:=m.getIndex(r,c);err==nil{ 
            visit(val)
        }
    }
}



