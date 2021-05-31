package task

// ApplyTo defines the grpc Ping method call
func (t *PatchTask) ApplyTo(remote string) (out string) {
	c := t.r.Get(remote)
	var targets []string
	for _, app := range c.Remote.Apps {
		targets = append(targets, app.Path)
	}
	return c.Ping("") // TODO
}

// ApplyToAll defines the grpc Ping method call to all remotes
func (t *PatchTask) ApplyToAll() (out map[string]string) {
	out = make(map[string]string)
	for _, c := range t.r.GetAll() {
		out[c.Remote.Name] = c.Ping("")
	}
	return
}
