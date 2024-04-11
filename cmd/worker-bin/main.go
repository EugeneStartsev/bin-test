package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	servers := newHttpServers(5, g)
	balancer := newBalancer(servers)

	g.GET("/bin-checker", balancer.handleGetByBalancer)

	go func() {
		err := g.Run(":4444")
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		time.Sleep(30 * time.Millisecond)
		go startReq(wg)
	}

	wg.Wait()

	signch := make(chan os.Signal)
	signal.Notify(signch, os.Interrupt)
	<-signch
}

func startReq(wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("http://localhost:4444/bin-checker?bin=%d", 518683)

	req, _ := http.NewRequest("GET", url, nil)

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
