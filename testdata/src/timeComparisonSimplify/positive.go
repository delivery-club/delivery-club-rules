package timeComparisonSimplify

import (
	"time"
)

func _(x, y time.Time) {
	if !x.Before(y) { // want `\Qsuggestion: x.After(y)`
		print(42)
	}

	if !y.After(x) { // want `\Qsuggestion: y.Before(x)`
		print(51)
	}

	print(!y.Before(x)) // want `\Qsuggestion: y.After(x)`
}
