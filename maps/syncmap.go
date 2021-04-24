package maps

import (
	"fmt"
	"sync"
	"time"
)

func SyncMAP() {
	ms := sync.Map{}
	ms.Store("item 1", "value of item 1")
	ms.Store("item 2", "value of item 2")
	ms.Store("item 3", "value of item 3")
	ms.Store("item 4", "value of item 4")
	ms.Store("item 5", "value of item 5")
	ms.Store("item 6", "value of item 6")
	ms.Store("item 7", "value of item 7")
	ms.Store("item 8", "value of item 8")
	ms.Store("item 9", "value of item 9")
	ms.Store("item 10", "value of item 10")

	value, hasValue := ms.Load("item 2")
	if hasValue {
		fmt.Println("the item was found", value.(string))
		<-time.After(5 * time.Millisecond)
	}

	ms.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key %v -> value: %v", key, value)
		return true
	})

	ms.Delete("item 10")
	fmt.Println("\n item 10 has been removed:")

	ms.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key %v -> value: %v", key, value)
		<-time.After(5 * time.Millisecond)
		return true
	})

	newItem, ok := ms.LoadOrStore("item 10", "value of item 10")

	if ok {
		fmt.Println("item 10 was already on the sync map")
	}

	fmt.Println("new item has been added to the sync map", newItem)

	ms.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key %v -> value: %v", key, value)
		<-time.After(5 * time.Millisecond)
		return true
	})
}
