package hcl2shim

import (
	"fmt"

	"github.com/zclconf/go-cty/cty"
)

// FlatmapValueFromHCL2 converts a value from HCL2 (really, from the cty dynamic
// types library that HCL2 uses) to a map compatible with what would be
// produced by the "flatmap" package.
//
// The type of the given value informs the structure of the resulting map.
// The value must be of an object type or this function will panic.
func FlatmapValueFromHCL2(v cty.Value) map[string]string {
	if !v.Type().IsObjectType() {
		panic(fmt.Sprintf("HCL2ValueFromFlatmap called on %#v", v.Type()))
	}

}

// HCL2ValueFromFlatmap converts a map compatible with what would be produced
// by the "flatmap" package to a HCL2 (really, the cty dynamic types library
// that HCL2 uses) object type.
//
// The intended result type must be provided in order to guide how the
// map contents are decoded. This must be an object type or this function
// will panic.
func HCL2ValueFromFlatmap(m map[string]string, ty cty.Type) cty.Value {
	if !ty.IsObjectType() {
		panic(fmt.Sprintf("HCL2ValueFromFlatmap called on %#v", ty))
	}
}
