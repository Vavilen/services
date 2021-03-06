package main

import (
	"net/http"
	"github.com/services/scanner/handler"
	"github.com/services/scanner/config"
	"fmt"
	"os"
	"github.com/gorilla/handlers"
	"github.com/services/scanner/scan"
)


func main() {

	config.LoadConfiguration("config.json")
	_, err := scan.NewBTCDClient(config.Config.BTCD.User, config.Config.BTCD.Pass)
	if err != nil {
		fmt.Printf("Can't connect btcd, error: ", err)
		os.Exit(1)
	} else {
		fmt.Println("Connect to btcd is established")
	}
	startServer()

}

func startServer() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static/dist")))
	mux.Handle("/getaddrs", handler.AddressHandler)
	mux.Handle("/getaddr", handler.GetAddressHandler)
	mux.Handle("/newaddrs", handler.AddAddressHandler)
	mux.Handle("/scanrange", handler.DiapasonHandler)
	mux.Handle("/scanmin", handler.MinScanHandler)
	mux.Handle("/scanmax", handler.MaxScanHandler)
	mux.Handle("/scanfar", handler.FarScanHandler)
	mux.Handle("/scanshort", handler.ShortScanHandler)
	http.ListenAndServe(":7755", handlers.CORS()(mux))

}








