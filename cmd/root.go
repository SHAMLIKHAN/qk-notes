package cmd

import (
	"log"
	"qk-note/api"
	"qk-note/core"
)

// Begin : Beginning of the app
func Begin() {
	db, err := prepareDatabase()
	if err != nil {
		log.Println("App : Database connection failed!")
		addr := getServerAddr()
		Router(addr, err)
	} else {
		app := api.App{
			QK: core.GetCapsule(db),
		}
		app.Serve(getServerAddr())
	}
}
