package trainerclient

// command is the struct used to move and parse RCSS commands
type command struct {
	time      int
	cmdString string
}

// String implements the native Stringer interface
func (c command) String() string {
	return string(c.cmdString)
}
