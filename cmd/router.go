package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"qk-note/consts"
)

// Router : Router in case of database failure
func Router(addr string, err error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Oops! Database connection failed! \n %s", err)
	})
	http.ListenAndServe(addr, nil)
}

func getServerAddr() string {
	port, ok := os.LookupEnv(consts.SERVER)
	if !ok {
		log.Println("App : SERVER environment variable required but not set")
		os.Exit(1)
	}
	addr := ":" + port
	return addr
}
