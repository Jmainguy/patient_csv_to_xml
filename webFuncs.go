package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, err := io.WriteString(h, strconv.FormatInt(crutime, 10))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("upload.gtpl")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = t.Execute(w, token)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		patClass := r.PostFormValue("patClass")
		fmt.Printf("PatClass is %s\n", patClass)
		buffer := convertCsvToXML(file, patClass)

		w.Header().Set("Content-Disposition", "attachment; filename=patient.xml")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		_, err = io.Copy(w, buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
