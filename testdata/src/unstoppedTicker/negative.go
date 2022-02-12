package unclosedResource

import (
	"time"
)

func negative1() {
	ff := time.NewTicker(time.Second)
	defer ff.Stop()

	ff = time.NewTicker(time.Second)

	print(<-ff.C)
	ff.Stop()
}

var globalVar *time.Ticker

func negative2() {
	globalVar = time.NewTicker(time.Second) // global var
	kk := <-globalVar.C

	print(kk)
}

func negative3() *time.Ticker {
	timer := time.NewTicker(time.Second) // var escape the function

	return timer
}

func negative4() {
	var timers []*time.Ticker
	timer := time.NewTicker(time.Second) // var escape the function in another var

	timers = append(timers, timer)
}

func negative5() {
	var filesMap map[time.Time]*time.Ticker
	timer := time.NewTicker(time.Second) // var escape the function in another var

	filesMap[<-timer.C] = timer
}

func negative6() {
	type st struct {
		*time.Ticker
	}
	var (
		timerDecorator1 st
		timerDecorator2 st
	)
	timer := time.NewTicker(time.Second) // var escape the function in another var

	timerDecorator1 = st{timer}
	timerDecorator2 = st{
		Ticker: timer,
	}

	kk, kkk := <-timerDecorator1.C, <-timerDecorator2.C

	print(kk, kkk)
}

func negative7() {
	var ch chan *time.Ticker
	timer := time.NewTicker(time.Second) // var escape the function in another var

	ch <- timer
}

func negative8() {
	closure := func() (*time.Ticker, error) {
		return nil, nil
	}

	f, _ := closure()
	defer f.Stop()
}

func negative9() {
	closure := func() *time.Ticker {
		return nil
	}

	f := closure()

	print(<-f.C)
}

type MyStruct struct {
	t *time.Ticker
}

func (k MyStruct) negative10() {
	kk := time.NewTicker(time.Second)
	k.t = kk
}

func negative11() *time.Ticker {
	var k *time.Ticker
	k = time.NewTicker(time.Second)

	kk := k

	return kk
}
