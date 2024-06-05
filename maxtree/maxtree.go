package maxtree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// def findMax(input, x, y):
//     if mem.get("x"+str(x)+"y"+str(y)) != None:
//         return mem["x"+str(x)+"y"+str(y)]
//     if y == len(input)-1:
//         return input[y][x]

//     mem["x"+str(x)+"y"+str(y+1)] = findMax(input, x,  y+1)
//     mem["x"+str(x+1)+"y"+str(y+1)] = findMax(input, x+1,  y+1)
//     return input[y][x] + max(mem["x"+str(x)+"y"+str(y+1)], mem["x"+str(x+1)+"y"+str(y+1)])

// print(findMax(input,0,0))
    

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func getName(x int, y int) string {
	return "x"+strconv.Itoa(x) + "y"+strconv.Itoa(y)
}

func findMax(mem map[string]int, input [][]int, x int, y int) int{
	name := getName(x,y)
	if mem[name] != 0 {
		return mem[name] 
	}else if y == len(input) -1{
		return input[y][x]
	}

	leftChildName := getName(x,y+1)
	rightChildName := getName(x+1,y+1)
	// mem[leftChildName] = new(int)
	mem[leftChildName] = findMax(mem, input, x, y+1) 
	// mem[rightChildName] = new(int)
	mem[rightChildName] = findMax(mem, input, x+1, y+1) 

	return input[y][x] + max(mem[rightChildName], mem[leftChildName])
}




func Tree(filename string){
	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Read the file contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	// Define a variable to hold the nested array
	var nestedArray [][]int

	// Unmarshal the JSON data into the variable
	err = json.Unmarshal(bytes, &nestedArray)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}

	mem := make(map[string]int)

	ans := findMax(mem, nestedArray, 0,0 )

	fmt.Println("ans: ",ans)
}
