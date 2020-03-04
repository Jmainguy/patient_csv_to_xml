package main

import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"mime/multipart"
	"net/http"
)

func convertCsvToXML(f multipart.File, patClass string) (buffer *bytes.Buffer) {
	// Root of our xml doc
	v := &Import{}

	// Read CSV File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	// Loop through CSV lines & turn into object
	for _, line := range lines {
		addPatient(v, line, patClass)
	}

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	err = enc.Encode(&v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	f.Close()
	buffer = &buf
	return buffer
}

func main() {

	http.HandleFunc("/", upload)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}

}
