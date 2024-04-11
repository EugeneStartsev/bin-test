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

	var binData structs.BinData

	if err := c.ShouldBindQuery(&query); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if (query.Bin < 100000) || (query.Bin > 99999999) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://binlist.io/lookup/%d", query.Bin)

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(body, &binData)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(binData)

	//c.JSON(http.StatusOK, binData)
}
