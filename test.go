package main

import (
"net/http"
)

func main() {
http.ListenAndServe("", http.NotFoundHandler())
}