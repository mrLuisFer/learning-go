// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {

// 		file, err := ioutil.ReadFile("read.txt")

// 		if err != nil {
//      fmt.Print(err)
//     }

//     str := string(file)

//     // show file data
//     fmt.Println(str)
// }

package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(path string) {
  file, err := os.Open(path)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}
 
func main() {
  var path string = "text.txt"
  openFile(path)
}
