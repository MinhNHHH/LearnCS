package exercises

import (
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

const (
	eq int = iota
	lt
	gt
)

type customSort struct {
	t       []*Track
	columns []ColumnCmp
}

func (x customSort) CompareYear(a, b *Track) int {
	switch {
	case a.Year > b.Year:
		return gt
	case a.Year < b.Year:
		return lt
	}
	return eq
}

func (x customSort) CompareLength(a, b *Track) int {
	switch {
	case a.Length > b.Length:
		return gt
	case a.Length < b.Length:
		return lt
	}
	return eq
}

func (x customSort) CompareTitle(a, b *Track) int {
	switch {
	case a.Title > b.Title:
		return gt
	case a.Title < b.Title:
		return lt
	}
	return eq
}

type ColumnCmp func(a, b *Track) int

func (c *customSort) Select(cmp ColumnCmp) {
	c.columns = append([]ColumnCmp{cmp}, c.columns...)
}

func (x customSort) Len() int { return len(x.t) }
func (x customSort) Less(i, j int) bool {
	for _, f := range x.columns {
		cmp := f(x.t[i], x.t[j])
		switch cmp {
		case gt:
			return false
		case lt:
			return true
		case eq:
			return false
		}
	}
	return false
}
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func NewCustomSort(tracks []*Track) *customSort {
	return &customSort{t: tracks}
}
