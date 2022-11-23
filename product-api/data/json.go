package data

import (
	"encoding/json"
	"io"
)

func ToJSON(i interface{}, wr io.Writer) error {
	e := json.NewEncoder(wr).Encode(i)
	return e
}

func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r).Decode(i)
	return e
}
