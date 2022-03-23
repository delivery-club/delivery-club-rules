package unclosedResource

import (
	"database/sql"
	"io/ioutil"
	"os"
)

func warning1() {
	f, err := os.Open("bar") //want `\Qf.Close() should be deferred right after the resource creation`
	if err == nil {
		print(f.Name())
	}

	f, err = os.Open("bar") //want `\Qf.Close() should be deferred right after the resource creation`
	print(f.Name())
}

func warning2() {
	f, err := os.Open("bar")
	if err == nil {
		defer f.Close()
	}

	f, err = os.Open("bar") //want `\Qf.Close() should be deferred right after the resource creation`
	print(f.Name())
}

func warning3() {
	var ff, err = os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the resource creation`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `\Qff.Close() should be deferred right after the resource creation`
	print(ff.Name())
}

func warning4() {
	ff, err := os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the resource creation`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `\Qff.Close() should be deferred right after the resource creation`
	print(ff.Name())
}

func warning5() {
	f, err := os.Open("bar") //want `\Qf.Close() should be deferred right after the resource creation`
	print(f.Name())

	ff, err := os.Open("bar")
	if err == nil {
		defer ff.Close()
	}
}

func warning6() []int {
	db, _ := sql.Open("", "")              //want `\Qdb.Close() should be deferred right after the resource creation`
	var rows, _ = db.QueryContext(nil, "") //want `\Qrows.Close() should be deferred right after the resource creation`

	var (
		i      int
		result []int
	)

	for rows.Next() {
		_ = rows.Scan(&i)
		result = append(result, i)
	}

	return result
}

func warning7() {
	closure := func() (*os.File, error) {
		return nil, nil
	}

	f, _ := closure() // want `\Qf.Close() should be deferred right after the resource creation`
	f.Name()
}

type MyStr struct {
	string
}
type Creator struct{}

func (m MyStr) Close() error {
	return nil
}

func (m MyStr) String() string {
	return m.string
}

func (Creator) OpenMyStruct() (MyStr, error) {
	return MyStr{}, nil
}

func NewCreator() Creator { return Creator{} }

func warning8() {
	m, _ := NewCreator().OpenMyStruct() // want `\Qm.Close() should be deferred right after the resource creation`

	print(m.String())
}
