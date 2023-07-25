//go:build boringcrypto || goexperiment.opensslcrypto || goexperiment.cngcrypto

package main

import "crypto/boring"

func init() {
	// This function only exists when a backend is enabled, and always returns true.
	if boring.Enabled() {
		backend = "enabled"
	} else {
		backend = "disabled"
	}
}
