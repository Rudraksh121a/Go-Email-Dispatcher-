package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	for _, records := range records[1:] {
		fmt.Println(records)
	}
	return nil

}
