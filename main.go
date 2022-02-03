package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------TESTING---------")

	// // Create a new context and add a timeout to it
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel() // The cancel should be deferred so resources are cleaned up

	// // Create the command with our context
	// cmd := exec.CommandContext(ctx, "ping", "-c 4", "-i 1", "8.8.8.8")

	// // This time we can simply use Output() to get the result.
	// out, err := cmd.CombinedOutput()

	// // We want to check the context error to see if the timeout was executed.
	// // The error returned by cmd.Output() will be OS specific based on what
	// // happens when a process is killed.
	// if ctx.Err() == context.DeadlineExceeded {
	// 	fmt.Println("Command timed out")
	// 	return
	// }

	// // If there's no context error, we know the command completed (or errored).
	// fmt.Println("Output:", string(out))
	// if err != nil {
	// 	fmt.Println("Non-zero exit code:", err)
	// }

	// newDir := "tmp/sdsd"
	// p, err := entity.CreateDirectoryIfNotExists(newDir)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("----------", p)

	// out, err := entity.ExecuteCmd("ps -a")
	// fmt.Println(string(out), err)

	// f := "x.tar.gz"
	// tar := entity.NewTar("", f, entity.C.AssetPath())
	// err := tar.Untar(entity.C.PatchPath())
	// fmt.Println(err)

	// pa, err := agent.NewPatchAgent("tmp", false)
	// fmt.Println(">>>>", pa, err)
	// out, err := pa.ListAssets()
	// fmt.Println(">>>>", out, err)
	// fmt.Println(len(out))

}
