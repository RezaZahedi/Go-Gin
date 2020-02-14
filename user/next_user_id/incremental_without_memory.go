package next_user_id

// NewIncrementalIDWithNoMemory is used to generate new User IDs. it is used so that other
// UUID generation algorithms can be replaced with the current one
func NewIncrementalIDWithNoMemory() func() int {
	var i int = 2
	return func() int {
		i++
		return i
	}
}
