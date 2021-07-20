package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(path string) {
  // O_RDWR: Flags to OpenFile wrapping those of the underlying system. 
  // Not all flags may be implemented on a given system
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
  /*
    for some reason it takes the path from the root
  */
  var path string = "./text.txt"
  var localPath string = "./src/read-file/textlocal.txt"
  openFile(path)
  openFile(localPath)
}
