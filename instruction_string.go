// Code generated by "stringer -type=Instruction"; DO NOT EDIT.

package molekcheatez

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoInstruction-0]
	_ = x[SlideLeft-1]
	_ = x[SlideRight-2]
	_ = x[Push-3]
	_ = x[Pull-4]
	_ = x[RotateLeft-5]
	_ = x[RotateRight-6]
	_ = x[AddHydrogen-7]
	_ = x[RemoveHydrogen-8]
	_ = x[Delete-9]
	_ = x[Output-10]
	_ = x[ShuntHydrogen-11]
}

const _Instruction_name = "NoInstructionSlideLeftSlideRightPushPullRotateLeftRotateRightAddHydrogenRemoveHydrogenDeleteOutputShuntHydrogen"

var _Instruction_index = [...]uint8{0, 13, 22, 32, 36, 40, 50, 61, 72, 86, 92, 98, 111}

func (i Instruction) String() string {
	if i >= Instruction(len(_Instruction_index)-1) {
		return "Instruction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Instruction_name[_Instruction_index[i]:_Instruction_index[i+1]]
}
