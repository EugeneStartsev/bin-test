package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	s := newHttpServer()

	err := s.run(fmt.Sprintf(":%d", 4000))
	if err != nil {
		log.Fatal(err)
	}
}
