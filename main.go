package main

import (
	"encoding/json"
	"fmt"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func main() {
	buff := duit_props.NewTranslateTransform().SetOffset(duit_props.NewOffset(10, 10))

	j, err := json.Marshal(buff)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(j))
}