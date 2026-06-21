package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(v any) {
	println("STRUCT______________________________________________________________________________________")
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(string(data))
}
