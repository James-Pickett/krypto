//go:build linux
// +build linux

package krypto

import (
	"io"

	"github.com/google/go-tpm/tpm2"
)

func newTpmEncoder() *tpmEncoder {
	return &tpmEncoder{
		openTpm: func() (io.ReadWriteCloser, error) {
			return tpm2.OpenTPM("/dev/tpm0")
		},
	}
}