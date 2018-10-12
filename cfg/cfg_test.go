package cfg // import "gopkg.in/webnice/cfg.v1/cfg"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestNew(t *testing.T) {
	var ok bool
	var obj = New()

	if obj == nil {
		t.Error("Function New() return nil")
		return
	}
	if _, ok = obj.(*impl); !ok {
		t.Error("Function New() return invalid object")
		return
	}
}
