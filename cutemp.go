// cutemp.go was designed to be invoked using curl or http
// You can cf push cutemp.go to TAS and then call it per the examples shown below:
//
// Example: cutemp is running locally (go run cutemp.go)
//
// curl 127.0.0.1:3000/user1-fact.apps.ourpcf.com/1000
//
// reponse:   [[2,0.069207]]    <--- json format
//
// curl 127.0.0.1:3000/user1-fact.apps.ourpcf.com/1000
//
// reponse:   [[3,0.059087]]    <--- json format
//

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
        "time"
)

const ( version = "v1.0.0" )

func headers(w http.ResponseWriter, req *http.Request) {
    log.Println("Requested Header API")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
            log.Println(fmt.Sprintf("%v: %v", name, h))
        }
    }
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}


func main() {

         var xa int = 0

	log.Println("Starting Factorial Application...")
        http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, version)
                                                                                   log.Println("Requested Version API")
                                                                                 })
        http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200)
                                                                                  log.Println("Requested Health API")
                                                                                })
        http.HandleFunc("/header", headers )
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            enableCors(&w)

                    pp := r.URL.Path[1:len(r.URL.Path)]
                    
                    // fmt.Fprint(w,"Veio com: >>",pp,"<< tamanho: >>",len(pp),"<<\n ")

                    if (len(pp)<12) { pp = "www.google.com" } else {  xa++; }

                    // fmt.Fprint(w,"Virou: >>",pp,"<< tamanho: >>",len(pp),"<<\n ")


         start := time.Now()
         
         result, err := http.Get("http://"+pp)
         
         if err != nil {
                 fmt.Println(err)
                 // os.Exit(1)
          } else { 
                    defer result.Body.Close()
                    elapsed := time.Since(start).Seconds() * 1000.0
                    fmt.Printf("http.Get to %s took %v milliseconds \n", pp, elapsed)
                    quero := fmt.Sprintf("%f", elapsed)
                    fmt.Fprint(w,"[[",xa,",",quero,"]]")

                 }

       	})

	var port string
	port = os.Getenv("PORT")
        if ( len(port) == 0 ) {
		port = "3000"
	}

        log.Println("Using port"+port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

	s := http.Server{Addr: ":" + port }

	go func() { log.Fatal(s.ListenAndServe()) }()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
