package rule

import "regexp"

type Mode interface {
	RegEx() regexp.Regexp
	Symbol() string
}
