package types

// TeammateTypes contain known player types from the same team
type TeammateTypes map[int]int

// DefaultTeammateTypes returns TeammateTypes object with default values
func DefaultTeammateTypes() TeammateTypes {
	t := make(TeammateTypes)

	for i := 1; i <= 11; i++ {
		t[i] = 0
	}

	return t
}
