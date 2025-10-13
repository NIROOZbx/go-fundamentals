package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Type something: ")
    input, _ := reader.ReadString('\n')
    fmt.Println("You typed:", input)

    fmt.Print("Type new: ")
    input2, _ := reader.ReadString('\n')
    fmt.Println("You typed:", input2)
	fmt.Print(input+input2)
}
