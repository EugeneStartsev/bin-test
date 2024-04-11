package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	s := newHttpServer()

	go func() {
		err := s.run(fmt.Sprintf(":%d", 4444))
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
