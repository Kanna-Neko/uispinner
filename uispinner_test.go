package uispinner

import (
	"testing"
	"time"

	"github.com/briandowns/spinner"
)

func TestSpinner(t *testing.T) {
	cj := New()
	spinner1 := cj.AddSpinner(spinner.CharSets[34], 1*time.Millisecond).SetComplete("helloWorld").SetPrefix("abc").SetSuffix("ab")
	spinner2 := cj.AddSpinner(spinner.CharSets[0], 100*time.Millisecond).SetComplete("good")
	cj.Start()
	time.Sleep(time.Second * 5)
	spinner1.Done()
	time.Sleep(time.Second * 5)
	spinner2.Done()
	spinner3 := cj.AddSpinner(spinner.CharSets[0], 100*time.Millisecond).SetComplete("goodBye")
	time.Sleep(time.Second * 5)
	spinner3.Reverse()
	time.Sleep(time.Second * 5)
	spinner3.SetCharSet(spinner.CharSets[39])
	time.Sleep(time.Second * 5)
	spinner3.Done()
	cj.Stop()
}
