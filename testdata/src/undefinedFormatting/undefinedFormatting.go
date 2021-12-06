package undefinedFormatting

import (
	"fmt"
	//"github.com/pkg/errors"
)

func fmtUnusedFormatting() {
	fmt.Sprintf("foo") // want "use function alternative without formatting"
	// TODO: modules not supported by analysistest: https://github.com.cnpmjs.org/golang/go/issues/46041, waiting for merge
	//errors.Errorf("on foo")                             // "use function alternative without formatting"
	//errors.WithMessagef(errors.New("on bar"), "on foo") // "use function alternative without formatting"
	//errors.Wrapf(errors.New("on bar"), "on foo")        // "use function alternative without formatting"

	foo := func(s string) string { return s + "123" }
	bar := "foo bar"

	fmt.Errorf(foo("bar")) // want `\Quse errors.New(foo("bar")) or fmt.Errorf("%s", foo("bar")) instead`
	fmt.Errorf(bar)        // want `\Quse errors.New(bar) or fmt.Errorf("%s", bar) instead`
}

func fmtFormatting() {
	fmt.Sprintf("foo %s", "bar")
	fmt.Errorf("kek happend %s", "foo")
	fmt.Errorf("kek happend")

	foo := "123"
	fmt.Errorf("%s", foo)

	//err := errors.New("on foo")
	//errors.WithMessagef(err, "on bar: id: %s", "foo")
	//errors.Errorf("on foo: %s", "on bar")
	//errors.Wrapf(err, "on foo: id: %s", "bar")
}
