package unclosedResource

import (
	"time"
)

func warning1() {
	f := time.NewTimer(time.Second) //want `unstopped timer`
	print(<-f.C)

	f = time.NewTimer(time.Second) //want `unstopped timer`
	print(<-f.C)
}

func warning2() {
	f := time.NewTimer(time.Second)
	defer f.Stop()

	f = time.NewTimer(time.Second) //want `unstopped timer`
	print(<-f.C)
}

func warning3() {
	var ff = time.NewTimer(time.Second) //want `unstopped timer`
	print(<-ff.C)

	ff = time.NewTimer(time.Second) //want `unstopped timer`
	print(<-ff.C)
}

func warning4() {
	ff := time.NewTimer(time.Second) //want `unstopped timer`
	if true {
		print(ff.Reset(time.Second))
	}

	ff = time.NewTimer(time.Minute + time.Minute) //want `unstopped timer`
	print(ff.Reset(time.Second))
}

func warning5() map[time.Time]string {
	f := time.NewTimer(time.Second) //want `unstopped timer`

	return map[time.Time]string{
		<-f.C: "",
	}
}

func warning6() chan time.Time {
	f := time.NewTimer(time.Second) //want `unstopped timer`

	var k chan time.Time

	k <- <-f.C

	return k
}

func warning7() {
	f := time.NewTimer(time.Second) //want `unstopped timer`

	var k chan time.Time

	k <- <-f.C
}

func warning8() time.Time {
	f := time.NewTimer(time.Second) //want `unstopped timer`
	return <-f.C
}

func warning9() interface{} {
	type timeDecor struct {
		time.Time
	}
	f := time.NewTimer(time.Second) //want `unstopped timer`

	return timeDecor{<-f.C}
}

func warning10() map[string]time.Time {
	f := time.NewTimer(time.Second) //want `unstopped timer`

	return map[string]time.Time{
		"": <-f.C,
	}
}
