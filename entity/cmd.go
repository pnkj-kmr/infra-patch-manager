package entity

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

const timeout time.Duration = 60 * time.Second

// ExecuteCmd - helps to execute cmd at os level
func ExecuteCmd(in string) (out []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	inputs := strings.Fields(in)
	cmd := exec.CommandContext(ctx, inputs[0], inputs[1:]...)

	result, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		// fmt.Println("Command timed out")
		return nil, ctx.Err()
	}
	if err != nil {
		return
	}
	// fmt.Println("Output:", string(result))
	return result, err
}
