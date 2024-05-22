package main

import (
	"encoding/json"
	"fmt"
)

var inputData = `
{
  "row1": "test1",
  "row2": "test2",
  "row3": {
    "row31": ["test33", "test33", "test33"],
    "row32": "test3"
  }
}
`

var inputMask = []string{"row1", "row2", "row3.row31"}

func main() {
	parsedData := make(map[string]interface{})
	if err := json.Unmarshal([]byte(inputData), &parsedData); err != nil {
		panic(err)
	}
	fmt.Println(parsedData)
}
