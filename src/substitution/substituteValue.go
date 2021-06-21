package substitution

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

func unescape(input []byte) ([]byte, error) {
	result, err := url.QueryUnescape(string(input))
	if err != nil {
		return input, err
	}
	return []byte(result), nil
}

// Swaps a <value:...> for the value from the valuesource
// input should contain no lf/cf
func (s *Substitutor) substituteValueWithError(input []byte) ([]byte, error) {
	reOuter := regexp.MustCompile(`^<\s*crumblecog\s*:\s*(.*[^\s])\s*>$`)
	tagFound := reOuter.FindSubmatch(input)
	if tagFound != nil {
		if len(tagFound[1]) > 0 {
			tag, err := unescape(tagFound[1])
			if err != nil {
				return nil, err
			}
			return s.lookupTag(tag)
		}
		return nil, errors.New(`Failed to find tag for substitution`)
	}
	// We pass through things we can't match at all. They shouldn't arrive here.
	return input, nil
}

// Swaps a <value:...> for the value from the valuesource
// input should contain no lf/cf
func (s *Substitutor) substituteValue(input []byte) []byte {
	res, err := s.substituteValueWithError(input)
	if err != nil {
		longerr := fmt.Errorf("Processing %s failed: %s", string(input), err)
		if s.errs == nil {
			s.errs = longerr
		} else {
			s.errs = fmt.Errorf("%s\n%s", s.errs, longerr)
		}
	}
	return res
}

// This is the actual work
func (s *Substitutor) lookupTag(input []byte) ([]byte, error) {
	switch string(input) {
	case `domain`:
		return []byte(fmt.Sprintf("<secret:%s~domain>", configSecretPath())), nil
	case `cf-api-token`:
		return []byte(fmt.Sprintf("<secret:%s/cloudflare~password>", configSecretPath())), nil
	default:
		return nil, fmt.Errorf("Unknown key %q", input)
	}

}
