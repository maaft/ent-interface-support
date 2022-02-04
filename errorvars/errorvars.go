package errorvars

import "errors"

var ErrAccessDenied = errors.New("access denied")
var ErrPasswordToShort = errors.New("password must be at least 8 characters long")
