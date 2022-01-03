package unclosedResource

import (
	"io/ioutil"
	"os"
)

func foo() {
	f, err := os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	if err == nil {
		print(f.Name())
	}

	f, err = os.Open("bar") //want `f.Close() should be deferred right after the os.Open error check`
}

func fooBar() {
	var ff, err = os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the os.Open error check`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `ff.Close() should be deferred right after the ioutil.TempFile error check`
}
