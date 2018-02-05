// https://www.hackerrank.com/challenges/arrays-ds/problem

package main
import "fmt"

func main() {
    var slice_size, cur_num int
    fmt.Scanf("%v\n", &slice_size)
    slice := make([]int, 0, slice_size)

    for i := 0; i < slice_size; i++ {
        fmt.Scanf("%v", &cur_num)
        slice = append(slice, cur_num)
    }
    for i := len(slice)-1; i >= 0; i-- {
        fmt.Printf("%v ", slice[i])
    }
    fmt.Println()
}