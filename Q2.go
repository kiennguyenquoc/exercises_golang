// 1. Create a simple loop with the for construct. Make it loop 10 times and print out
//   the loop counter with the fmt package.
// 2. Rewrite the loop from 1. to use goto . The keyword for may not be used.
// 3. Rewrite the loop again so that it fills an array and then prints that array to the
//   screen

package main
import "fmt"

func main() {
  // for i := 0; i < 10; i++ {
  //  fmt.Println(i)
  // }
  i := 0
Loop:
  fmt.Printf("%d\n", i)
  if i < 10 {
    i++
    goto Loop
  }
}
