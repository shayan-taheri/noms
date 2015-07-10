package types

import (
	"github.com/attic-labs/noms/ref"
)

type String struct {
	s  string
	cr *cachedRef
}

func NewString(s string) String {
	return String{s, &cachedRef{}}
}

func (fs String) Blob() Blob {
	return NewBlob([]byte(fs.s))
}

func (fs String) String() string {
	return fs.s
}

func (fs String) Ref() ref.Ref {
	return fs.cr.Ref(fs)
}

func (fs String) Equals(other Value) bool {
	if other == nil {
		return false
	} else {
		return fs.Ref() == other.Ref()
	}
}
