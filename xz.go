package xz

import (
	"errors"
	"io"

	uxz "github.com/ulikunitz/xz"
)

func Xzcat(compressed io.Reader, wtr io.Writer) error {
	xrdr, e := uxz.NewReader(compressed)
	if nil != e {
		return e
	}
	_, e = io.Copy(wtr, xrdr)
	return e
}

func Xz(plain io.Reader, wtr io.Writer) error {
	xwtr, e := uxz.NewWriter(wtr)
	if nil != e {
		return e
	}
	_, e = io.Copy(xwtr, plain)
	return errors.Join(e, xwtr.Close())
}
