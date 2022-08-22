package pattern

import (
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Template struct {
	Segments []*Segment `parser:"@@ ( '/' @@ )*"`
}

type Segment struct {
	Value    string    `parser:"@(Ident | '-') | "`
	Variable *Variable `parser:"@@"`
}

type Variable struct {
	Value string `parser:"'{' @(Ident ( '.' Ident )*) '}'"`
}

var parser = participle.MustBuild[Template]()

// Template = Segment { "/" Segment } ;
// Segment = LITERAL | Variable ;
// Variable = "{" LITERAL "}" ;

// Match returns false if all is set to true and the pattern doesn't fully match
// the specified value.
func MatchString(pattern, value string, all bool) bool {
	_, ok := FindString(pattern, value, all)
	return ok
}

// Find returns false if all is set to true and the pattern doesn't fully match
// the specified value.
func FindString(pattern, value string, all bool) (map[string]string, bool) {
	m := make(map[string]string)

	ast, err := parser.ParseString("", pattern)
	if err != nil {
		return nil, false
	}

	segments := strings.Split(strings.Trim(value, "/"), "/")
	if len(ast.Segments) != len(segments) && (all || len(ast.Segments) > len(segments)) {
		return nil, false
	}

	for i, segment := range segments {
		if len(ast.Segments) < i+1 {
			break
		}

		if ast.Segments[i].Variable != nil {
			if segment == "" {
				return nil, false
			}

			m[ast.Segments[i].Variable.Value] = segment
		} else if segment != ast.Segments[i].Value {
			return nil, false
		}
	}

	return m, true
}

func FindStringSubmatch(pattern, value string) (string, bool) {
	ast, err := parser.ParseString("", pattern)
	if err != nil {
		return "", false
	}

	segments := strings.Split(strings.Trim(value, "/"), "/")
	if len(ast.Segments) > len(segments) {
		return "", false
	}

	var s []string
	for i, segment := range segments {
		if len(ast.Segments) < i+1 {
			break
		}

		if ast.Segments[i].Variable != nil {
			if segment == "" {
				return "", false
			}
		} else if segment != ast.Segments[i].Value {
			return "", false
		}

		s = append(s, segment)
	}

	return strings.Join(s, "/"), true
}

func ReplaceString(pattern string, variables map[string]string) (string, bool) {
	ast, err := parser.ParseString("", pattern)
	if err != nil {
		return "", false
	}

	var s []string
	for _, i := range ast.Segments {
		if i.Variable != nil {
			if variables == nil {
				return "", false
			}

			variable, ok := variables[i.Variable.Value]
			if !ok {
				return "", false
			}

			s = append(s, variable)
		} else {
			s = append(s, i.Value)
		}
	}

	return strings.Join(s, "/"), true
}
