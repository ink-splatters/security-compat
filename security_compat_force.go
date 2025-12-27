//go:build darwin && cgo && sectrust_compat

package security_compat

/*
#cgo CFLAGS: -DSECTRUST_COMPAT
*/
import "C"
