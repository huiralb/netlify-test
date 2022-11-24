package main

import (
    "fmt"
    "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.URL.Path == "/" {
        fmt.Fprint(w, "<h1>Selamat datang di skillplus</h1>")
    } else if r.URL.Path == "/aboutus" {
        fmt.Fprint(w, "<h1>Skillplus adalah web tutorial seputar informasi teknologi</h1>")
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
    }
}

func main() {
    http.HandleFunc("/", handlerFunc)
    http.ListenAndServe(":3000", nil)
}
