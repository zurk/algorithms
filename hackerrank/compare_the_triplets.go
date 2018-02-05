// https://www.hackerrank.com/challenges/compare-the-triplets/problem

package main
import "fmt"

func main() {
    var a1, a2, a3, b1, b2, b3 int
    fmt.Scanf("%v %v %v\n%v %v %v", &a1, &a2, &a3, &b1, &b2, &b3)

    rating_a := 0
    rating_b := 0

    if a1 > b1 {
    	rating_a += 1 
    } else if a1 < b1 {
    	rating_b += 1
    }

    if a2 > b2 {
    	rating_a += 1 
    } else if a2 < b2 {
    	rating_b += 1
    }

    if a3 > b3 {
    	rating_a += 1 
    } else if a3 < b3 {
    	rating_b += 1
    }

    fmt.Println(rating_a, rating_b)
}