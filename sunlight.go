package sunlightApi

import "net/url"

// SunlightAPI is the interface for all of the
// Sunlight Foundation's APIs
type SunlightAPI interface {
	call(string, url.Values) ([]byte, error)
}
