package uniq2

/*
AdjacentUniqer is an implementation of uniqer for adjacent lines.
*/
type AdjacentUniqer struct {
	prev      string
	firstLine bool
}

/*
NewAdjacentUniqer creates an instance of AdjacentUniqer.
*/
func NewAdjacentUniqer() *AdjacentUniqer {
	return &AdjacentUniqer{prev: "", firstLine: true}
}

/*
StreamLine tries to remove the duplicated line.
*/
func (au *AdjacentUniqer) StreamLine(line string) bool {
	if au.firstLine {
		au.prev = line
		au.firstLine = false
		return true
	}
	if au.prev == line {
		return false
	}
	au.prev = line
	return true
}
