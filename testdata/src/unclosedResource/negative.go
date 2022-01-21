package unclosedResource

import (
	"database/sql"
	"io/ioutil"
	"os"
)

func negative1() {
	ff, err := ioutil.TempFile("/fo", "bo")
	if err != nil {
		print(err)
	}
	defer ff.Close()

	ff, err = ioutil.TempFile("/fo", "bo")
	if err != nil {
		print(err)
	}
	print(123)
	ff.Close()
}

var f *os.File

func negative2() {
	f, _ = ioutil.TempFile("", "") // global var
}

func negative3() *os.File {
	file, _ := ioutil.TempFile("", "") // var escape the function

	return file
}

func negative4() {
	var files []*os.File
	file, _ := ioutil.TempFile("", "") // var escape the function in another var

	files = append(files, file)
}

func negative5() {
	var filesMap map[string]*os.File
	file, _ := ioutil.TempFile("", "") // var escape the function in another var

	filesMap[file.Name()] = file
}

func negative6() {
	type st struct {
		*os.File
	}
	var (
		fileDecorator1 st
		fileDecorator2 st
	)
	file, _ := ioutil.TempFile("", "") // var escape the function in another var

	fileDecorator1 = st{file}
	fileDecorator2 = st{
		File: file,
	}

	kk, kkk := fileDecorator1.Name(), fileDecorator2.Name()

	print(kk, kkk)
}

func negative7() {
	var ch chan *os.File
	file, _ := ioutil.TempFile("", "") // var escape the function in another var

	ch <- file
}

func negative8() []int {
	db, _ := sql.Open("", "")
	defer db.Close()

	rows, _ := db.QueryContext(nil, "")

	var (
		i      int
		result []int
	)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&i)
		result = append(result, i)
	}

	return result
}

// test for: https://github.com/quasilyte/go-ruleguard/issues/366
func dataRace() {
	f, err := os.Open("bar")
	print(f.Name())

	f, err = os.Open("bar")
	if err == nil {
		defer f.Close()
	}
}
