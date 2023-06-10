package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)



func main() {

	fileName := flag.String("csv","problems.csv","give the name to csv file containing problems")
	flag.Parse()

	file,err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(file)
	questionAndAnswers,err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	quiz := addToDataStructure(questionAndAnswers)
	fmt.Println(quiz)
}

type Quiz struct {
	problem string
	choice []string
	answer string
}

func addToDataStructure(questionAndAnswers [][]string) []Quiz {

	quiz := make([]Quiz,len(questionAndAnswers))
	
	for i,val := range questionAndAnswers {

		quiz[i] = Quiz{
			problem: val[0],
			choice: val[1:len(val)-1],
			answer: val[len(val)-1],
		}
	} 

	return quiz
}