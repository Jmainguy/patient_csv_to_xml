package main

import (
    "fmt"
    "net/http"
    "io"
    "crypto/md5"
    "time"
    "strconv"
    "html/template"
)

func upload(w http.ResponseWriter, r *http.Request) {
       fmt.Println("method:", r.Method)
       if r.Method == "GET" {
           crutime := time.Now().Unix()
           h := md5.New()
           io.WriteString(h, strconv.FormatInt(crutime, 10))
           token := fmt.Sprintf("%x", h.Sum(nil))

           t, _ := template.ParseFiles("upload.gtpl")
           t.Execute(w, token)
       } else {
           r.ParseMultipartForm(32 << 20)
           file, _, err := r.FormFile("uploadfile")
           if err != nil {
               fmt.Println(err)
               return
           }
           buffer := convertCsvToXml(file)

           w.Header().Set("Content-Disposition", "attachment; filename=patient.xml")
           w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
           io.Copy(w, buffer)
       }
}
