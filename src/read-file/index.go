package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(path string) {
  file, err := os.OpenFile(path, os.O_RDWR, 0755)

  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Print("Opening file... \n")
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }
}
 
func main() {
  var path string = "./text.txt"
  var localPath string = "./src/read-file/textlocal.txt"
  openFile(path)
  openFile(localPath)
}
