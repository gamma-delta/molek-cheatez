package molekcheatez

import (
	"bytes"
	"encoding/binary"
)

//UnsafeMarshal marshalls a solution to bytes without doing any checks on it.
//This will let you put parts out of bounds, for instance.
//TODO: Safe marshal function
func (s *Solution) UnsafeMarshal() []byte {
	out := new(bytes.Buffer)

	//Write header
	writeint32(out, SolutionStart)

	//Puzzle ID
	writeint32(out, int32(s.PuzzleID))

	//Do name of solution
	nameLen := len(s.Name)
	writeint32(out, int32(nameLen)) //the length...
	out.WriteString(s.Name)         //...and the name

	//Puzzle stats
	if s.Solved {
		writeint32(out, SolvedPuzzle)
		//Write stats
		writeint32(out, CyclesHeader)
		writeint32(out, *s.Cycles)
		writeint32(out, ModulesHeader)
		writeint32(out, *s.Modules)
		writeint32(out, SymbolsHeader)
		writeint32(out, *s.Symbols)
	} else {
		//No need to write stats
		writeint32(out, UnsolvedPuzzle)
	}

	//Write part count
	writeint32(out, int32(len(s.Parts)))

	//Writing the parts
	for _, part := range s.Parts {
		if part.IsInput {
			writeint32(out, TypeIsInput1)
		} else {
			writeint32(out, TypeIsEmitter1)
		}

		//Position and rotation
		writeint32(out, part.Position.Q)
		writeint32(out, part.Position.R)
		writeint32(out, part.Rotation)

		//Whether it's an input, again
		if part.IsInput {
			//Write header
			out.WriteByte(TypeIsInput2)
			//Write precursor ID
			writeint32(out, int32(*part.Precursor))
		} else {
			//Write header
			out.WriteByte(TypeIsEmitter2)
			//Write emitter ID
			writeint32(out, *part.EmitterID)
		}

		//Magic mstery who knows aaaaa
		writeint32(out, part.Mystery)

		//Opcodes
		if part.IsInput {
			//jk write the magic data
			out.Write(EmitterPad)
		} else {
			//opcodes
			for _, instr := range part.Instructions {
				out.WriteByte(byte(instr))
			}
		}
	}

	return out.Bytes()
}

//writeint32 helps me write ints as byte slices
func writeint32(b *bytes.Buffer, i int32) {
	binary.Write(b, binary.LittleEndian, i)
}
