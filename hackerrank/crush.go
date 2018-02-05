// Array Manipulation
// https://www.hackerrank.com/challenges/crush/problem

package main
import "fmt"
import "math"

func main() {
    var slice_size, command_num int
    fmt.Scanf("%v %v\n", &slice_size, &command_num)

    slice := make([]int, slice_size)

    var first, last, inc int
    for i := 0; i < command_num; i++ {
        fmt.Scanf("%v %v %v", &first, &last, &inc)
        for j := first; j <= last; j++ {
            slice[j] += inc
        } 
    }
    fmt.Println(math.Max(slice))
}
