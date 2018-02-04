package structures

import "testing"

type itemString struct {
	Item
	value string
}

var listElements = []string{"alpha", "bravo", "charlie"}

func listFromStrings(lstr []string) *LinkedList {
	ll := NewLinkedList()
	for i := range lstr {
		item := &Item{lstr[i], nil, nil}
		ll.Append(item)
	}
	return ll
}

func TestList_Create(t *testing.T) {
	ll := listFromStrings(listElements)

	i := 0
	for item := ll.head; item != nil; item = item.Next {
		if item.Value.(string) != listElements[i] {
			t.Fatalf("got value '%v', expected '%s'", item.Value, listElements[i])
		}
		i++
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	ll := NewLinkedList()

	for i := range listElements {
		item := &Item{listElements[i], nil, nil}
		ll.Prepend(item)
	}

	item := ll.head
	for i := len(listElements) - 1; i >= 0; i-- {
		//fmt.Printf("i: %d, ", i)
		//fmt.Printf("item: %v\n", item.Value)
		val, ok := item.Value.(string)
		if !ok || val != listElements[i] {
			t.Fatalf("got value '%v', expected '%s'", item.Value, listElements[i])
		}
		item = item.Next
	}
}
