package main

type DynamicArray struct {
	items []string
	count int
}

func NewDynamicArray(capacity int) *DynamicArray {
	array := new(DynamicArray)
	array.items = make([]string, capacity)
	return array
}

func (array *DynamicArray) Add(item string) {
	if array.Size() == array.Capacity() {
		array.growArray()
	}
	array.items[array.count] = item
	array.count++
}

func (array DynamicArray) Capacity() int {
	if cap(array.items) >= 1 {
		return cap(array.items)
	}
	return 0
}

func (array DynamicArray) Size() int {
	if array.count >= 1 {
		return array.count
	}
	return 0
}

func (array DynamicArray) Get(index int) string {
	result := array.items[index]
	if result != "" {
		return result
	}
	return ""
}

func (array *DynamicArray) growArray() {
	if cap(array.items) == 0 {
		newSlice := make([]string, 1)
		copy(newSlice, array.items)
		array.items = newSlice
	} else {
		newSlice := make([]string, cap(array.items)*2)
		copy(newSlice, array.items)
		array.items = newSlice
	}
}

func (array *DynamicArray) Remove(index int) {
	copy(array.items[index:], array.items[index+1:])
	array.count--
}

func (array *DynamicArray) Insert(index int, item string) {
	array.items = append(array.items, "")
	copy(array.items[index+1:], array.items[index:])
	array.items[index] = item
	array.count++
}

func (array DynamicArray) Contains(item string) bool {
	for _, val := range array.items {
		if val == item {
			return true
		}
	}
	return false
}

func (array DynamicArray) IndexOf(item string) int {
	for index, val := range array.items {
		if val == item {
			return index
		}
	}
	return 0
}

func (array *DynamicArray) Clear() {
	array.items = array.items[:0]
	array.count = 0
}

func main() {


}

//array.items = append(array.items[:index], array.items[index+1:]...)
//array.count--
