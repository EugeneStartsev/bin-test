package main

import (
	"bin-checker/structs"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type httpServer struct {
	id     int
	router *gin.Engine
	stor   storage
	cache  *memCache
}

func newHttpServers(count int, engine *gin.Engine, stor storage, cache *memCache) []*httpServer {
	servers := make([]*httpServer, 0, 5)

	for i := 0; i < count; i++ {
		s := httpServer{
			id:     i + 1,
			router: engine,
			stor:   stor,
			cache:  cache,
		}

		servers = append(servers, &s)
	}

	return servers
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

	binData, ok := s.cache.get(strconv.Itoa(query.Bin))
	if ok {
		c.JSON(http.StatusOK, binData)
		return
	}

	binData, err := s.stor.getBin(strconv.Itoa(query.Bin))
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if binData.Bin != "" {
		c.JSON(http.StatusOK, binData)
		return
	}

	var saveBinData structs.SaveBinData

	url := fmt.Sprintf("https://binlist.io/lookup/%d", query.Bin)

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &saveBinData)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = s.stor.saveBin(saveBinData)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	s.cache.set(binData.Bin, binData)

	c.JSON(http.StatusOK, binData)
}
