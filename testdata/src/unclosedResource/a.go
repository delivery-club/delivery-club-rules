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

	f, err = os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	print(f.Name())
}

func foo2() {
	f, err := os.Open("bar")
	if err == nil {
		defer f.Close()
	}

	f, err = os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	print(f.Name())
}

func fooBar() {
	var ff, err = os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the os.Open error check`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `\Qff.Close() should be deferred right after the ioutil.TempFile error check`
	print(ff.Name())
}

func fooBar2() {
	ff, err := os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the os.Open error check`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `\Qff.Close() should be deferred right after the ioutil.TempFile error check`
	print(ff.Name())
}

func negative() {
	ff, err := ioutil.TempFile("/fo", "bo")
	if err != nil {
		print(err)
	}
	defer ff.Close()

	ff, err = ioutil.TempFile("/fo", "bo")
	if err != nil {
		print(err)
	}
	defer func() {
		print(123)
		ff.Close()
	}()
}

func falseNegative() {
	f, err := os.Open("bar") // TODO must be warning
	print(f.Name())

	f, err = os.Open("bar")
	if err == nil {
		defer f.Close()
	}
}
