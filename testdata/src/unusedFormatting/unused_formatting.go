package unusedFormatting

import (
	"fmt"
	"os"
	// modules not supported by analysistest: https://github.com.cnpmjs.org/golang/go/issues/46041, waiting for merge
	//"github.com/pkg/errors"
)

func fmtUnusedFormatting() {
	fmt.Sprintf("foo")           // want "use function alternative without formatting"
	fmt.Printf("bar")            // want "use function alternative without formatting"
	fmt.Fscanf(os.Stdin, "bar")  // want "use function alternative without formatting"
	fmt.Fprintf(os.Stdin, "bar") // want "use function alternative without formatting"
	fmt.Scanf("foo")             // want "use function alternative without formatting"
	fmt.Sscanf("foo", "bar")     // want "use function alternative without formatting"
	//errors.Errorf("on foo")                             // "use function alternative without formatting"
	//errors.WithMessagef(errors.New("on bar"), "on foo") // "use function alternative without formatting"
	//errors.Wrapf(errors.New("on bar"), "on foo")        // "use function alternative without formatting"

	fmt.Errorf("bar happend") // want "use errors.New instead"
}

func fmtFormatting() {
	fmt.Sprintf("foo %s", "bar")
	fmt.Printf("foo %s", "bar")
	fmt.Fscanf(os.Stdin, "foo %s", "bar")
	fmt.Fprintf(os.Stdin, "foo %s", "bar")
	fmt.Scanf("foo %s", "bar")
	fmt.Errorf("kek happend %s", "foo")
	fmt.Sscanf("foo bar", "foo %s", "bar")

	//err := errors.New("on foo")
	//errors.WithMessagef(err, "on bar: id: %s", "foo")
	//errors.Errorf("on foo: %s", "on bar")
	//errors.Wrapf(err, "on foo: id: %s", "bar")
}
