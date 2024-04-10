package main

import (
	"bin-checker/structs"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type httpServer struct {
	router *gin.Engine
}

func newHttpServer() *httpServer {
	s := httpServer{
		router: gin.Default(),
	}

	s.router.GET("/bin-checker", s.handleGetBin)

	return &s
}

func (s *httpServer) run(listenAddr string) error {
	return s.router.Run(listenAddr)
}

func (s *httpServer) handleGetBin(c *gin.Context) {
	var query struct {
		Bin int `form:"bin"`
	}

	var binData structs.Bin2

	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if (query.Bin < 100000) || (query.Bin > 999999) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://bin-ip-checker.p.rapidapi.com/?bin=%d", query.Bin)

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("X-RapidAPI-Key", "4963f1d9afmshe0620822d03e013p1a907ejsn85b8b5ec5e91")
	req.Header.Add("X-RapidAPI-Host", "bin-ip-checker.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(body, &binData)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, binData)
}
