package cli_table

// Stringer is a type constraint that requires the type argument to have
// a String method and permits the generic function to call String.
// The String method should return a string representation of the value.
type Stringer interface {
	String() string
}

type Cells struct {
	data   []string
	height int
	width  int
}

func (c Cells) String() string {
	//TODO implement me
	panic("implement me")
}
