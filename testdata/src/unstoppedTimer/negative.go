package unclosedResource

import (
	"time"
)

func negative1() {
	ff := time.NewTimer(time.Second)
	defer ff.Stop()

	ff = time.NewTimer(time.Second)

	print(<-ff.C)
	ff.Stop()
}

var globalVar *time.Timer

func negative2() {
	globalVar = time.NewTimer(time.Second) // global var
	kk := <-globalVar.C

	print(kk)
}

func negative3() *time.Timer {
	timer := time.NewTimer(time.Second) // var escape the function

	return timer
}

func negative4() {
	var timers []*time.Timer
	timer := time.NewTimer(time.Second) // var escape the function in another var

	timers = append(timers, timer)
}

func negative5() {
	var filesMap map[time.Time]*time.Timer
	timer := time.NewTimer(time.Second) // var escape the function in another var

	filesMap[<-timer.C] = timer
}

func negative6() {
	type st struct {
		*time.Timer
	}
	var (
		timerDecorator1 st
		timerDecorator2 st
	)
	timer := time.NewTimer(time.Second) // var escape the function in another var

	timerDecorator1 = st{timer}
	timerDecorator2 = st{
		Timer: timer,
	}

	kk, kkk := <-timerDecorator1.C, <-timerDecorator2.C

	print(kk, kkk)
}

func negative7() {
	var ch chan *time.Timer
	timer := time.NewTimer(time.Second) // var escape the function in another var

	ch <- timer
}

func negative8() {
	closure := func() (*time.Timer, error) {
		return nil, nil
	}

	f, _ := closure()
	defer f.Stop()
}

func negative9() {
	closure := func() *time.Timer {
		return nil
	}

	f := closure()

	print(<-f.C)
}

type MyStruct struct {
	t *time.Timer
}

func (k MyStruct) negative10() {
	kk := time.NewTimer(time.Second)
	k.t = kk
}

func negative11() *time.Timer {
	var k *time.Timer
	k = time.NewTimer(time.Second)

	kk := k

	return kk
}
