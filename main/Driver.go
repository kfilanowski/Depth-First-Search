// Package main - The main driver of the program.
//
// Author - Kevin Filanowski
//
// Version - October 2018
package main

import (
   "errors"
   "KF_Project2_DFS/graph"
   "fmt"
   "os"
   "bufio"
)

// readFile - Reads the file input
// May throw errors if there is invalid input or the file cannot be read.
func readFile(fileName string) ([]string, error) {
   // result variable will hold the contents of the file.
   var result []string

   // Exit if the input file could not be read.
	inFile, err := os.Open(fileName)
	if err != nil { return result, errors.New("input file cannot be read") }

   // Set up a scanner to read the file.
   scanner := bufio.NewScanner(inFile)
   scanner.Split(bufio.ScanWords)

   // Counter will catch vertices mapping to nowhere.
   counter := 0

   // Read file into variable result
   for scanner.Scan() {
      result = append(result,scanner.Text())
      counter++
   }
   // Error checking
   if err := scanner.Err(); err != nil { return result, err }
   // Close the file.
   inFile.Close()
   // Ensure there are enough elements in the file
   if counter % 2 != 0 {
      return result, errors.New("A vertex in the file maps to nowhere." +
                           " Please check input.")
   }
   return result, nil
}

func main() {
   // Exit if there is not the correct number of arguments.
	if len(os.Args) != 2 {
	  fmt.Println("Usage: go run Driver.go <input_file>")
	  os.Exit(2)
	}

   // Read in the file and retrieve the data.
   result, err := readFile(os.Args[1])
   if (err != nil) {
      fmt.Println(err)
      os.Exit(1)
   }

   // Create a graph.
   g := graph.Graph{}

   // Run the program.
   err = g.Run(result)
   if (err != nil) {
      fmt.Println(err)
      os.Exit(1)
   }
}
