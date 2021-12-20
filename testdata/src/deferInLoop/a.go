package deferInLoop

import "time"

func testForStmt() {
	defer println(111)

	for range []int{1, 2, 3, 4} {
		println(222)
		break
	}

	defer println(222)
}

func testRangeStmt() {
	defer println(222)

	for i := 0; i < 10; i++ {
		println(111)
	}

	defer println(222)
}

func testClosure() {
	func() {
		for {
			break
		}
		defer println(1)

		for {
			for {
				break
			}
			break
		}

		defer println(1)
	}()

	func() {
		defer println(123)

		for range []int{1, 2, 3, 4} {
			println(222)
		}

		defer println(123)

		for range []int{1, 2, 3, 4} {
			for range []int{1, 2, 3, 4} {
				println(222)
			}
		}

		defer println(123)
	}()

	for {
		func() {
			defer println(123)
		}()
		break
	}

	for {
		go func() {
			defer println()
		}()

		break
	}
}

func testBlock() {
	{
		for {
			func() {
				defer println()
			}()
			break
		}
	}
	{
		for {
			go func() {
				defer println()
			}()
			break
		}
	}
	{
		for range []int{1, 2, 3, 4} {
			go func() {
				defer println()
			}()
			break
		}
	}
	{
		for range []int{1, 2, 3, 4} {
			{
				func() {
					{
						defer println()
					}
				}()
			}
			break
		}
	}
}

func negativeAssign() {
	x := func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	var xx = func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	var xxx func() = func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	_ = func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}

	var _ = func() {
		{
			defer println(123)
			for {
				break
			}
			defer println(123)
		}
	}

	var _ func() = func() {
		{
			defer println(123)
			for {
				break
			}
			defer println(123)
		}
	}

	var _ = (func() {
		{
			defer println(123)
			for range []int{1, 2, 3} {
			}
			defer println(123)
		}
	})

	x()
	xx()
	xxx()
}

func negativeFuncArgs() {
	time.AfterFunc(time.Second, func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	})

	time.AfterFunc(time.Second, (func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	}))

	{
		time.AfterFunc(time.Second, func() {
			{
				for {
					break
				}
				defer println(123)
			}
		})
	}

	x("").closureExec(func() {
		defer println(123)
		for {
			break
		}
		defer println(123)
	})

	h := x("")
	{
		go h.closureExec(func() {
			{
				defer println(123)
				{
					defer println(123)
					for range []int{1, 2, 3} {
					}
					defer println(123)
				}
			}
		})
	}
}

func deferWithCall() {
	for {
		defer println("test") // want `Possible resource leak, 'defer' is called in the 'for' loop`
		break
	}

	for range []int{1, 2, 3, 4} {
		defer println("test") // want `Possible resource leak, 'defer' is called in the 'for' loop`
	}
}

func deferWithClosure() {
	for {
		defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`

		break
	}

	for range []int{1, 2, 3, 4} {
		defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`
	}
}

func innerLoops() {
	for {
		for {
			defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`

			break
		}
		break
	}

	for {
		for {
			defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`

			break
		}
		break
	}

	for range []int{1, 2, 3, 4} {
		defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`

		for range []int{1, 2, 3, 4} {
			defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`
		}
	}
}

func anonFunc() {
	func() {
		for range []int{1, 2, 3, 4} {
			defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`

			for range []int{1, 2, 3, 4} {
				defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`
			}
		}
	}()

	for {
		defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
		break
	}

	func() {
		for {
			defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`

			for range []int{1, 2, 3, 4} {
				defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`
			}

			break
		}

		for {
			defer func() {}() // want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
	}()

	go func() {
		for {
			defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
	}()
}

func assignStmt() {
	f, ff := func() {
		{
			defer println(123)
			{
				defer println(123)
				for i := 0; i < 10; {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
					i++
				}
				defer println(123)
			}
		}
	}, func() {
		{
			i := 0
			defer println(123)
			{
				defer println(123)
				for i < 10 {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
					i++
				}
				defer println(123)
			}
		}
	}

	fff := func() {
		i := 0
		{
			defer println(123)
			{
				defer println(123)
				for ; i < 10; i++ {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				}
				defer println(123)
			}
		}
	}

	var t = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				}
				defer println(123)
			}
		}
	}

	var tt func() = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				}
				defer println(123)
			}
		}
	}

	var _ = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				}
				defer println(123)
			}
		}
	}

	_ = func() {
		{
			defer println(123)
			{
				defer println(123)
				for range []int{1, 2, 3} {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				}
				defer println(123)
			}
		}
	}

	var ttt = (func() {
		{
			defer println(123)
			for x := 0; x < 5; x++ {
				defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
			}
			defer println(123)
		}
	})

	var _ = (func() {
		{
			defer println(123)
			for range []int{1, 2, 3, 4} {
				defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
			}
			defer println(123)
		}
	})

	f()
	ff()
	fff()
	t()
	tt()
	ttt()
}

func contextBlock() {
	{
		go func() {
			for {
				defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				break
			}
		}()
	}

	go (func() {
		for {
			defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
	})()

	{
		func() {
			for {
				defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
				break
			}
		}()
	}

	{
		{
			func() {
				for {
					defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
					break
				}
			}()
		}
	}

	{
		for {
			defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
	}
}

type x string

func funcArgs() {
	time.AfterFunc(time.Second, func() {
		for {
			defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
		}
	})

	x("").closureExec(func() {
		defer println(123)
		for {
			defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
		defer println(123)
	})

	h := x("")
	{
		go h.closureExec(func() {
			{
				defer println(123)
				{
					defer println(123)
					for range []int{1, 2, 3} {
						defer println(123) // want `Possible resource leak, 'defer' is called in the 'for' loop`
					}
					defer println(123)
				}
			}
		})
	}
}

func (x x) closureExec(f func()) { f() }

func falsePositive() {
	// TODO: now we cant find this cases due to additional block scope
	for {
		{
			defer println(123) // TODO: want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
	}

	for range []int{1, 2, 3} {
		{
			defer println(123) // TODO: want `Possible resource leak, 'defer' is called in the 'for' loop`
			break
		}
	}

	go func() {
		{
			for {
				{
					defer println(123) // TODO: want `Possible resource leak, 'defer' is called in the 'for' loop`
					break
				}
			}
		}
	}()
	{
		time.AfterFunc(time.Second, func() {
			{
				for {
					{
						defer println(123) // TODO: want `Possible resource leak, 'defer' is called in the 'for' loop`
						break
					}
				}
				defer println(123)
			}
		})
	}
}
