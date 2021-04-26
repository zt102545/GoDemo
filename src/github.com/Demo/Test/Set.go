package main

import(

)

type Set struct{
	data map[interface{}]struct{}
}

func New(items ...interface{}) *Set {
	s := &Set{}
	s.data = make(map[interface{}]struct{})
	return s
}

func (s *Set) Add(items ...interface{}) error{
	for _, item :=range items{
		s.data[item]= struct{}{}
	}
	return nil
}

func (s *Set) Remove(item interface{}) error{
	delete(s.data, item)
	return nil
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.data[item]
	return ok
}

//交集
func (s *Set) Intersection(zs *Set) map[interface{}]struct{}{
	newSet:=make(map[interface{}]struct{})
	for key:= range zs.data {
		if s.Contains(key){
			newSet[key]= struct{}{}
		}
	}
	return newSet
}

func main(){

}
