package _00_499

import "fmt"

//412. Fizz Buzz
func fizzBuzz(n int) []string {

    var i = 1
    var s []string
    for i <= n {

        switch {
        case i%3 == 0 && i%5 == 0:
            s = append(s, `FizzBuzz`)
        case i%3 == 0:
            s = append(s, `Fizz`)
        case i%5 == 0:
            s = append(s, `Buzz`)
        default:
            s = append(s, fmt.Sprintf("%d", i))
        }
        i++
    }

    return s

}