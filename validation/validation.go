package validation

import (
	"unicode"

	"github.com/gobwas/glob"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

const (
	UsernameMinLength = 3
	UsernameMaxLength = 32

	GivenNameMaxLength  = 32
	FamilyNameMaxLength = 32
)

func isValidName(name string) bool {
	for _, r := range name {
		if unicode.IsSymbol(r) ||
			unicode.IsControl(r) ||
			!unicode.IsPrint(r) {
			return false
		}
	}

	return true
}

func ValidateResoureName(name, pattern string, fldPath *field.Path) ErrorList {
	allErrs := ErrorList{}

	if len(name) == 0 {
		allErrs = append(allErrs, Required(fldPath))
	}
	if !glob.MustCompile(pattern).Match(name) {
		allErrs = append(allErrs, Invalid(fldPath, name, "resource name does not match a valid pattern"))
	}

	return allErrs
}
