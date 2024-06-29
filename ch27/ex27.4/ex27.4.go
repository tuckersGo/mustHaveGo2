//ch27/ex27.4/ex27.4.go
package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static"))) // ❶
	http.ListenAndServe(":3000", nil)
}
