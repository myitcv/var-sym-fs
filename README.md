
# ** IMPORTANT **

This code is unmainted and more significantly, does not work. It is here for historical purposes only.

## var-sym-fs

A very basic and rough implementation of [variant symlinks](https://wiki.freebsd.org/200808DevSummit?action=AttachFile&do=get&target=variant-symlinks-for-freebsd.pdf) using [`go-fuse`](https://github.com/hanwen/go-fuse)

With time this will accept mount requests of the form:

```
var-sym-fs <ROOT_DIR> <ENV_VAR_NAME> <MOUNT_POINT>
```

So for example:

```
var-sym-fs /home/myitcv/.gos GO_VERSION /home/myitcv/go
```

with the command:

```bash
GO_VERSION=go1.2.1 ls /home/myitcv/go
```

would list the contents of `/home/myitcv/.gos/go1.2.1`. 

Furthermore (and this is potentially the most powerful use case), `PATH`, `GOPATH` etc could be defined to include paths that are variant symlinks. This would do away with the need for [`gvm`](https://github.com/moovweb/gvm), [`rbenv`](https://github.com/sstephenson/rbenv) and other such version managers. Indeed it would also do away with the need for many package managers (slight caveat here because there are some important use cases `var-sym-fs` would not handle that `gvm pkgset` does, for example local `pkgset`'s)


# Status

Work in progress. Currently blocked on two *critical* items:

* Cannot `execve` a file on the mount. Blocked on [an answer from the LKML](https://lkml.org/lkml/2014/3/17/492)
* A slight change to [`github.com/hanwen/go-fuse`](https://github.com/hanwen/go-fuse/pull/21)

# License

See the top-level `LICENSE` file
