package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

// создание списока и возвращение указателя
func NewList() *List {
	newList := &List{}
	return newList
}

// возращение длины списка из поля len структуры List
func (l *List) Len() (len int64) {
	return l.len
}

// добавление элемента в список и возвращение его индекса
func (l *List) Add(value int64) (id int64) {
	newNode := &node{
		value: value,
		next:  nil,
	} // новый узел с заданным значением value добавляем в конец списка

	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		current := l.firstNode
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	} // Пустой список? Устанавливаем новый узел как первый узел.
	/*В противном случае, мы перебираем список, находим последний узел,
	добавляем новый узел в конец списка*/

	l.len++
	return l.len - 1 /* Возвращение индекса добавленного элемента,
	который равен текущей длине списка - 1 */
}

// удаление элемента списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	} // Индекс в допустимых приделах?

	if id == 0 {
		// Удаляемый первый элемент списка - переназначаем первый узел
		l.firstNode = l.firstNode.next
	} else {
		/* В противном случае, находим узел перед удаляемым узлом
		   и переназначаем указатель на следующий узел*/
		current := l.firstNode
		for i := int64(0); i < id-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	l.len-- /* удаляем элемент из списка, обнавляем указатели
	узлов таким образом, чтобы исключить удаляемый элемент.*/
}

// удаляение элемента списка по значению
func (l *List) RemoveByValue(value int64) {
	if l.len == 0 {
		return
	} // в том случае если список пуст

	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
		return
	} // удаляемый элемент в начале списка

	current := l.firstNode
	for current.next != nil && current.next.value != value {
		current = current.next
	} // поиск элемента для удаления

	if current.next == nil {
		return
	} // элемент не найден - выходим

	// удаляем элемент, переназначаем указатель на следующий узел
	current.next = current.next.next
	l.len--
}

// удаление всех элементов списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.len == 0 {
		return
	} // в том случае если список пуст

	for l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	} // удаляем нужные элементы в начале списка

	current := l.firstNode
	for current != nil && current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.len--
		} else {
			current = current.next
		}
	} // поиск и удаление остальных элементов
}

// возвращение значения элемента по индексу
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 || id >= l.len {
		return 0, false
	} // индекс недопустим - возвращаем 0 и false

	current := l.firstNode
	for i := int64(0); i < id; i++ {
		current = current.next
	}

	return current.value, true
	// элемент найден - возвращаем значение и true
}

// возвращение индекса первого найденного элемента по значению
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	current := l.firstNode
	id = int64(0)

	for current != nil {
		if current.value == value {
			return id, true
		} //  элемент найден - возвращаем значение и true
		current = current.next
		id++
	}

	return 0, false
	// элемент с указанным значением не найден - возвращаем 0 и false
}

// возвращение индексов всех найденных элементов по значению
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	current := l.firstNode

	for id := int64(0); current != nil; id++ {
		if current.value == value {
			ids = append(ids, id)
		}
		current = current.next
	}

	if len(ids) > 0 {
		return ids, true
	} // элементы с указанным значением найдены - возвращаем их и true

	return nil, false
	// элементы с указанным значением не найдены - возвращаем nil и false,
}

// возвращение всех элементов списка
func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false // список пуст - возвращаем nil и false
	}

	values = make([]int64, l.len)
	current := l.firstNode
	index := 0

	for current != nil {
		values[index] = current.value
		current = current.next
		index++
	}

	return values, true
	// возвращаем найденные элементы и true
}

// очистка списка
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}

// вывод списка в консоль
func (l *List) Print() {
	current := l.firstNode

	fmt.Print("Список: ")

	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}

	fmt.Println()
}
