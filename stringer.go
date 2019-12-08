package molekcheatez

import (
	"fmt"
	"strings"
)

//String makes Solution implement Stringer
func (s Solution) String() string {
	out := strings.Builder{}

	out.WriteString(fmt.Sprintf("%s: %s\n", s.PuzzleID, s.Name))
	out.WriteString(fmt.Sprintf("%d parts\n", len(s.Parts)))

	if s.Solved {
		out.WriteString(fmt.Sprintf("Solved. Cycles: %d; Modules: %d; Symbols: %d\n", *s.Cycles, *s.Modules, *s.Symbols))
	} else {
		out.WriteString("Unsolved.\n")
	}


	for c, part := range s.Parts {
		out.WriteString(part.String())
		if c < len(s.Parts)-1 {
			//on all but the last
			out.WriteString("\n---\n")
		}
	}

	return out.String()
}

//String makes Part implement Stringer
func (p Part) String() string {
	out := strings.Builder{}

	if p.IsInput {
		out.WriteString(fmt.Sprintf("Input %s", *p.Precursor))
	} else {
		out.WriteString(fmt.Sprintf("Emitter #%d", *p.EmitterID))
	}
	out.WriteString(fmt.Sprintf(" @ (%d, %d) Rot %d", p.Position.Q, p.Position.R, p.Rotation))
	out.WriteString(fmt.Sprintf(" #%d", p.Mystery))

	if !p.IsInput {
		//Instructions
		out.WriteString(" [")
		for _, instr := range p.Instructions {
			out.WriteRune(instr.Rune())
		}
		out.WriteRune(']')
	}

	return out.String()
}

//Rune gives one rune representing an instruction for brevity
func (i Instruction) Rune() rune {
	return rune(instructionRunes[i])
}

//Reference string for above
const instructionRunes = `.<>^v\/+-x#~`
var runeToInstr map[rune]Instruction
func init() {
	runeToInstr = make(map[rune]Instruction, len(instructionRunes))

	for c, r := range instructionRunes {
		runeToInstr[r] = Instruction(c)
	}
}