package cfg // import "gopkg.in/webnice/cfg.v1/cfg"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
//import ()

// New Function creates object and return as interface
func New() Interface {
	var trt = new(impl)
	return trt
}
