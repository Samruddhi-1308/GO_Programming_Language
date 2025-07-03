//Find the Area of any rectangle

package main
import "fmt"

func main() {
	var length, width int
	fmt.Scan(&length)
    fmt.Scan(&width)
    area := length * width
    fmt.Println("Area of the rectangle is:", area)

}
