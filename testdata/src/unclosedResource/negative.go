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

var globalVar *os.File

func negative2() {
	globalVar, _ = ioutil.TempFile("", "") // false negative because "_" doesn't have error type
	kk := globalVar.Name()

	print(kk)
}

func negativeGlobalVar() {
	var err error
	globalVar, err = ioutil.TempFile("", "") // global var
	kk := globalVar.Name()

	print(kk, err)
}

func negative3() *os.File {
	file, _ := ioutil.TempFile("", "") // var escape the function

	return file
}

//TODO: false positive `files` var escape the function
func negative4() []*os.File {
	var files []*os.File
	file, _ := ioutil.TempFile("", "") // want `\Qfile.Close() should be deferred right after the resource creation`

	files = append(files, file)
	return files
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

func negative9() {
	closure := func() (*os.File, error) {
		return nil, nil
	}

	f, _ := closure()
	f.Name()
	defer f.Close()
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

type MyStruct struct {
	f *os.File
}

func (m MyStruct) negative10() {
	kk, _ := os.Open("123")
	m.f = kk
}

func negative11() *os.File {
	var k *os.File
	k, _ = os.Open("123")

	kk := k

	return kk
}
