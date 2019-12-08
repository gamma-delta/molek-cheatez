package molekcheatez

import (
	"errors"
)

//Solution is one solution to a puzzle.
type Solution struct {
	PuzzleID     PuzzleID `json:"puzzle_id"`
	Name         string   `json:"name"`
	Solved       bool     `json:"solved"`
	Cycles       *int32   `json:"cycles,omitempty"`  //Nil if Solved == false
	Modules      *int32   `json:"modules,omitempty"` //Nil if Solved == false
	Symbols      *int32   `json:"symbols,omitempty"`                   //Nil if Solved == false

	Parts []Part `json:"parts"`
}

//Part is either an emitter or a input.
type Part struct {
	IsInput   bool `json:"is_input"`
	Position  Position `json:"position"`
	Rotation  int32 `json:"rotation"`
	Precursor *PrecursorID `json:"precursor,omitempty"` //Nil if IsInput == False
	EmitterID *int32 `json:"emitterid,omitempty"` //Nil if IsInput == True

	Mystery       int32 `json:"mystery"`         //I really don't know
	Instructions InstructionSeq `json:"instructions,omitempty"` //Is nil if IsInput == true
}

//Position is the spot something is on the board.
type Position struct {
	Q int32 `json:"q"` //Hexes right of center
	R int32 `json:"r"` //Hexes up-right of center
}

//InstructionSeq is the 24 instructions an emitter can perform.
type InstructionSeq []Instruction

//MarshalJSON makes the JSON output nicer
func (is InstructionSeq) MarshalJSON() ([]byte, error) {
	out := []byte{'"'}
	for _, instr := range is {
		out = append(out, byte(instr.Rune()))
	}
	out = append(out, '"')

	return out, nil
}

//UnmarshalJSON takes the string provided by JSON and unmarshals it
func (is *InstructionSeq) UnmarshalJSON(b []byte) error {
	out := make(InstructionSeq, 24)
	str := string(b[1:len(b)-1])
	if len(str) != 24 {
		return errors.New("instruction sequence wrong length")
	}
	for c, char := range str {
		instr, ok := runeToInstr[char]
		if !ok {
			return errors.New("invalid byte in instruction sequence")
		}
		out[c] = instr
	}

	*is = out
	return nil
}