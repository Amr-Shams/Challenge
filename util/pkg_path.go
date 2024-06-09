package util
import(
    "fmt"
	"path/filepath"
	"runtime"
    "os"
)
// find the path of the current executable
func RootPath() (string,error){
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        return "", fmt.Errorf("No caller information")
    }
    return filepath.Dir(filename),nil
}
func DaysPath(day int) (string,error){
   root,err:=RootPath()
   if err!=nil{
       return "",err
   }
   dayPath:=fmt.Sprintf("day%d",day)
   fullPath:=filepath.Join(root[:len(root)-len("util")],"challenge",dayPath)
   if _, err := os.Stat(fullPath); os.IsNotExist(err) {
    os.Mkdir(fullPath, os.ModePerm)
   }
   return fullPath,nil
}
