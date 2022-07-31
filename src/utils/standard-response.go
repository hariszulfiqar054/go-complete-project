package utils

import (
	"encoding/json"
	"log"
)

type Data = interface{}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func New(s bool, m string, d Data) Response {
	return Response{s, m, d}
}

func (r *Response) MustMarshal() []byte {
	j, err := json.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}

	return j
}
