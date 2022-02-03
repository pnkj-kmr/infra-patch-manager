package entity

// MultiEntity - new type for sorting purpose
type MultiEntity []Entity

// Len is part of sort.Interface.
func (f MultiEntity) Len() int {
	return len(f)
}

// Swap is part of sort.Interface.
func (f MultiEntity) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (f MultiEntity) Less(i, j int) bool {
	return f[i].ModTime().Before(f[j].ModTime())
}
