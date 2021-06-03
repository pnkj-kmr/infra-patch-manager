package task_test

// task, err := task.NewPatchTask()
// if err != nil {
// 	log.Fatal("task object errors:", err)
// }

// // PING-PONG check with Remotes
// msg := "ping"
// remoteStat := task.PingToAll(msg)
// for k, v := range remoteStat {
// 	log.Println(">>>>> HOST host:", k, "req:", msg, "res:", v)
// }

// // PATCH UPLOAD TO REMOTES
// path := filepath.Join("tmp3/test.tar.gz")
// result := task.PatchFileUploadToAll(path)
// for _, r := range result {
// 	log.Println("FILE UPLOAD : ", r.Remote, r.File, r.Size, r.Err)
// 	for _, f := range r.Data {
// 		// log.Println(f)
// 		log.Println("File-- : ", f.GetPath(), f.GetFile(), f.GetSize(), f.GetIsdir(), f.GetTime())
// 	}
// }

// // PATCH APPLY TO REMOTES
// result2 := task.ApplyPatchToAll()
// for _, r := range result2 {
// 	// log.Print("\n\n\n")
// 	for _, r2 := range r {
// 		log.Println("PATCH APPLY : ", r2.Remote, r2.RemoteApp, r2.Verified, r2.Err)
// 		for _, f := range r2.Data {
// 			log.Println("File-- : ", f.GetPath(), f.GetFile(), f.GetSize(), f.GetIsdir(), f.GetTime())
// 		}
// 	}
// }
