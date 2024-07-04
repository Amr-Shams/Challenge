package util 
import "strconv"
import "log"
func Atoi(s string) (int) {
    i, err := strconv.Atoi(s)
    if err != nil {
        log.Fatal(err)
        return 0
    }
    return i
}

func Atof(s string) (float64) {
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        log.Fatal(err)
        return 0
    }
    return f
}
// create a func that convert a list of strings to a list of intergers or a single string into a list of an intergers
func AtoiList(s ...string) ([]int) {
    var list []int
    for _, i := range s {
        list = append(list, Atoi(i))
    }
    return list
}
