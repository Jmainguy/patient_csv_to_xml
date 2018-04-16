package main

import (
    "encoding/csv"
    "io"
    "net/http"
    "mime/multipart"
    "bytes"
    "fmt"
    "encoding/xml"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func convertCsvToXml(f multipart.File) (buffer *bytes.Buffer){
    // Root of our xml doc
    v := &Import{}

    // Read CSV File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }

    // Loop through CSV lines & turn into object
    for _, line := range lines {
        addPatient(v, line)
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

    http.ListenAndServe(":8000", nil)

}
