package validation

import (
	"net/url"

	"github.com/gobwas/glob"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateResourceName(name, pattern string, fldPath *field.Path) ErrorList {
	allErrs := ErrorList{}

	if len(name) == 0 {
		return append(allErrs, Required(fldPath))
	}
	if !glob.MustCompile(pattern).Match(name) {
		allErrs = append(allErrs, Invalid(fldPath, name, "resource name does not match a valid pattern"))
	}

	return allErrs
}

func ValidateURI(uri string, fldPath *field.Path) ErrorList {
	allErrs := ErrorList{}

	if len(uri) == 0 {
		return append(allErrs, Required(fldPath))
	}

	_, err := url.Parse(uri)
	if err != nil {
		return append(allErrs, Invalid(fldPath, uri, "could not parse URI"))
	}

	return allErrs
}
