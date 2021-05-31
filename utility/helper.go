package utility

import "strings"

// Ping helps to validate the ping-pong echo
// input as "PING" output will be "PONG" else ""
func Ping(in string) (out string) {
	if ok := strings.EqualFold(in, "PING"); ok {
		out = "PONG"
	}
	return
}
