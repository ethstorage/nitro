// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package storage

import (
	"os"
	"syscall"
)

func newFileLock(path string, readOnly bool) (fl fileLock, err error) {
	return nil, nil
}

func setFileLock(f *os.File, readOnly, lock bool) error {
	return nil
}

func rename(oldpath, newpath string) error {
	return syscall.ENOTSUP
}

func isErrInvalid(err error) bool {
	return false
}

func syncDir(name string) error {
	return nil
}
