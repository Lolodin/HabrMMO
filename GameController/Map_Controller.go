package GameController

import (
	"encoding/json"
	"exampleMMO/Chunk"
	"fmt"
	"net/http"
)

func Map_Handler(w http.ResponseWriter, r *http.Request) {
	c := Chunk.NewChunk(Chunk.Coordinate{1, 1})
	js, e := json.Marshal(c)
	if e != nil {
		fmt.Println(e.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
