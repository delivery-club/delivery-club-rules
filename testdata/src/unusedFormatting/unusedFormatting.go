package unusedFormatting

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
}

func fmtFormatting() {
	fmt.Sprintf("foo %s", "bar")
	fmt.Sprint("foo", "bar")

	//errors.WithMessagef(err, "on bar: id: %s", "foo")
	//errors.Errorf("on foo: %s", "on bar")
	//errors.Wrapf(err, "on foo: id: %s", "bar")
}
