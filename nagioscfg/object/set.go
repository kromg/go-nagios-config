package object

import ()

type Set map[interface{}]byte

func NewSet(elements ...string) Set {
	s := make(Set)
	for _, e := range elements {
		s.Add(e)
	}
	return s
}

func (s Set) Add(element interface{}) {
	s[element] = 1 // This value is actually not important
}

func (s Set) Remove(element interface{}) {
	delete(s, element)
}

func (s Set) Elements() []interface{} {
	elements := make([]interface{}, 0)
	for e, _ := range s {
		elements = append(elements, e)
	}
	return elements
}

func (s Set) AddAll(as Set) {
	for e, _ := range as {
		s.Add(e)
	}
}

func (s Set) Contains(element interface{}) bool {
	for e, _ := range s {
		if e == element {
			return true
		}
	}
	return false
}
