package syserror

import ("fmt")

type sysError struct {
	msg string	// message associated to the error
	code int	// error code
}

// Predefined errors
var NoErr = sysError{"Success", 0} // Success
var ErrFatal = sysError{"Fatal error", -1} // Generic fatal error
var ErrOutOfRes = sysError{"Out of resources", -2} // Out of resources (e.g., out of memory)
var ErrNotAvailable = sysError{"Not available", -3} // The target is not available (e.g., a file does not exist)


func (err *sysError) Error() string {
	if err.code != 0 {
		return fmt.Sprintf ("ERROR: %d: %s", err.code, err.msg)
	} else {
		return fmt.Sprintf ("%d: %s", err.code, err.msg)
	}
}

func (_error sysError) getMsg() string {
	return _error.msg
}

func (_error sysError) getCode() int {
	return _error.code
}

