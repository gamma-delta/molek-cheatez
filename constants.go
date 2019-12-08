package molekcheatez

import (
	"errors"
)

//These are stored little-endian, but go source code is big-endian.
const (
	SolutionStart = 0x00002727 //Solutions always start with this

	//Tells whether if the puzzle is solved or not
	UnsolvedPuzzle = 0x00000000
	SolvedPuzzle   = 0x00000003

	//If the puzzle is solved, the 3 stats have an int before them like a header.
	//They're always in the same order, so I'm not sure why that's a thing.
	//Anyway I should probably check for them anyways.
	CyclesHeader  = 0x00000000
	ModulesHeader = 0x00000001
	SymbolsHeader = 0x00000002

	//The type of the part is in 3 places.
	//In the solution file, the part starts with either 1 (input) or 3 (emitter). This is #1 here.
	//The 17th byte is a bool with true being IsInput. This is #2 here.
	//Finally, the opcodes for inputs are always 0x18 then 23 * 0x00, and the Part ID of arms are always 0x18. (I think)
	//If these don't match somehow I should probably throw an enormous error, like how tf did you manage that
	TypeIsInput1   = 0x00000001
	TypeIsEmitter1 = 0x00000003
	TypeIsInput2   = 0x01
	TypeIsEmitter2 = 0x00

	//Board size, in hexes.
	//abs(Emitters' positions) can be equal to this, but you should keep parts below it.
	BoardSize = 7
)

//EmitterPad is the 28 bytes present in every emitter after the actual useful data.
//I have no idea why this is.
var EmitterPad = []byte{0x18, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}

//Errors
var (
	ErrBadStartHeader        = errors.New(`solution did not start with ''`)
	ErrEOF                   = errors.New(`solution ended too soon`)
	ErrBadPuzzleSolvedHeader = errors.New(`invalid puzzle solved status`)
	ErrBadCycleCountHeader   = errors.New(`invalid cycle count header`)
	ErrBadModuleCountHeader  = errors.New(`invalid module count header`)
	ErrBadSymbolCountHeader  = errors.New(`invalid symbol count header`)
	ErrBadPartType1          = errors.New(`first part type is invalid`)
	ErrBadPartType2          = errors.New(`second part type is invalid`)
	ErrMismatchedPartTypes   = errors.New(`part does not agree with itself on what type it is`)
	ErrPartOutOfBounds       = errors.New(`part is out of bounds`)
)
