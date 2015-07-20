package phineas

type Resources []Resource

func (s Resources) Len() int {
	return len(s)
}

func (s Resources) Less(i, j int) bool {
	return len(s[i].BuildPath()) < len(s[j].BuildPath())
}

func (s Resources) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
