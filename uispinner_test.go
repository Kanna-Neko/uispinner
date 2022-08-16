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
		spinner2 := spinner1.AddSpinner(spinner.CharSets[9], 200*time.Millisecond).SetComplete(fmt.Sprintf("part %d had completed", i+1)).SetPrefix(fmt.Sprintf("part os process %d", i+1))
		for j := 0; j < 3; j++ {
			spinner3 := spinner2.AddSpinner(spinner.CharSets[0], 200*time.Millisecond).SetPrefix(fmt.Sprintf("this is process %d", j+1)).SetComplete(fmt.Sprintf("process %d done", j+1))
			for k := 0; k < 3; k++ {
				spinner4 := spinner3.AddSpinner(spinner.CharSets[0], 200*time.Millisecond).SetPrefix(fmt.Sprintf("this is atomic process %d", k+1)).SetComplete(fmt.Sprintf("atomic process %d done", k+1))
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
	sp := cj.AddSpinner(spinner.CharSets[20], 499*time.Millisecond).SetPrefix("test").SetSuffix("test").SetComplete("helloWorld").Reverse()
	fmt.Fprintf(cj.Bypass(),"Bypass Test\n")
	time.Sleep(time.Second)
	sp.SetInterval(1000 * time.Millisecond)
	sp.SetCharSet(spinner.CharSets[21])
	time.Sleep(time.Second)
	sp.Done()
	cj.Stop()
}
