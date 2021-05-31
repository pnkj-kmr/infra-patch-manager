package task

// PingTo defines the grpc Ping method call
func (t *PatchTask) PingTo(remote, msg string) (out string) {
	c := t.r.Get(remote)
	return c.Ping(msg)
}

// PingToAll defines the grpc Ping method call to all remotes
func (t *PatchTask) PingToAll(msg string) (out map[string]string) {
	out = make(map[string]string)
	for _, c := range t.r.GetAll() {
		out[c.Remote.Name] = c.Ping(msg)
	}
	return
}
