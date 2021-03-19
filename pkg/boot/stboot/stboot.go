// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stboot

const (
	// OSPackageExt is the file extension of OS packages
	OSPackageExt string = ".zip"
	// DefaultOSPackageName is the file name of the archive, which is expected to contain
	// the stboot configuration file along with the corresponding files
	DefaultOSPackageName string = "ospkg.zip"
	// ManifestName is the name of OS packages' internal configuration file
	ManifestName string = "manifest.json"
)
