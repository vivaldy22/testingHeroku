package main

import (
	"github.com/gorilla/mux"
	"github.com/vivaldy22/testingHeroku/viper"
	"log"
	"net/http"
)

func main() {
	//g := gin.Default()
	//
	//g.GET("/", func(ctx *gin.Context) {
	//	ctx.Writer.Write([]byte("Hai"))
	//})
	//
	//g.Run()

	port := viper.ViperGetEnv("PORT", "8080")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ga cair cair euy :D"))
	})
	r.HandleFunc("/cair", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Masih belum cair gan"))
	})

	log.Println("Starting Server...")

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
