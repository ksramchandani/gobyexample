// package main

// import (
// 	_ "expvar"
// 	"fmt"
// 	"net"
// 	"net/http"
// )

// func main() {
// 	// Create socket
// 	sock, err := net.Listen("tcp", "localhost:8123")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	go func() {
// 		fmt.Println("expvar http service listening on port 8123")
// 		http.Serve(sock, nil)
// 	}()

// }

package main

import (
	"expvar"
	"fmt"
	"net/http"
)

var (
	openPorts    = expvar.NewInt("open_ports")
	stringMetric = expvar.NewString("kamal")
)

func main() {
	forever := make(chan bool)

	// Start a local http server for expvar in a separate goroutine
	go func() {
		http.ListenAndServe(":8123", http.DefaultServeMux)
	}()

	openPorts.Add(1)
	openPorts.Add(1)
	openPorts.Add(1)
	openPorts.Add(1)
	openPorts.Add(1)
	openPorts.Add(1)

	stringMetric.Set("kr")
	fmt.Println(stringMetric.String())

	<-forever
}
