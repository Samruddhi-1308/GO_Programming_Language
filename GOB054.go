//Taking user input
//You have already learned that fmt.Print() is used to output (print) values.
//Now we will use fmt.scan() to get user input.



package main
import "fmt"

func main() {
	var a int
    fmt.Scan(&a)
    fmt.Println("Your number is: ", a)

}