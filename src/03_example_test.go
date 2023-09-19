package examples

import (
	"sync/atomic"
	"testing"
)

/* execute the test by running:
go test ./... --race
*/
func TestDataRaceConditions(t *testing.T) {
	var state int32

	for i := 0;i < 10;i++ {
		go func(i int) {
			// state += int32(i)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
