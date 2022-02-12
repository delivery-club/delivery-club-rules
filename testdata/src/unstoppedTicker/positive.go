package unclosedResource

import (
	"time"
)

func warning1() {
	f := time.NewTicker(time.Second) //want `unstopped ticker`
	print(<-f.C)

	f = time.NewTicker(time.Second) //want `unstopped ticker`
	print(<-f.C)
}

func warning2() {
	f := time.NewTicker(time.Second)
	defer f.Stop()

	f = time.NewTicker(time.Second) //want `unstopped ticker`
	print(<-f.C)
}

func warning3() {
	var ff = time.NewTicker(time.Second) //want `unstopped ticker`
	print(<-ff.C)

	ff = time.NewTicker(time.Second) //want `unstopped ticker`
	print(<-ff.C)
}

func warning4() {
	ff := time.NewTicker(time.Second) //want `unstopped ticker`
	if true {
		ff.Reset(time.Second)
	}

	ff = time.NewTicker(time.Minute + time.Minute) //want `unstopped ticker`
	ff.Reset(time.Second)
}

func warning5() map[time.Time]string {
	f := time.NewTicker(time.Second) //want `unstopped ticker`

	return map[time.Time]string{
		<-f.C: "",
	}
}

func warning6() chan time.Time {
	f := time.NewTicker(time.Second) //want `unstopped ticker`

	var k chan time.Time

	k <- <-f.C

	return k
}

func warning7() {
	f := time.NewTicker(time.Second) //want `unstopped ticker`

	var k chan time.Time

	k <- <-f.C
}

func warning8() time.Time {
	f := time.NewTicker(time.Second) //want `unstopped ticker`
	return <-f.C
}

func warning9() interface{} {
	type timeDecor struct {
		time.Time
	}
	f := time.NewTicker(time.Second) //want `unstopped ticker`

	return timeDecor{<-f.C}
}

func warning10() map[string]time.Time {
	f := time.NewTicker(time.Second) //want `unstopped ticker`

	return map[string]time.Time{
		"": <-f.C,
	}
}
