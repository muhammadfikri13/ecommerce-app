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

import (
	"log"
	"os"

	"my-backend/controllers"
	"my-backend/database"
	"my-backend/middleware"
	"my-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() { // function to start the server
	port := os.Getenv("PORT") // mengakses port Get Environment
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users")) // variable to handle other routes

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
