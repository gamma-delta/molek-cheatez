package molekcheatez

import (
	"sort"
)

//because they can't just let you write an anonymous function...
type partsSorterOMatic9000 []Part

func (p partsSorterOMatic9000) Len() int {
	return len(p)
}

func (p partsSorterOMatic9000) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p partsSorterOMatic9000) Less(i, j int) bool {
	//Return true if p[i] goes before p[j]
	p1 := p[i]
	p2 := p[j]

	//Inputs always go first
	if p1.IsInput != p2.IsInput {
		return p1.IsInput //if true, p1 goes first, else p2 goes first
	}

	if p1.IsInput {
		//then both are inputs, so sort by precursor name
		return p1.Precursor.String() < p2.Precursor.String()
	}
	//both are emitters; sort by ID
	return *p1.EmitterID < *p2.EmitterID
	
}

//SortParts sorts this solution's parts in place.
//Inputs are put first in alphabetical order by precursor name
//Arms are next, in numberical order.
func (s *Solution) SortParts() {
	sort.Sort(partsSorterOMatic9000(s.Parts))
}