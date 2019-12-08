package molekcheatez

import (
	"bytes"
	"encoding/binary"
	"io"
)

//Unmarshal processes the bytes of a .solution file.
func Unmarshal(input []byte) (*Solution, error) {
	out := &Solution{}
	reader := bytes.NewReader(input)

	//Get header
	header, err := readint32(reader)
	if err != nil {
		return nil, ErrEOF
	}
	if header != SolutionStart {
		return nil, ErrBadStartHeader
	}

	//PuzzleID
	puzzid, err := readint32(reader)
	if err != nil {
		return nil, ErrEOF
	}
	out.PuzzleID = PuzzleID(puzzid)

	//Handle name of solution
	nameLen, err := readint32(reader)
	if err != nil {
		return nil, ErrEOF
	}
	nameBytes := make([]byte, nameLen)
	_, err = reader.Read(nameBytes)
	if err != nil {
		return nil, ErrEOF
	}
	out.Name = string(nameBytes)

	//Puzzle stats
	solvedHeader, err := readint32(reader)
	if err != nil {
		return nil, ErrEOF
	}
	if solvedHeader == SolvedPuzzle {
		out.Solved = true

		//Get stats
		//there's gotta be a better way to do this
		cyclesHead, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		if cyclesHead != CyclesHeader {
			return nil, ErrBadCycleCountHeader
		}
		cycleCount, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		out.Cycles = &cycleCount

		modulesHead, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		if modulesHead != ModulesHeader {
			return nil, ErrBadModuleCountHeader
		}
		moduleCount, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		out.Modules = &moduleCount

		symbolsHead, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		if symbolsHead != SymbolsHeader {
			return nil, ErrBadSymbolCountHeader
		}
		symbolCount, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		out.Symbols = &symbolCount

	} else if solvedHeader == UnsolvedPuzzle {
		//Setting it false is not required (0 value) but I like the symmetry/
		out.Solved = false
	} else {
		return nil, ErrBadPuzzleSolvedHeader
	}

	//Part count is already encoded in the number of parts, so it can be elided.
	_, err = readint32(reader)
	if err != nil {
		return nil, ErrEOF
	}

	//Parts
	out.Parts = []Part{}
	for {
		part := Part{}

		partType1, err := readint32(reader)
		if err != nil {
			//done gathering parts!
			break
		}
		if partType1 == TypeIsInput1 {
			part.IsInput = true
		} else if partType1 == TypeIsEmitter1 {
			part.IsInput = false
		} else {
			println(partType1)
			return nil, ErrBadPartType1
		}

		//Position and rotation
		posQ, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		if (part.IsInput && (posQ >= BoardSize || -posQ >= BoardSize)) || (!part.IsInput && (posQ > BoardSize || -posQ > BoardSize)) {
			return nil, ErrPartOutOfBounds
		}
		posR, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		if (part.IsInput && (posR >= BoardSize || -posR >= BoardSize)) || (!part.IsInput && (posR > BoardSize || -posR > BoardSize)) {
			return nil, ErrPartOutOfBounds
		}
		rot, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		part.Position = Position{Q: posQ, R: posR}
		part.Rotation = rot

		//Check if it's an input... again... zach why are you like this
		isInputAgainSlice := make([]byte, 1)
		_, err = reader.Read(isInputAgainSlice)
		if err != nil {
			return nil, ErrEOF //how I wish I had a C pre-processor...
		}
		isInputAgain := isInputAgainSlice[0]
		if isInputAgain != TypeIsInput2 && isInputAgain != TypeIsEmitter2 {
			return nil, ErrBadPartType2
		}
		if (isInputAgain == TypeIsInput2) != part.IsInput {
			return nil, ErrMismatchedPartTypes
		}

		//Precursor ID / Arm ID
		precArm, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		if part.IsInput {
			//Precursor ID
			converted := PrecursorID(precArm)
			part.Precursor = &converted
		} else {
			//Arm ID
			part.EmitterID = &precArm
		}
		

		partID, err := readint32(reader)
		if err != nil {
			return nil, ErrEOF
		}
		part.Mystery = partID

		//Opcodes!
		if !part.IsInput {
			part.Instructions = InstructionSeq{}
			for c := 0; c < 24; c++ {
				instrSlice := make([]byte, 1)
				_, err = reader.Read(instrSlice)
				if err != nil {
					return nil, ErrEOF
				}
				part.Instructions = append(part.Instructions, Instruction(instrSlice[0]))
			}
		} else {
			//for some reason, there's 28 bytes of useless info after every input...
			//the first number always seems to be 0x18 then null bytes
			//weird
			discard := make([]byte, 28)
			_, err = reader.Read(discard)
			if err != nil {
				return nil, ErrEOF
			}
		}

		out.Parts = append(out.Parts, part)
	}

	return out, nil
}

//Helper function to read an int32 from a byte slice
func readint32(in io.Reader) (ret int32, err error) {
	data := make([]byte, 4)
	_, err = io.ReadFull(in, data)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}
