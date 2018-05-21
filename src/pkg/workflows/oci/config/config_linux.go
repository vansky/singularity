// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package config

const Platform = "linux"

type RuntimeOciPlatform struct {
	Linux   RuntimeOciLinux
	Solaris interface{}
	Windows interface{}
}