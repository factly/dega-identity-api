package main

import (
	"net/http"
	"os"

	"github.com/factly/kavach-server/util"

	"github.com/factly/kavach-server/model"
	"github.com/joho/godotenv"

	"github.com/factly/kavach-server/action"
)

func main() {
	godotenv.Load()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "6620"
	}
	port = ":" + port

	// db setup
	model.SetupDB()

	util.InitLogging()

	r := action.RegisterRoutes()
	http.ListenAndServe(port, r)
}
