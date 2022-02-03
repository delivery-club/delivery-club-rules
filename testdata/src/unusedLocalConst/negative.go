package uselessLocalConst

func negative() {
	const bar = .3
	print(bar)

	const bar2 byte = '1'
	_ = func() {
		print(bar2)
	}
}
