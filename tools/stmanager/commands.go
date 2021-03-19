// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/u-root/u-root/pkg/boot/stboot"
)

func packOSPackage(out, label, pkgURL, kernel, initramfs, cmdline, tboot, tbootArgs string, acms []string) error {
	ospkg, err := stboot.InitOSPackage(out, label, pkgURL, kernel, initramfs, cmdline, tboot, tbootArgs, acms)
	if err != nil {
		return err
	}

	err = ospkg.Pack()
	if err != nil {
		return err
	}

	newArchive := filepath.Base(ospkg.Archive)
	os.Stdout.WriteString(newArchive)
	return nil
}

func addSignatureToOSPackage(osPackage, privKey, cert string) error {
	ospkg, err := stboot.OSPackageFromArchive(osPackage)
	if err != nil {
		return err
	}

	log.Print("Signing OS package ...")
	log.Printf("private key: %s", privKey)
	log.Printf("certificate: %s", cert)
	err = ospkg.Sign(privKey, cert)
	if err != nil {
		return err
	}

	if err = ospkg.Pack(); err != nil {
		return err
	}

	log.Printf("Signatures included: %d", len(ospkg.Descriptor.Signatures))
	return nil
}

func unpackOSPackage(ospkgPath string) error {
	ospkg, err := stboot.OSPackageFromArchive(ospkgPath)
	if err != nil {
		return err
	}

	log.Println("Archive unpacked into: " + ospkg.Archive)
	return nil
}
