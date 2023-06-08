package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)


func main() {
	// here a command line flag is defined --- used to add options or parameters
	// in this case user can use different csv file if they don't want to use the default one
	//  the .String specfies that the input will be coming in string format
	// So we have to send data in (-csv string) format since we have defined those things below
	csvFilename := flag.String("csv","problems.csv","a csv file in the format of question, answer")
	// The below statement is used to parse the command line given input and replace the default one.
	flag.Parse()
	// the assignment is done to compile our programme temporarily
	_ = csvFilename

	// since scvFilename is a pointer to a string. since we need value of the string we will dereference.
	file,err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open csv file: %s\n",*csvFilename))
	}
	
	// we need a reader in order to read from a file. since it is a csv file. we are using csv.NewReader to create a new reader.
	r := csv.NewReader(file)

	// Here we will get a 2d slice. with each line consisting of a row, and each of the element will be automatically generated by comma separation
	lines,err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file")
	}

	fmt.Println(lines)
	problems := parseLines(lines)
	fmt.Println(problems)
  
	marks := 0
	// here we are ranging through the slice of problem struct and we are printing the question and waiting for the users answer
	// the scanf have a newline character in order for the computer to understand the newline that automatically gets get applied when we press enter
	for i,p := range problems {
		fmt.Printf("problem %d : %s = ",i+1,p.q)
		var ans string
		fmt.Scanf("%s\n",&ans)
		if p.a == ans {
			marks++
		}
	}
	fmt.Printf("You scored %d out of %d",marks,len(problems))
}



// here each row is received in ranging and from that the question and answer will be added to two fields of a single struct (problems)
// we are using struct as it can be used for many part of the program and it is a common composite data type which can be used in variety of instances.
func parseLines(lines [][]string) []problems {

	ret := make([]problems,len(lines))
	for i,line := range lines {
		ret[i] = problems{
			q: line[0],
			a: strings.TrimSpace(line[1]), // in order to trim spaces if it exists
		}
	}

	return ret

}


type problems struct {
	q string
	a string
}


func exit(msg string) {
	
	fmt.Println(msg)
	os.Exit(1)

}


// the packages that I have used

// encoding/csv - for creating new reader of csv file type
// flag - used to pass a particular type of data from command line
// os - for reading a file or to create a reader
// strings - in order to trim space when extracting values