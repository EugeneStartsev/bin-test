package main

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type Balancer struct {
	m       sync.RWMutex
	servers []*httpServer
	current int
}

func newBalancer(servers []*httpServer) *Balancer {
	return &Balancer{
		servers: servers,
	}
}

func (b *Balancer) getNextServer() *httpServer {
	b.m.RLock()
	defer b.m.RUnlock()

	server := b.servers[b.current]
	b.current = (b.current + 1) % len(b.servers)
	return server
}

func (b *Balancer) handleGetByBalancer(ctx *gin.Context) {
	server := b.getNextServer()
	server.handleGetBin(ctx)
}
