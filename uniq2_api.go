package uniq2

import (
	"strings"
)

/*
Uniqer removes duplicated lines of given files by certain algorithm.
*/
type Uniqer interface {
	/*
	   StreamLine pours lines from reader and returns the given line should show or not show.
	*/
	StreamLine(line string) (uniqFlag bool)
}

/*
InverseUniqer negates the result from the other Uniqer.
*/
type InverseUniqer struct {
	uniqer Uniqer
}

/*
StreamLine tries to remove the duplicated line.
*/
func (iu *InverseUniqer) StreamLine(line string) bool {
	return !iu.uniqer.StreamLine(line)
}

/*
FilterUniqer composes Filter and Uniqer.
*/
type FilterUniqer interface {
	StreamLine(line string) (uniqFlag bool)
	Filter(line string) string
}

/*
BasicFilterUniqer is the default implementation of FilterUniqer.
*/
type BasicFilterUniqer struct {
	filter Filter
	uniqer Uniqer
}

/*
Filter filters given string.
*/
func (bfu *BasicFilterUniqer) Filter(line string) string {
	return bfu.filter.Filter(line)
}

/*
StreamLine tries to remove the duplicated line.
*/
func (bfu *BasicFilterUniqer) StreamLine(line string) (uniqFlag bool) {
	return bfu.uniqer.StreamLine(bfu.Filter(line))
}

/*
Filter is an interface for filtering given line.
*/
type Filter interface {
	Filter(line string) string
}

/*
Parameters represents option parameter values.
*/
type Parameters struct {
	Adjacent    bool
	ShowCounts  bool
	DeleteLines bool
	IgnoreCase  bool
}

func (params *Parameters) String() string {
	types := []string{}
	if params.Adjacent {
		types = append(types, "adjacent")
	}
	if params.ShowCounts {
		types = append(types, "show-counts")
	}
	if params.DeleteLines {
		types = append(types, "delete-lines")
	}
	if params.IgnoreCase {
		types = append(types, "ignore-case")
	}
	return strings.Join(types, ",")
}

/*
BuildUniqer creates suitable Uniqer following params, the receiver.
*/
func (params *Parameters) BuildUniqer() Uniqer {
	filter := params.buildFilter()
	uniqer := createUniqer(params.Adjacent)
	if params.DeleteLines {
		uniqer = &InverseUniqer{uniqer: uniqer}
	}
	return &BasicFilterUniqer{filter: filter, uniqer: uniqer}
}

func createUniqer(adjacent bool) Uniqer {
	if adjacent {
		return NewAdjacentUniqer()
	}
	return NewWholeLineUniqer()
}

func (params *Parameters) buildFilter() Filter {
	filters := []Filter{}
	if params.IgnoreCase {
		filters = append(filters, &IgnoreCaseFilter{})
	}
	if params.ShowCounts {
		filters = append(filters, &CountLineFilter{counts: map[string]int{}})
	}
	return &MultipleFilter{filters: filters}
}

/*
MultipleFilter contains multiple filters, and apply filters by the order.
*/
type MultipleFilter struct {
	filters []Filter
}

/*
Filter filters given string.
*/
func (mf *MultipleFilter) Filter(line string) string {
	for _, filter := range mf.filters {
		line = filter.Filter(line)
	}
	return line
}

/*
CountLineFilter counts lines.
*/
type CountLineFilter struct {
	counts map[string]int
}

/*
Filter filters given string.
*/
func (clf *CountLineFilter) Filter(line string) string {
	clf.counts[line] = clf.counts[line] + 1
	return line
}

/*
Counts returns count of the line.
*/
func (clf *CountLineFilter) Counts(line string) int {
	return clf.counts[line]
}

/*
IgnoreCaseFilter shows an filter of ignoring case.
*/
type IgnoreCaseFilter struct {
}

/*
Filter filters given string.
*/
func (icf *IgnoreCaseFilter) Filter(line string) string {
	return strings.ToLower(line)
}
