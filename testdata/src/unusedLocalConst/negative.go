package uselessLocalConst

const globalConst = "123"
const GlobalConst = "123"

func negative() {
	const bar = .3
	print(bar)

	const bar2 byte = '1'
	_ = func() {
		print(bar2)
	}
}
