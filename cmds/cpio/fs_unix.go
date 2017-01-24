// Copyright 2013-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
	"time"
)

// modes sets the modes, changing the easy ones first and the harder ones last.
// In this way, we set as much as we can before bailing out. It's not an error
// to not be able to set uid and gid, at least not yet.
// For now we also ignore sticky bits.
func modes(f *File) error {
	if err := os.Chmod(f.Name, os.FileMode(f.Mode)); err != nil {
		return err
	}
	if err := os.Chtimes(f.Name, time.Now(), time.Unix(int64(f.Mtime), 0)); err != nil {
		return err
	}
	if err := os.Chown(f.Name, int(f.UID), int(f.GID)); err != nil {
		return err
	}
	// TODO: only set SUID and GUID if we can set the owner.
	return nil
}

func create(f *File) error {

	m, err := cpioModetoMode(f.Mode)
	if err != nil {
		return err
	}

	switch m {
	case os.ModeSocket, os.ModeNamedPipe:
		return fmt.Errorf("%v: type %v: we do not create IPC endpoitns", f.Name, m)
	case os.FileMode(0):
		nf, err := os.Create(f.Name)
		if err != nil {
			return err
		}
		_, err = io.Copy(nf, f.Data)
		if err != nil {
			return err
		}
		return modes(f)
	case os.ModeDir:
		err = os.MkdirAll(f.Name, os.FileMode(f.Mode))
		if err != nil {
			return err
		}
		return modes(f)
	case os.ModeDevice:
		if err = syscall.Mknod(f.Name, perm(f)|syscall.S_IFBLK, dev(f)); err != nil {
			return err
		}
		return modes(f)
	case os.ModeCharDevice:
		if err = syscall.Mknod(f.Name, perm(f)|syscall.S_IFCHR, dev(f)); err != nil {
			return err
		}
		return modes(f)
	default:
		return fmt.Errorf("%v: Unknown type %#o", f.Name, m)
	}

}
