# var-sym-fs

# Status

Work in progress. Currently blocked on two *critical* items:

* Cannot `execve` any file on the mount. Blocked on [an answer from the LKML](https://lkml.org/lkml/2014/3/17/492)
* A slight change to [`github.com/hanwen/go-fuse`](https://github.com/hanwen/go-fuse/pull/21)
