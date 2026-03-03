package hard

import (
	"fmt"
	"math/rand"
)

type Lf9p7 struct {
	IsComplete int
	Name       string
	Descr      string
}

func Lf9p7Init() *Lf9p7 {
	return &Lf9p7{
		IsComplete: 1,
		Name:       "https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/",
		Descr: `RandomizedCollection — это структура данных, содержащая набор чисел, возможно с дубликатами (т.е. мультимножество). Она должна поддерживать вставку и удаление определенных элементов, а также предоставление случайного элемента.

Реализуйте класс RandomizedCollection:

RandomizedCollection(): Инициализирует пустой объект RandomizedCollection.
bool insert(int val): Вставляет элемент val в мультимножество, даже если элемент уже присутствует. Возвращает true, если элемента не было, и false в противном случае.
bool remove(int val): Удаляет элемент val из мультимножества, если он присутствует. Возвращает true, если элемент присутствовал, и false в противном случае. Если у val несколько вхождений в мультимножестве, удаляется только одно из них.
int getRandom(): Возвращает случайный элемент из текущего мультимножества. Вероятность возврата каждого элемента пропорциональна числу вхождений этого элемента в мультимножество.
Реализуйте функции класса так, чтобы каждая функция работала в среднем за O(1) времени.
`,
	}
}

type RandomizedCollection struct {
	dict map[int]map[int]struct{}
	list []int
}

type IRandomizedCollection interface {
	Incert(val int) bool
	Remove(val int) bool
	GetRandom() int
}

func Constructor() *RandomizedCollection {
	return &RandomizedCollection{
		dict: make(map[int]map[int]struct{}),
		list: []int{},
	}
}

func (rc *RandomizedCollection) Incert(val int) bool {
	_, isExists := rc.dict[val]
	if !isExists {
		rc.dict[val] = make(map[int]struct{})
	}
	rc.dict[val][len(rc.list)] = struct{}{}
	rc.list = append(rc.list, val)
	return isExists
}

func (rc *RandomizedCollection) Remove(val int) bool {
	_, isExists := rc.dict[val]
	if !isExists {
		return false
	}
	index := 0
	for i, _ := range rc.dict[val] {
		index = i
		break
	}
	delete(rc.dict[val], index)
	lastElement := rc.list[len(rc.list)-1]
	rc.list[index] = lastElement
	rc.dict[lastElement][index] = struct{}{}
	delete(rc.dict[lastElement], len(rc.list)-1)
	rc.list = rc.list[:len(rc.list)-1]
	if len(rc.dict[val]) == 0 {
		delete(rc.dict, val)
	}
	return true
}

func (rc *RandomizedCollection) GetRandom() int {
	return rc.list[rand.Intn(len(rc.list))]
}

func (f *Lf9p7) Run() error {
	rc := Constructor()
	fmt.Println(rc.list)

	rc.Incert(1)
	fmt.Println(rc.list)

	rc.Incert(1)
	fmt.Println(rc.list)

	rc.Incert(1)
	fmt.Println(rc.list)

	fmt.Println(rc.GetRandom())

	rc.Remove(1)
	fmt.Println(rc.list)

	fmt.Println(rc.GetRandom())

	return nil
}
