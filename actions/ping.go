package actions

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/pnkj-kmr/infra-patch-manager/service"
	"github.com/pnkj-kmr/infra-patch-manager/service/client"
)

// PingTo defines the grpc Ping method call
func (a *Action) PingTo(remote, msg string) (out *service.Remote) {
	c := a.r.Get(remote)
	return getPingResult(c, msg)
}

// PingToAll defines the grpc Ping method call to all remotes
func (a *Action) PingToAll(msg string) (out []*service.Remote) {
	start := time.Now()
	// for _, c := range a.r.GetAll() {
	// 	out = append(out, getPingResult(c, msg))
	// }
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	data := a.r.GetAll()
	wg.Add(len(data))
	for _, cc := range data {
		// concurrency with mutli host environment
		go func(r *[]*service.Remote, c *client.Client, msg string) {
			defer wg.Done()
			res := getPingResult(c, msg)
			mutex.Lock()
			*r = append(*r, res)
			mutex.Unlock()
		}(&out, cc, msg)
	}
	wg.Wait()
	log.Println("PING: receieved response with data -", len(out), "T:", time.Since(start))
	return
}

func getPingResult(c *client.Client, msg string) (out *service.Remote) {
	log.Println(c.Remote.Name, "PING: sending request with data -", msg)
	res, err := c.Ping(msg)
	var ok bool
	var serr string
	if strings.EqualFold(res, "PONG") && err == nil {
		ok = true
	}
	if err != nil {
		serr = err.Error()
	}
	out = &service.Remote{
		Name:    c.Remote.Name,
		Address: c.Remote.Address,
		Apps:    c.Remote.Apps,
		Status:  service.Status{Ok: ok, Err: serr},
	}
	log.Println(c.Remote.Name, "PING: receieved response with data -", res)
	return
}
