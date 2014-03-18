package main

import (
	"fmt"
	"time"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

// NewDefaultFileSystem creates a filesystem that responds ENOSYS for
// all methods
func NewDefaultFileSystem() pathfs.FileSystem {
	return (*defaultFileSystem)(nil)
}

// defaultFileSystem implements a FileSystem that returns ENOSYS for every operation.
type defaultFileSystem struct{}

func (fs *defaultFileSystem) SetDebug(debug bool) {
	fmt.Println("SetDebug")
}

// defaultFileSystem
func (fs *defaultFileSystem) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	fmt.Println("GetAttr")
	return nil, fuse.ENOSYS
}

func (fs *defaultFileSystem) GetXAttr(name string, attr string, context *fuse.Context) ([]byte, fuse.Status) {
	fmt.Println("GetXAttr")
	return nil, fuse.ENOSYS
}

func (fs *defaultFileSystem) SetXAttr(name string, attr string, data []byte, flags int, context *fuse.Context) fuse.Status {
	fmt.Println("SetXAttr")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) ListXAttr(name string, context *fuse.Context) ([]string, fuse.Status) {
	fmt.Println("ListXAttr")
	return nil, fuse.ENOSYS
}

func (fs *defaultFileSystem) RemoveXAttr(name string, attr string, context *fuse.Context) fuse.Status {
	fmt.Println("RemoveXAttr")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Readlink(name string, context *fuse.Context) (string, fuse.Status) {
	fmt.Println("Readlink")
	return "", fuse.ENOSYS
}

func (fs *defaultFileSystem) Mknod(name string, mode uint32, dev uint32, context *fuse.Context) fuse.Status {
	fmt.Println("Mknod")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Mkdir(name string, mode uint32, context *fuse.Context) fuse.Status {
	fmt.Println("Mkdir")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Unlink(name string, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Unlink")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Rmdir(name string, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Rmdir")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Symlink(value string, linkName string, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Symlink")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Rename(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Rename")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Link(oldName string, newName string, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Link")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Chmod(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Chmod")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Chown(name string, uid uint32, gid uint32, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Chown")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Truncate(name string, offset uint64, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Truncate")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	fmt.Println("Open")
	return nil, fuse.ENOSYS
}

func (fs *defaultFileSystem) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, status fuse.Status) {
	fmt.Println("OpenDir")
	return nil, fuse.ENOSYS
}

func (fs *defaultFileSystem) OnMount(nodeFs *pathfs.PathNodeFs) {
	fmt.Println("OnMount")
}

func (fs *defaultFileSystem) OnUnmount() {
	fmt.Println("OnUnmount")
}

func (fs *defaultFileSystem) Access(name string, mode uint32, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Access")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) Create(name string, flags uint32, mode uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	fmt.Println("Create")
	return nil, fuse.ENOSYS
}

func (fs *defaultFileSystem) Utimens(name string, Atime *time.Time, Mtime *time.Time, context *fuse.Context) (code fuse.Status) {
	fmt.Println("Utimens")
	return fuse.ENOSYS
}

func (fs *defaultFileSystem) String() string {
	return "defaultFileSystem"
}

func (fs *defaultFileSystem) StatFs(name string) *fuse.StatfsOut {
	fmt.Println("StatFs")
	return nil
}
