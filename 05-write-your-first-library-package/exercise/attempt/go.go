package attempt

import "runtime"


func Version() string {
	return runtime.Version()
}