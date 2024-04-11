package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	for i := 0; i < 1000; i++ {
		url := fmt.Sprintf("http://localhost:4444/bin-checker?bin=%d", 518683)

		req, _ := http.NewRequest("GET", url, nil)

		_, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
	}
}
