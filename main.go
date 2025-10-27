package main

import "fmt"

// Stack представляет структуру стека
type Stack struct {
	items []interface{}
}

// NewStack создает новый экземпляр стека
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push добавляет элемент на вершину стека
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size возвращает количество элементов в стеке
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear очищает стек
func (s *Stack) Clear() {
	s.items = make([]interface{}, 0)
}

// Peek возвращает элемент с вершины без удаления
func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.items[len(s.items)-1], true
}

func main() {
	// Создаем новый стек
	stack := NewStack()

	fmt.Println("Тестирование стека")

	// Проверяем пустой стек
	fmt.Printf("Стек пустой: %t\n", stack.IsEmpty())
	fmt.Printf("Размер стека: %d\n", stack.Size())

	// Добавляем элементы
	fmt.Println("\n Добавляем элементы")
	stack.Push("первый")
	stack.Push(42)
	stack.Push(3.14)
	stack.Push([]string{"a", "b", "c"})

	fmt.Printf("Стек после добавления: %v\n", stack)
	fmt.Printf("Размер стека: %d\n", stack.Size())
	fmt.Printf("Стек пустой: %t\n", stack.IsEmpty())

	// Просматриваем верхний элемент
	if top, ok := stack.Peek(); ok {
		fmt.Printf("Верхний элемент: %v\n", top)
	}

	// Извлекаем элементы
	fmt.Println("\n Извлекаем элементы")
	for !stack.IsEmpty() {
		if item, ok := stack.Pop(); ok {
			fmt.Printf("Извлечен: %v, тип: %T\n", item, item)
		}
		fmt.Printf("Текущий стек: %v\n", stack)
	}

	// Попытка извлечь из пустого стека
	fmt.Println("\n Попытка извлечь из пустого стека ")
	if item, ok := stack.Pop(); ok {
		fmt.Printf("Извлечен: %v\n", item)
	} else {
		fmt.Println("Стек пустой, нельзя извлечь элемент")
	}

	// Тестируем Clear
	fmt.Println("\n Тестируем Clear ")
	stack.Push("a1")
	stack.Push("a2")
	fmt.Printf("Стек до очистки: %v\n", stack)
	stack.Clear()
	fmt.Printf("Стек после очистки: %v\n", stack)
	fmt.Printf("Размер после очистки: %d\n", stack.Size())

	// Пример с разными типами данных
	fmt.Println("\n Работа с разными типами ")
	multiTypeStack := NewStack()
	multiTypeStack.Push("строка")
	multiTypeStack.Push(100)
	multiTypeStack.Push(true)
	multiTypeStack.Push(struct{ name string }{name: "тест"})

	for multiTypeStack.Size() > 0 {
		if item, ok := multiTypeStack.Pop(); ok {
			fmt.Printf("Элемент: %v, тип: %T\n", item, item)
		}
	}
}
