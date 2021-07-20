package main

import (
	"bufio"
	"fmt"
	"os"
)

func openFile(path string) []string {
  file, err := os.OpenFile(path, os.O_RDWR, 0755)
	var a []string
	
	if err != nil {
		fmt.Println(err)
  }
	
  defer file.Close()	
  scanner := bufio.NewScanner(file)
	
	fmt.Print("Opening file with names... \n")
  for scanner.Scan() {
		var textFile string = scanner.Text()
		fmt.Println(textFile)
		// a = append(a, textFile)
		//// Inverse
		a = append([]string{textFile}, a...)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }

	return a
}
 
func main() {
  var localPath string = "./src/read-names/names.txt"
 	var names = openFile(localPath)
	fmt.Println("Getting names")
	fmt.Print(names)

}
