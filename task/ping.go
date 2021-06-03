package task

import (
	"log"
	"strings"

	"github.com/pnkj-kmr/patch/module/jsn"
)

// PingTo defines the grpc Ping method call
func (t *PatchTask) PingTo(remote, msg string) (out jsn.Remote) {
	c := t.r.Get(remote)
	log.Println(c.Remote.Name, "PING: sending request with data -", msg)
	res := c.Ping(msg)
	if strings.EqualFold(res, "PONG") {
		c.Remote.Status = true
	}
	out = c.Remote
	log.Println(c.Remote.Name, "PING: receieved response with data -", res)
	return
}

// PingToAll defines the grpc Ping method call to all remotes
func (t *PatchTask) PingToAll(msg string) (out []jsn.Remote) {
	for _, c := range t.r.GetAll() {
		log.Println(c.Remote.Name, "PING: sending request with data -", msg)
		res := c.Ping(msg)
		if strings.EqualFold(res, "PONG") {
			c.Remote.Status = true
		}
		out = append(out, c.Remote)
		log.Println(c.Remote.Name, "PING: receieved response with data -", res)
	}
	return
}
