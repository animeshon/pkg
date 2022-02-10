package validation

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/gobwas/glob"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
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

	// Make sure the ids are int 64
	resourceParts := strings.Split(name, "/")
	categoriesCount := len(resourceParts) / 2
	for catNumber := 0; catNumber < categoriesCount; catNumber++ {
		_, err := strconv.ParseInt(resourceParts[catNumber*2+1], 10, 64)
		if err != nil {
			allErrs = append(allErrs, Invalid(fldPath, name, "resource ids must be valid int64"))
		}
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

func ValidateFieldMask(mask *fieldmaskpb.FieldMask, m protoreflect.ProtoMessage, fldPath *field.Path) ErrorList {
	allErrs := ErrorList{}

	if mask == nil {
		return allErrs
	}

	mask.Normalize()
	if !mask.IsValid(m) {
		return append(allErrs, Invalid(fldPath, mask, "the specified field mask is invalid"))
	}

	return allErrs
}

func ValidatePageSize(size int32, fldPath *field.Path) ErrorList {
	allErrs := ErrorList{}

	if size < 0 {
		allErrs = append(allErrs, Invalid(fldPath, size, "page size must be non-negative"))
	}

	return allErrs
}
