package main
import ("fmt")

func main() {
  var i int
  var name string
  var f float32

  fmt.Print("Type a number: ")
  fmt.Scan(&i)

  fmt.Print("Enter a string: ")
  fmt.Scan(&name)

  fmt.Print("Enter a float number: ")
  fmt.Scan(&f)

  fmt.Println("Your number is:", i)
  fmt.Println("Your name is:", name)
  fmt.Println("Your float number is:", f)
}
