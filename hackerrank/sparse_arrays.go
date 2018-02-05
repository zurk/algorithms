// https://www.hackerrank.com/challenges/sparse-arrays/problem

package main
import "fmt"

func main() {
    var array_size int
    var s string
    counter := make(map[string]int)
    fmt.Scanf("%v\n", &array_size)

    for i := 0; i < array_size; i++ {
        fmt.Scanf("%s", &s)
        counter[s] += 1
    }

    fmt.Scanf("%v\n", &array_size)
    for i := 0; i < array_size; i++ {
        fmt.Scanf("%s", &s)
        fmt.Println(counter[s])
    }
}