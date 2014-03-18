package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	//"time"

	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 3 {
		log.Fatal("Usage:\n  var_sym ROOT_DIR ENV_VAR_NAME MOUNTPOINT")
	}

	v := &VarSymFs{
		root_dir:    flag.Arg(0),
		env_var:     flag.Arg(1),
		mount_point: flag.Arg(2),
		FileSystem:  NewDefaultFileSystem(),
	}
	nfs := pathfs.NewPathNodeFs(v, nil)

	// opts := nodefs.Options{
	// 	EntryTimeout:    time.Millisecond,
	// 	AttrTimeout:     time.Millisecond,
	// 	NegativeTimeout: time.Millisecond,
	// }

	server, _, err := nodefs.MountRoot(v.mount_point, nfs.Root(), nil)
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}

	// unmount when we die...
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)
	go func() {
		for sig := range ch {
			buf := make([]byte, 32)
			for {
				n := runtime.Stack(buf, true)
				if n < len(buf) {
					break
				}
				buf = make([]byte, len(buf)*2)
			}
			fmt.Println(string(buf))
			err := server.Unmount()
			if sig == os.Interrupt {
				var i int
				if err != nil {
					i = 1
				}
				os.Exit(i)
			}
		}
	}()

	// serve
	server.Serve()
}
