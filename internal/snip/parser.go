package snip

import (
	"fmt"
)

// ParseSnipArguments takes a Snip and a slice of raw arguments,
// and returns a map of argument names to their values, or an error.
func ParseSnipArguments(s *Snip, rawArgs []string) (map[string]string, error) {
	if len(rawArgs) != len(s.Args) {
		return nil, fmt.Errorf("expected %d arguments (%v), got %d", len(s.Args), s.Args, len(rawArgs))
	}

	parsedArgs := make(map[string]string)
	for i, argName := range s.Args {
		parsedArgs[argName] = rawArgs[i]
	}

	return parsedArgs, nil
}
