//Storing Boxes
//Suppose we now want to find the area and perimeter of the rectangle again which has length = 11 and width = 13.
//    But now the width has changed from 13 to 15.
package main
import "fmt"

func main() {
	var length int = 11
	var width int = 13
	fmt.Println(length * width)
	fmt.Println(2 * (length + width))
	width = 15 // Replace the underscore (_) with the updated value of width.
	fmt.Println(length * width)
	fmt.Println(2 * (length + width))
}
