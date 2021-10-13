package main

import "fmt"
import "../gorez"

func main() {
	fmt.Println("hello")
	file := gorez.NewREZFile("../../../../../../temp/Game.rez")
	err := file.Open()
	if err != nil {
		fmt.Printf("Erro %v\n", err)
	}
	fmt.Println(file.Filename())
	fmt.Println(file.Size())
	header := file.Header()
	fmt.Printf("FileFormatVersion %v\n", header.FileFormatVersion)
	fmt.Printf("RootDirPos %v\n", header.RootDirPos)
	fmt.Printf("RootDirSize %v\n", header.RootDirSize)
	fmt.Printf("RootDirTime %v\n", header.RootDirTime)
	fmt.Printf("NextWritePos %v\n", header.NextWritePos)
	fmt.Printf("Time %v\n", header.Time)
	fmt.Printf("LargestKeyAry %v\n", header.LargestKeyAry)
	fmt.Printf("LargestDirNameSize %v\n", header.LargestDirNameSize)
	fmt.Printf("LargestRezNameSize %v\n", header.LargestRezNameSize)
	fmt.Printf("LargestCommentSize %v\n", header.LargestCommentSize)
	fmt.Printf("IsSorted %v\n", header.IsSorted)

	err2 := file.Read()
	if err2 != nil {
		fmt.Printf("Erro %v\n", err2)
	}	

	file.Close()

}