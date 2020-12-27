package oauth

import (
	"strings"
)

const (
	ScopesMatched    = true
	ScopesNotMatched = false
)

func CompareScope(source []string, scope string) bool {
	isExist := make(map[string]bool)

	scopes := strings.Split(scope, " ")
	for _, scp := range scopes {
		isExist[scp] = true
	}

	for _, scp := range source {
		if isExist[scp] == false {
			return ScopesNotMatched
		}
	}

	return ScopesMatched
}
