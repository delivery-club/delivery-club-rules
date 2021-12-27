package regexpCompileInLoop

import (
	"regexp"
)

func warnings() {
	for {
		r := regexp.MustCompile(`qwe`) //want "don't compile regex in the loop, move it outside of the loop"

		r.Match([]byte{})
		break
	}

	var (
		ok  bool
		err error
	)

	for i := 0; i < 10; i++ {
		print(123)

		ok, err = regexp.Match(`qwe`, []byte{}) //want "don't compile regex in the loop, move it outside of the loop"
		if !ok {
			print(err)
		}
	}

	for range []string{"1", "2"} {
		ok, err = regexp.MatchReader(`qwe`, nil) //want "don't compile regex in the loop, move it outside of the loop"
		if !ok {
			print(err)
		}
	}
}
