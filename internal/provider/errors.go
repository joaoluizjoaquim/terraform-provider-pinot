package provider

import "strings"

// isNotFoundError reports whether err represents a 404 response from the Pinot
// controller. go-pinot-api returns plain string errors (no typed/sentinel
// errors), so we match on the status code embedded in the message, e.g.
// `client: request failed with status code: 404, body: {...}`.
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "status code: 404")
}
