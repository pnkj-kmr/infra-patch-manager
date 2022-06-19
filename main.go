package main

import (
	"fmt"
	"time"
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

	// a := [...]int{1, 2, 3}
	// fmt.Println(&a[1])
	// // a[1] = 33
	// // fmt.Println(a)

	// var b []int
	// fmt.Println(b == nil)
	// // b[1] = 1
	// fmt.Println(b)

	// // // b := append(a[:], a...)
	// // // fmt.Println(b)
	// var m map[int]bool
	// fmt.Println(m == nil)

	// const (
	// 	A = 1
	// 	B = 3
	// 	C = 6
	// )
	// a := [...]int{A: 2, B: 4, C: 44}
	// fmt.Println(a)

	// err := agent.BackupRollback()
	// fmt.Println(err)

	// t := entity.NewTar("", "ROLLBACK__1645727283.tar.gz", "resources/assets")
	// err := t.Untar("resources/patch")
	// fmt.Println(err)

	// fmt.Println(strings.Compare("XX", "XXX"))
	// format := "%v\t%v\t%v\t%v\t%v\t\n"
	// tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	// fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	// fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	// fmt.Fprintf(tw, format, "	-----", "------", "----dfkdjfkdjfkdjfkdfj-", "----", "------")
	// // for _, t := range tracks {
	// // 	fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	// // }
	// tw.Flush() // calculate column widths and print table

	st := time.Now()
	out := make(chan string, 5)
	// done := make(chan interface{})

	for i := 0; i < 5; i++ {
		go test(i+1, out)
	}
	var x int
	for y := range out {
		fmt.Println(".....", y)
		x++
		if x == 5 {
			close(out)
		}

	}
	fmt.Println("....", time.Since(st))
	st = time.Now()
	for i := 0; i < 5; i++ {
		test2(i + 1)
	}
	fmt.Println("....", time.Since(st))

}

func test(x int, out chan<- string) {
	time.Sleep(time.Millisecond)
	out <- "hello - " + fmt.Sprintf("%v", x)
}

func test2(x int) string {
	time.Sleep(time.Millisecond)
	return "hello - " + fmt.Sprintf("%v", x)
}
