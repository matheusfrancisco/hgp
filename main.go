package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: go run main.go <filename>")
    os.Exit(1)
  }
  //fmt.Println("Filename: ", os.Args[1])

	//filename := "10m-v2.txt"
	filename := os.Args[1]
	infoFile := ReadFileInformation(filename)
	bitset := makeBitsetAndPopulate(infoFile)

	//fmt.Println("File Information", infoFile)
	//fmt.Printf("Number of lines in %s: %d\n", filename, infoFile.NumberOfLines)
	//fmt.Printf("Bitset: %b\n", bitset)

	query := ParseQueryInput("4 79 13 80 56")
	start := time.Now()
	query.SelectedBitset(bitset)

  //This is straightforward: the count of players who matched all 5 
  //numbers does not overlap with other counts.
  //c(n,k) = n! / (k! * (n-k)!)
	countFive := query.countWinnerIntersection(5)
  // countFour includes players who matched 4 or 5 numbers.
  //Subtract 5*countFive because: 
  //A player who matched all 5 numbers contributes to countFour in 5 ways (one for each combination of 4 numbers out of 5).
  //Use the multiplier 5 because there are C(5, 4) = 5 combinations of 4 numbers from 5.
	countFour := query.countWinnerIntersection(4) - 5*countFive
  //countThree includes players who matched 3, 4, or 5 numbers.
  //Subtract: 4*countFour: A player who matched exactly 4 numbers contributes to countThree in 4 ways (one for each combination of 3 numbers out of 4).
  // 10*countFive: A player who matched all 5 numbers contributes to countThree in 10 ways (one for each combination of 3 numbers out of 5).
  // Use the multiplier 10 because there are C(5, 3) = 10 combinations of 3 numbers from 5.
	countThree := query.countWinnerIntersection(3) - 4*countFour - 10*countFive
  //countTwo includes players who matched 2, 3, 4, or 5 numbers.
  //Subtract:
  //3*countThree: A player who matched 3 numbers contributes to countTwo in 3 ways (one for each combination of 2 numbers out of 3).
  //6*countFour: A player who matched 4 numbers contributes to countTwo in 6 ways (one for each combination of 2 numbers out of 4).
  //10*countFive: A player who matched all 5 numbers contributes to countTwo in 10 ways (one for each combination of 2 numbers out of 5).
	countTwo := query.countWinnerIntersection(2) - 3*countThree - 6*countFour - 10*countFive
	elapsed := time.Since(start)
	fmt.Println("count: ", countFive, countFour, countThree, countTwo)
	fmt.Println("Time: ", elapsed.Milliseconds(), " ms")

}
