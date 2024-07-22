// package main // inisiasi titik masuk program go

// import (
// 	"fmt"      // import package fmt (format I/O ex: mencetak ke konsol atau write)
// 	"net/http" // import package http (mengakses web)
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // inisiasi fungsi handler
// 		w.Header().Set("Access-Control-Allow-Origin", "*") // Header CORS
// 		fmt.Fprintf(w, "Hello, World!")                    // mengisi objek w dengan string
// 	})

// 	http.ListenAndServe(":8080", nil) // inisiasi port
// }

package main
