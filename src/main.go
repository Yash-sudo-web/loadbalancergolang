package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

type LoadBalancer struct {
	port            string
	roundRobinIndex int
	servers         []Server
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		fmt.Printf("parse server url failed, err:%v\n", err)
		os.Exit(1)
	}
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinIndex: 0,
		servers:         servers,
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	return true
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.nextServer()
	fmt.Printf("proxy to server: %s\n", targetServer.Address())
	targetServer.Serve(rw, req)
}

func (lb *LoadBalancer) nextServer() Server {
	server := lb.servers[lb.roundRobinIndex%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinIndex++
		server = lb.servers[lb.roundRobinIndex%len(lb.servers)]
	}
	lb.roundRobinIndex++

	return server
}

func main() {
	servers := []Server{
		newSimpleServer("http://youtube.com"),
		newSimpleServer("http://facebook.com"),
		newSimpleServer("http://google.com"),
	}
	lb := newLoadBalancer("8080", servers)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Load Balancer started at localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
