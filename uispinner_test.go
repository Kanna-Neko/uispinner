package uispinner

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/briandowns/spinner"
)

func TestSpinner(t *testing.T) {
	pool := sync.WaitGroup{}
	cj := New()
	spinner1 := cj.AddSpinner(spinner.CharSets[8], 300*time.Millisecond).SetComplete("All task had completed")
	for i := 0; i < 3; i++ {
		spinner2 := spinner1.AddSpinner(spinner.CharSets[9], 200*time.Millisecond).SetComplete(fmt.Sprintf("part %d had completed", i))
		for j := 0; j < 3; j++ {
			spinner3 := spinner2.AddSpinner(spinner.CharSets[0], 200*time.Millisecond).SetPrefix(fmt.Sprintf("this is process %d", j)).SetComplete(fmt.Sprintf("process %d done", j))
			for k := 0; k < 3; k++ {
				spinner4 := spinner3.AddSpinner(spinner.CharSets[0], 200*time.Millisecond).SetPrefix(fmt.Sprintf("this is atomic process %d", k)).SetComplete(fmt.Sprintf("atomic process %d done", k))
				pool.Add(1)
				go func() {
					time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
					spinner4.Done()
					pool.Done()
				}()
			}
		}
	}
	cj.Start()
	pool.Wait()
	cj.Stop()
}
