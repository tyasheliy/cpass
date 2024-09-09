package entry

import "sort"

type Aggregate struct {
	entries []Entry
}

func NewAggregate(entries ...Entry) *Aggregate {
	return &Aggregate{
		entries: entries,
	}
}

func (a *Aggregate) Slice() []Entry {
	return a.entries
}

func (a *Aggregate) Append(entries ...Entry) {
	a.entries = append(a.entries, entries...)
}

func (a *Aggregate) SortByPassName() *Aggregate {
	passNameMap := make(map[string]Entry)
	passNames := make([]string, len(a.entries))

	for i, entry := range a.entries {
		passName := PassName(entry)
		passNameMap[passName] = entry
		passNames[i] = passName
	}

	sort.Strings(passNames)

	for i := range passNames {
		passName := passNames[i]
		entry := passNameMap[passName]
		a.entries[i] = entry
	}

	return a
}
