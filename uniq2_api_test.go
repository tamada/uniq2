package uniq2

import "testing"

func TestParametersString(t *testing.T) {
	testdata := []struct {
		giveParams *Parameters
		wontString string
	}{
		{&Parameters{Adjacent: false, ShowCounts: false, DeleteLines: false, IgnoreCase: false}, ""},
		{&Parameters{Adjacent: true, ShowCounts: false, DeleteLines: false, IgnoreCase: false}, "adjacent"},
		{&Parameters{Adjacent: false, ShowCounts: true, DeleteLines: false, IgnoreCase: false}, "show-counts"},
		{&Parameters{Adjacent: false, ShowCounts: false, DeleteLines: true, IgnoreCase: false}, "delete-lines"},
		{&Parameters{Adjacent: false, ShowCounts: false, DeleteLines: false, IgnoreCase: true}, "ignore-case"},
		{&Parameters{Adjacent: true, ShowCounts: true, DeleteLines: false, IgnoreCase: false}, "adjacent,show-counts"},
		{&Parameters{Adjacent: true, ShowCounts: true, DeleteLines: true, IgnoreCase: false}, "adjacent,show-counts,delete-lines"},
		{&Parameters{Adjacent: true, ShowCounts: true, DeleteLines: true, IgnoreCase: true}, "adjacent,show-counts,delete-lines,ignore-case"},
		{&Parameters{Adjacent: false, ShowCounts: true, DeleteLines: true, IgnoreCase: false}, "show-counts,delete-lines"},
		{&Parameters{Adjacent: false, ShowCounts: true, DeleteLines: true, IgnoreCase: true}, "show-counts,delete-lines,ignore-case"},
		{&Parameters{Adjacent: false, ShowCounts: false, DeleteLines: true, IgnoreCase: true}, "delete-lines,ignore-case"},
	}
	for _, td := range testdata {
		gotString := td.giveParams.String()
		if td.wontString != gotString {
			t.Errorf("Parameters(%s).String() did not match, wont %s, got %s", td.giveParams.String(), td.wontString, gotString)
		}
	}
}

func TestUniq2BuildUniqer(t *testing.T) {
	testdata := []struct {
		params         *Parameters
		filterSize     int
		inverseUniqer  bool
		adjacentUniqer bool
		wholeUniqer    bool
	}{
		{&Parameters{}, 0, false, false, true},
		{&Parameters{Adjacent: true}, 0, false, true, false},
		{&Parameters{DeleteLines: true}, 0, true, false, false},
		{&Parameters{ShowCounts: true, Adjacent: true}, 1, false, true, false},
		{&Parameters{IgnoreCase: true, ShowCounts: true, Adjacent: true}, 2, false, true, false},
	}
	for _, td := range testdata {
		uniqer, _ := td.params.BuildUniqer().(*BasicFilterUniqer)
		filter, _ := uniqer.filter.(*MultipleFilter)
		if len(filter.filters) != td.filterSize {
			t.Errorf("filter size of buildUniqer({ %s }) did not match, wont %d, got %d", td.params.String(), td.filterSize, len(filter.filters))
		}
		_, inverseFlag := uniqer.uniqer.(*InverseUniqer)
		_, adjacentFlag := uniqer.uniqer.(*AdjacentUniqer)
		_, wholeLineFlag := uniqer.uniqer.(*WholeLineUniqer)
		if inverseFlag != td.inverseUniqer {
			t.Errorf("type error InverseUniqer by buildUniqer({ %s })", td.params.String())
		}
		if adjacentFlag != td.adjacentUniqer {
			t.Errorf("type error AdjacentUniqer by buildUniqer({ %s })", td.params.String())
		}
		if wholeLineFlag != td.wholeUniqer {
			t.Errorf("type error InverseUniqer by buildUniqer({ %s })", td.params.String())
		}
	}
}
