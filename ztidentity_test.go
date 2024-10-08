package ztidentity

import (
	"testing"
)

func TestNewZeroTierIdentity(t *testing.T) {

	ident := NewZeroTierIdentity()
	if len(ident.IDString()) != 10 {
		t.Errorf("Identity should be 10 digits. Got %s", ident.IDString())
    }
}
