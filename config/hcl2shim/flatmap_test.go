package hcl2shim

import (
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestFlatmapValueFromHCL2(t *testing.T) {
	tests := []struct {
		Value cty.Value
		Want  map[string]string
	}{}
}

func TestHCL2ValueFromFlatmap(t *testing.T) {
	tests := []struct {
		Value map[string]string
		Want  cty.Value
	}{}
}
