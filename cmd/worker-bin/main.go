package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	dbConn := flag.String("db",
		"host=localhost dbname=bin user=bin password=bin sslmode=disable",
		"database connection string")
	httpPort := flag.Int("http-port", 4444, "HTTP API port")
	flag.Parse()

	stor, err := newDb(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	cache := newMemCache()

	err = cache.recoverFromPostgres(stor)

	servers := newHttpServers(5, g, stor, cache)
	balancer := newBalancer(servers)

	g.GET("/bin-checker", balancer.handleGetByBalancer)

	go func() {
		err = g.Run(fmt.Sprintf(":%d", *httpPort))
		if err != nil {
			log.Fatal(err)
		}
	}()

	signch := make(chan os.Signal)
	signal.Notify(signch, os.Interrupt)
	<-signch
}
