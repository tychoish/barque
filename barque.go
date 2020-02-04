/*
Package barque holds a a number of application level constants and
shared resources for the cedar application.
*/
package barque

import (
	"time"
)

const (
	ShortDateFormat = "2006-01-02T15:04"
)

// BuildRevision stores the commit in the git repository at build time
// and is specified with -ldflags at build time
var BuildRevision = ""

const (
	AuthTokenCookie  = "barque-token"
	APIUserHeader    = "Api-User"
	APIKeyHeader     = "Api-Key"
	TokenExpireAfter = time.Hour
)
