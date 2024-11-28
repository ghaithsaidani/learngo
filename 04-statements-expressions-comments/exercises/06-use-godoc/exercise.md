## EXERCISE

- Print the documentation of `runtime.NumCPU` function in the command line
- Print also its source code using in the command line

## HINT

You should use correct `go doc` tools


### documentation
package runtime // import "runtime"

func NumCPU() int
    NumCPU returns the number of logical CPUs usable by the current process.

    The set of available CPUs is checked by querying the operating system at
    process startup. Changes to operating system CPU allocation after process
    startup are not reflected.
 
### source code
package runtime // import "runtime"

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
func NumCPU() int {
        return int(ncpu)
}