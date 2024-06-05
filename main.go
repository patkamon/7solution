package main

import (
	"fmt"

	beef "example.com/m/beef"
	decode "example.com/m/decode"
	maxTree "example.com/m/maxtree"

	"github.com/gin-gonic/gin"
)





func main(){
	 // 1. reading from hard.json
	maxTree.Tree("hard.json")
	
	// 2. reading from input 
	fmt.Println(decode.Decode("LLRR=", "")) // output = 210122
	fmt.Println(decode.Decode("==RLL", "")) // output = 000210
	fmt.Println(decode.Decode("=LLRR", "")) // output = 221012
	fmt.Println(decode.Decode("RRL=R", "")) // output = 012001


	// 3. gin server
	r := gin.Default()
    r.GET("/beef/summary", beef.HandleSummaryRequest)
    r.Run(":8080")

	
}