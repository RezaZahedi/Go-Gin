package next_user_id

func NewIncrementalIDWithNoMemory() func() int {
	var i int = 1
	return func() int {
		i++
		return i
	}
}
