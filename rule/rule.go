package rule

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// Rule is a single rule that will get checked on every HTTP request.
type Rule struct {
	// ID is the unique id of the rule. It can be at most 190 characters long, but the layout of the ID is up to you.
	// You will need this ID later on to update or delete the rule.
	ID string

	// MatchesMethods as an array of HTTP methods (e.g. GET, POST, PUT, DELETE, ...). When ORY Oathkeeper searches for rules
	// to decide what to do with an incoming request to the proxy server, it compares the HTTP method of the incoming
	// request with the HTTP methods of each rules. If a match is found, the rule is considered a partial match.
	MatchesMethods []string

	// MatchesURLCompiled is a regular expression of paths this rule matches.
	MatchesURLCompiled *regexp.Regexp

	// MatchesURL is a regular expression of paths this rule matches.
	MatchesURL string

	// RequiredScopes is a list of scopes that are required by this rule.
	RequiredScopes []string

	// RequiredScopes is the action this rule requires.
	RequiredAction string

	// RequiredScopes is the resource this rule requires.
	RequiredResource string

	// AllowAnonymousModeEnabled if set true allows anonymous access to the specified paths and methods.
	AllowAnonymousModeEnabled bool

	// PassThroughModeEnabled if set true disables firewall capabilities.
	PassThroughModeEnabled bool

	// BasicAuthorizationModeEnabled if set true disables checking access control policies.
	BasicAuthorizationModeEnabled bool

	// Description describes the rule.
	Description string
}

func (r *Rule) IsMatching(method string, u *url.URL) error {
	if !stringInSlice(method, r.MatchesMethods) {
		return errors.Errorf("Method %s does not match any of %v", method, r.MatchesMethods)
	}

	if !r.MatchesURLCompiled.MatchString(u.String()) {
		return errors.Errorf("Path %s does not match %s", u.String(), r.MatchesURL)
	}

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.EqualFold(a, b) {
			return true
		}
	}
	return false
}
