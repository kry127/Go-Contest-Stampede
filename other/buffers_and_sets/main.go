package main

import "fmt"

// set structure
// https://gist.github.com/bgadrian/cb8b9344d9c66571ef331a14eb7a2e80
type Set struct {
	list map[int]struct{} //empty structs occupy 0 memory
}

func (s *Set) Has(v int) bool {
	_, ok := s.list[v]
	return ok
}

func (s *Set) Add(v int) {
	s.list[v] = struct{}{}
}

func (s *Set) Remove(v int) {
	delete(s.list, v)
}

func (s *Set) Clear() {
	s.list = make(map[int]struct{})
}

func (s *Set) Size() int {
	return len(s.list)
}

func (s *Set) List() []int {
	// https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
	keys := make([]int, len(s.list))
	i := 0
	for k := range s.list {
		keys[i] = k
		i++
	}
	return keys
}

func NewSet() *Set {
	s := &Set{}
	s.list = make(map[int]struct{})
	return s
}

//optional functionalities

//AddMulti Add multiple values in the set
func (s *Set) AddMulti(list ...int) {
	for _, v := range list {
		s.Add(v)
	}
}

type FilterFunc func(v int) bool

// Filter returns a subset, that contains only the values that satisfies the given predicate P
func (s *Set) Filter(P FilterFunc) *Set {
	res := NewSet()
	for v := range s.list {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set) Union(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		res.Add(v)
	}

	for v := range s2.list {
		res.Add(v)
	}
	return res
}

func (s *Set) Intersect(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from s, that doesn't exists in s2 (param)
func (s *Set) Difference(s2 *Set) *Set {
	res := NewSet()
	for v := range s.list {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}
// end of set structure

var messageBuffer = make(chan int, 3) // Создаём очередь размером в 3 инта
var finishedProducing = make(chan bool)
var finishedConsuming = make(chan bool)

func producer(amount int) {
	for i := 0; i < amount; i++ {
		suppliedValue := i
		fmt.Printf("Produced: %v\n", suppliedValue)
		messageBuffer <- suppliedValue
	}

	finishedProducing <- true
}

func consumer() {
	for {
		select {
		case <- finishedProducing: // извлекаем из очереди "finishedProducing" значение
			finishedConsuming <- true
			return
		case result := <- messageBuffer:
			fmt.Printf("Consumed: %v\n", result)
		}
	}
}

func main() {
	// section 1: cool coroutines
	fmt.Printf("Hello, Go world!")
	go producer(400)
	go consumer()
	<- finishedConsuming
	fmt.Printf("All corutines finished!\n")

	// work with sets
	// https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
	mySet := NewSet()
	fmt.Println("has 5", mySet.Has(5)) //false
	mySet.Add(5)
	fmt.Println("has 5", mySet.Has(5)) //true
	mySet.Remove(5)
	fmt.Println("has 5", mySet.Has(5)) //false

	mySet.AddMulti(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 3, 2, 1)

	p := func(v int) bool { return v >= 5 }
	subSet := mySet.Filter(p)
	fmt.Println("all >= 5", subSet.List())


	mySet1 := NewSet()
	mySet1.AddMulti(1, 2, 3, 4, 5, 6)

	mySet2 := NewSet()
	mySet2.AddMulti(18, 2, 3, 15, 16, 17)

	fmt.Println("mySet1:", mySet1.List())
	fmt.Println("mySet2:", mySet2.List())

	fmt.Println("mySet1 U mySet2:", mySet1.Union(mySet2).List())
	fmt.Println("mySet1 n mySet2:", mySet1.Intersect(mySet2).List())
	fmt.Println("mySet1 \\ mySet2:", mySet1.Difference(mySet2).List())
	fmt.Println("mySet2 \\ mySet1:", mySet2.Difference(mySet1).List())

}
