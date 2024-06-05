package beef

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Beef struct {
	Beef *map[string]*int  `json:"beef"`
}

type Chunk struct {
	StartIndex int
	EndIndex   int
}

func splitmem(data []byte, numGoRoutime int) []Chunk {
	chunk := make([]Chunk, numGoRoutime)
	num := len(data) / numGoRoutime
	chunk[0].StartIndex = 0
	for i := 1; i < numGoRoutime; i++ {
		start := false
		for j := i * num; j < i*num+30; j++ {
			if !start && (data[j] == ' ' || data[j] == '.' || data[j] == ',' || data[j] == '\n') {
				chunk[i-1].EndIndex = j
				start = true
			} else if start && !(data[j] == ' ' || data[j] == '.' || data[j] == ',' || data[j] == '\n') {
				chunk[i].StartIndex = j
				break
			}
		}
	}
	return chunk
}

func parallel(pipeline chan (map[string]*int), data []byte, startIndex int, endIndex int) {
	start := false
	prev := endIndex
	beef := ""
	collection := make(map[string]*int)
	for i := startIndex; i < endIndex; i++ {
		if start && (data[i] == ' ' || data[i] == '.' || data[i] == ',' || data[i] == '\n') {
			beef = string(bytes.ToLower(data[prev:i]))
			if collection[beef] == nil {
				collection[beef] = new(int)
				*collection[beef] = 1
			} else {
				*collection[beef] += 1
			}
			start = false
		} else if !start && !(data[i] == ' ' || data[i] == '.' || data[i] == ',' || data[i] == '\n') {
			start = true
			prev = i
		}
	}
	pipeline <- collection
}

// Min returns the minimum of x and y.
func Min(x, y int) int {
	return y ^ ((x ^ y) & ((x - y) >> 63))
}

func CustomMmap(input string) (*Beef, error){
	response, err := http.Get(input) //use package "net/http"

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Can't get input url")
	}

	defer response.Body.Close()

	// Copy data from the response to standard output
	data, err1 := io.ReadAll(response.Body) //use package "io" and "os"
	if err != nil {
		fmt.Println(err1)
		return nil, fmt.Errorf("Can't read body")
	}
	maxGoroutime := Min(runtime.NumCPU(), runtime.GOMAXPROCS(0))

	chunks := splitmem(data, maxGoroutime)

	pipeline := make(chan map[string]*int)
	for i := 0; i < maxGoroutime; i++ {
		go parallel(pipeline, data, chunks[i].StartIndex, chunks[i].EndIndex)
	}

	ans := &Beef{}
	collection := make(map[string]*int)
	ans.Beef = &collection

	for i := 0; i < maxGoroutime; i++ {
		beefs := <-pipeline
		for b, value := range beefs {
			if collection[b] == nil {
				collection[b] = value
			} else {
				*collection[b] = *collection[b] + *value
			}
		}
	}
	return ans, nil
}

func HandleSummaryRequest(c *gin.Context) {
    // Perform any necessary logic here
	summary, err := CustomMmap("https://baconipsum.com/api/?type=meat-and-filler&paras=150&format=text")
	if err != nil{
		panic("something wrong")
	}

    c.JSON(http.StatusOK, summary)
}
