// https://www.hackerrank.com/challenges/simple-array-sum/problem

package main
import "fmt"

func main() {
    var array_size, cur_num int
    fmt.Scanf("%v\n", &array_size)

    sum := 0
	for i := 0; i < array_size; i++ {
    	fmt.Scanf("%v", &cur_num)
    	sum += cur_num
	}

    fmt.Println(sum)
}