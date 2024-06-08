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
