/*
 * Copyright(c)         Geoffroy Vallee
 *                      All rights reserved
 */

package syserror

import ("fmt")

type SysError struct {
	msg string	// message associated to the error
	code int	// error code
}

// Predefined errors
var NoErr = SysError{"Success", 0} // Success
var ErrFatal = SysError{"Fatal error", -1} // Generic fatal error
var ErrOutOfRes = SysError{"Out of resources", -2} // Out of resources (e.g., out of memory)
var ErrNotAvailable = SysError{"Not available", -3} // The target is not available (e.g., a file does not exist)
var ErrDataOverflow = SysError{"Data overflow", -4} // Data overflow

func (err *SysError) Error() string {
	if err.code != 0 {
		return fmt.Sprintf ("ERROR: %d: %s", err.code, err.msg)
	} else {
		return fmt.Sprintf ("%d: %s", err.code, err.msg)
	}
}

func (_error SysError) getMsg() string {
	return _error.msg
}

func (_error SysError) getCode() int {
	return _error.code
}

