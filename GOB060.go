// Cubes and Squares

package main
import "fmt"

func main() {
	var a int
	var square int
	var cube int

	fmt.Scan(&a)
    square = a * a
    cube = square * a
    fmt.Println("Square is:", square)
    fmt.Println("Cube is:", cube)
}
