package timewheel

import (
	"fmt"
	gotimewheel "github.com/rfyiamcool/go-timewheel"
	"gotest.tools/v3/assert"
	"time"
)
import "testing"

func TestTimeWheel(t *testing.T) {
	tick := 100 * time.Millisecond

	tw, err := gotimewheel.NewTimeWheel(100*time.Millisecond, 100)
	assert.NilError(t, err)

	tw.Start()
	defer tw.Stop()

	delayTime := time.Second

	notify := make(chan time.Time, 2)

	now := time.Now()
	tw.Add(delayTime, func() {
		now := time.Now()
		fmt.Printf("dely call, %v\n", now)
		notify <- now
	})
	fmt.Printf("add call, %v\n", now)

	occuredTime := <-notify

	assert.Equal(t, occuredTime.Sub(now) >= delayTime, true)
	assert.Equal(t, occuredTime.Sub(now) < delayTime+tick, true)
}

func TestTimeWheelCron(t *testing.T) {
	tw, err := gotimewheel.NewTimeWheel(100*time.Millisecond, 100)
	assert.NilError(t, err)

	tw.Start()
	defer tw.Stop()

	delayTime := time.Second

	tw.AddCron(delayTime, func() {
		now := time.Now()
		fmt.Printf("dely call, %v\n", now)
	})

	time.Sleep(10 * time.Minute)
}
