package main

import (
	"fmt"
	"bufio"
	"os"
	
)

func main(){
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)

	for scan.Scan(){
		words[scan.Text()]++
	}		

	// fmt.Println(len(words),"unique words")
	// fmt.Println(words)



	type wordCount struct{
		word string
		count int  

	}


	var wordList []wordCount
	for word,count := range words{
		wordList  = append(wordList,wordCount{word,count})
	}

	fmt.Println(wordList)

}

