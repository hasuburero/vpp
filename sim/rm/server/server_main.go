package server

import (
	"http/net"
	"rm/except"
)

func Start(port string) {
	server := http.Server{
		Addr: ":" + port,
	}

	// add handler
	http.HandleFunc("")

	go func(){
		err := server.ListenAndServe()
		if err != nil{
			except.Error <- err
		}
	}
}
