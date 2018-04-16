package main

import (
    "encoding/csv"
    "os"
    "encoding/xml"
    "fmt"
    "io"
)

func main() {

    // Csv to read from
    filename := "/tmp/export.csv"
    // XML to write to
    writefilename := "/tmp/patients.xml"
    // Root of our xml doc
    v := &Import{}

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        panic(err)
    }

    // Loop through lines & turn into object
    for _, line := range lines {
        addPatient(v, line)
    }

    file, _ := os.Create(writefilename)

    xmlWriter := io.Writer(file)

    enc := xml.NewEncoder(xmlWriter)
    enc.Indent("", "  ")
    if err := enc.Encode(v); err != nil {
        fmt.Printf("error: %v\n", err)
    }
}

