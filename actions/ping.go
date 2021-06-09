package actions

import (
	"log"
	"strings"

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
	for _, c := range a.r.GetAll() {
		out = append(out, getPingResult(c, msg))
	}
	log.Println("PING: receieved response with data -", len(out))
	return
}

func getPingResult(c *client.Client, msg string) (out *service.Remote) {
	log.Println(c.Remote.Name, "PING: sending request with data -", msg)
	res, err := c.Ping(msg)
	var ok bool
	if strings.EqualFold(res, "PONG") && err == nil {
		ok = true
	}
	out = &service.Remote{
		Name:    c.Remote.Name,
		Address: c.Remote.Address,
		Apps:    c.Remote.Apps,
		Status:  service.Status{Ok: ok, Err: err.Error()},
	}
	log.Println(c.Remote.Name, "PING: receieved response with data -", res)
	return
}
