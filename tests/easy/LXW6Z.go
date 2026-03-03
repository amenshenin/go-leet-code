package easy

type LXW6Z struct {
	IsComplete int
	Name       string
	Descr      string
}

func LXW6ZInit() *LXW6Z {
	return &LXW6Z{
		IsComplete: 0,
		Name:       "https://leetcode.com/problems/flatten-2d-vector/",
		Descr: `Разработайте итератор для разворачивания двумерного вектора. Он должен поддерживать операции next и hasNext.

Реализуйте класс Vector2D:

Vector2D(int[][] vec) инициализирует объект двумерным вектором vec.
next() возвращает следующий элемент из двумерного вектора и перемещает указатель на один шаг вперед. Вы можете предположить, что все вызовы next допустимы.
hasNext() возвращает true, если в векторе еще остались элементы, и false в противном случае.`,
	}
}

func (f *LXW6Z) Run() error {

	return nil
}
