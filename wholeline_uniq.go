package uniq2

/*
WholeLineUniqer is an implementation of uniqer for whole lines.
*/
type WholeLineUniqer struct {
	lines map[string]bool
}

/*
NewWholeLineUniqer creates an instance of WholeLineUniqer.
*/
func NewWholeLineUniqer() *WholeLineUniqer {
	return &WholeLineUniqer{lines: map[string]bool{}}
}

/*
StreamLine tries to remove the duplicated line.
*/
func (wlu *WholeLineUniqer) StreamLine(line string) (isUniq bool) {
	_, ok := wlu.lines[line]
	if !ok {
		wlu.lines[line] = true
	}
	return !ok
}
