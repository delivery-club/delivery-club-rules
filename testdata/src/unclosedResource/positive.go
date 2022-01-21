package unclosedResource

import (
	"database/sql"
	"io/ioutil"
	"os"
)

func warning1() {
	f, err := os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	if err == nil {
		print(f.Name())
	}

	f, err = os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	print(f.Name())
}

func warning2() {
	f, err := os.Open("bar")
	if err == nil {
		defer f.Close()
	}

	f, err = os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	print(f.Name())
}

func warning3() {
	var ff, err = os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the os.Open error check`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `\Qff.Close() should be deferred right after the ioutil.TempFile error check`
	print(ff.Name())
}

func warning4() {
	ff, err := os.Open("foo.txt") //want `\Qff.Close() should be deferred right after the os.Open error check`
	if err != nil {
		print(ff.Fd())
	}

	ff, err = ioutil.TempFile("/kek", "foo") //want `\Qff.Close() should be deferred right after the ioutil.TempFile error check`
	print(ff.Name())
}

func warning5() {
	f, err := os.Open("bar") //want `\Qf.Close() should be deferred right after the os.Open error check`
	print(f.Name())

	ff, err := os.Open("bar")
	if err == nil {
		defer ff.Close()
	}
}

func warning6() []int {
	db, _ := sql.Open("", "")              //want `\Qdb.Close() should be deferred right after the sql.Open error check`
	var rows, _ = db.QueryContext(nil, "") //want `\Qrows.Close() should be deferred right after the db.QueryContext error check`

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
