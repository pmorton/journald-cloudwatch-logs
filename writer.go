package main

import (
	"encoding/json"
	"os"
)

type Writer struct {
	encoder *json.Encoder
}

func NewWriter() (*Writer, error) {

	return &Writer{
		encoder: json.NewEncoder(os.Stdout),
	}, nil
}

func (w *Writer) WriteBatch(records []Record) (string, error) {
	for _, record := range records {
		w.encoder.Encode(record)
	}
	return "",nil
}
