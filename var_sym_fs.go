package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type VarSymFs struct {
	root_dir    string
	env_var     string
	mount_point string
	pathfs.FileSystem
}

var debug bool = true

func do_debug(s string, args ...interface{}) {
	if debug {
		fmt.Println(fmt.Sprintf(s, args))
	}
}

func (me *VarSymFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	do_debug("GetAttr %v", name)
	root, err := me.contextToDir(context)
	if err != nil {
		fmt.Println(err)
		return nil, fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	a, b := loopback.GetAttr(name, context)
	return a, b
}

func (me *VarSymFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
	do_debug("OpenDir %v", name)
	root, err := me.contextToDir(context)
	if err != nil {
		return nil, fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.OpenDir(name, context)
}

func (me *VarSymFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	do_debug("Open %v", name)
	root, err := me.contextToDir(context)
	if err != nil {
		return nil, fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Open(name, flags, context)
}

// func (fs *VarSymFs) SetDebug(debug bool) {
// 	root, err := fs.contextToDir(context)
// 	if err != nil {
// 		return nil, fuse.ENOENT
// 	}
// 	loopback := pathfs.NewLoopbackFileSystem(root)
// 	return loopback.SetDebug(debug)
// }

func (fs *VarSymFs) GetXAttr(name string, attr string, context *fuse.Context) ([]byte, fuse.Status) {
	do_debug("GetXAttr %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return nil, fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.GetXAttr(name, attr, context)
}

func (fs *VarSymFs) SetXAttr(name string, attr string, data []byte, flags int, context *fuse.Context) fuse.Status {
	do_debug("SetXAttr %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.SetXAttr(name, attr, data, flags, context)
}

func (fs *VarSymFs) ListXAttr(name string, context *fuse.Context) ([]string, fuse.Status) {
	do_debug("ListXAttr %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return nil, fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.ListXAttr(name, context)
}

func (fs *VarSymFs) RemoveXAttr(name string, attr string, context *fuse.Context) fuse.Status {
	do_debug("RemoveXAttr %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.RemoveXAttr(name, attr, context)
}

func (fs *VarSymFs) Readlink(name string, context *fuse.Context) (string, fuse.Status) {
	do_debug("Readlink %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return "", fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Readlink(name, context)
}

func (fs *VarSymFs) Mknod(name string, mode uint32, dev uint32, context *fuse.Context) fuse.Status {
	do_debug("Mknod %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Mknod(name, mode, dev, context)
}

func (fs *VarSymFs) Mkdir(name string, mode uint32, context *fuse.Context) fuse.Status {
	do_debug("Mkdir %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Mkdir(name, mode, context)
}

func (fs *VarSymFs) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	do_debug("Unlink %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Unlink(name, context)
}

func (fs *VarSymFs) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	do_debug("Rmdir %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Rmdir(name, context)
}

func (fs *VarSymFs) Symlink(value string, linkName string, context *fuse.Context) (code fuse.Status) {
	do_debug("Symlink %v, %v", value, linkName)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Symlink(value, linkName, context)
}

func (fs *VarSymFs) Rename(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	do_debug("Rename %v, %v", oldName, newName)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Rename(oldName, newName, context)
}

func (fs *VarSymFs) Link(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	do_debug("Link %v, %v", oldName, newName)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Link(oldName, newName, context)
}

func (fs *VarSymFs) Chmod(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	do_debug("Chmod %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Chmod(name, mode, context)
}

func (fs *VarSymFs) Chown(name string, uid uint32, gid uint32, context *fuse.Context) (code fuse.Status) {
	do_debug("Chown %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Chown(name, uid, gid, context)
}

func (fs *VarSymFs) Truncate(name string, offset uint64, context *fuse.Context) (code fuse.Status) {
	do_debug("Truncate %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Truncate(name, offset, context)
}

// func (fs *VarSymFs) OnMount(nodeFs *pathfs.PathNodeFs) {
// 	fmt.Println("OnMount")
// }

// func (fs *VarSymFs) OnUnmount() {
// 	fmt.Println("OnUnmount")
// }

func (fs *VarSymFs) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	do_debug("Access %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Access(name, mode, context)
}

func (fs *VarSymFs) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	do_debug("Create %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return nil, fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Create(name, flags, mode, context)
}

func (fs *VarSymFs) Utimens(name string, Atime *time.Time, Mtime *time.Time, context *fuse.Context) (code fuse.Status) {
	do_debug("Utimens %v", name)
	root, err := fs.contextToDir(context)
	if err != nil {
		return fuse.ENOENT
	}
	loopback := pathfs.NewLoopbackFileSystem(root)
	return loopback.Utimens(name, Atime, Mtime, context)
}

func (fs *VarSymFs) String() string {
	return "VarSymFs"
}

// func (fs *VarSymFs) StatFs(name string) *fuse.StatfsOut {
// 	fmt.Println("StatFs")
// 	return nil
// }

func (f *VarSymFs) contextToDir(context *fuse.Context) (string, error) {
	if context == nil {
		return "", errors.New("Nil context passed")
	}
	i := 1
	pid := fmt.Sprintf("%v", context.Pid)

	// walk until we hit the parent if needs be
	for i <= 1 {
		fmt.Println(fmt.Sprintf("/proc/%v/environ", pid))
		env, err := os.Open(fmt.Sprintf("/proc/%v/environ", pid))
		if err != nil {
			return "", errors.New(fmt.Sprintf("Could not get env for %v", pid))
		}

		re := bufio.NewReader(env)
		for {
			line, err := re.ReadString(0)
			if err != nil {
				break
			}
			// strip the trailing null
			line = line[:len(line)-1]

			sp := strings.SplitN(line, "=", 2)
			key, val := sp[0], sp[1]
			if key == f.env_var {
				return filepath.Join(f.root_dir, val), nil
			}
		}

		// at this point move to the parent
		fi, _ := os.Open(fmt.Sprintf("/proc/%v/stat", pid))
		sc := bufio.NewScanner(fi)
		sc.Split(bufio.ScanWords)
		for j := 1; sc.Scan() && j <= 4; j++ {
			pid = sc.Text()
		}

		i++
	}

	return "", errors.New(fmt.Sprintf("Could not find %v in env for %v or parent %v", f.env_var, context.Pid, pid))
}
