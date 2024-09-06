package interpreter

type Univers struct {
	memory  [30000]int8 // array of 30000 cells (maybe implement a way to have infinite cell later)
	pointer uint        // pointer to a cell in the memory
}

var (
	instructionSet []string = []string{
		">", // MovRight
		"<", // MovLeft
		"+", // IncCell
		"-", // DecCell
		".", // Output
		",", // Input
		"[", // JmpPast
		"]", // JmpBack
	}
)

const (
	// BrainF basic instruction set
	MovRight uint = iota
	MovLeft
	IncCell
	DecCell
	Output
	Input
	JmpPast
	JmpBack

	// BrainF variation
)
