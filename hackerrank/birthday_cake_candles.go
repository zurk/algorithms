// https://www.hackerrank.com/challenges/birthday-cake-candles/problem

package main
import "fmt"

func main() {
    var array_size, cur_num int
    counter := make(map[int]int)
    fmt.Scanf("%v\n", &array_size)

    max := 0
    for i := 0; i < array_size; i++ {
        fmt.Scanf("%v", &cur_num)
        if max < cur_num {
            max = cur_num
        }
        counter[cur_num] += 1
    }

    fmt.Println(counter[max])
}