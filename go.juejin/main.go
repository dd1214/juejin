package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/juejin", handlerHello)
	r := gin.Default()

	initRouter(r)

	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
