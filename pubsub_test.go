package newstore_pubsub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSimpleSubscribe(t *testing.T) {
	event := "shoe_added"

	Subscribe(event, func(eventName string, payload Payload) {})

	if _, ok := pubsub[event]; !ok {
		t.Fatalf("no %s event was added", event)
	}
}

func TestSimplePublish(t *testing.T) {
	event := "newShoe"

	var state bool

	Subscribe(event, func(eventName string, payload Payload) {
		state = true
	})

	Publish(event, Payload{})

	if state == false {
		t.Fatal("the registered callback didn't run")
	}
}

func TestDifferentCallbacks(t *testing.T) {
	event := "different callbacks"

	var counter int = 0

	cb1 := func(eventName string, payload Payload) {
		counter = counter + 2
		fmt.Println("ran cb 1")
	}
	Subscribe(event, cb1)

	cb2 := func(eventName string, payload Payload) {
		counter = counter + 2
		fmt.Println("ran cb 2")
	}
	Subscribe(event, cb2)

	cb3 := func(eventName string, payload Payload) {
		counter = counter + 2
		fmt.Println("ran cb 3")
	}

	callbacks := []CallbackFunc{cb1, cb2, cb3}
	for _, cb := range callbacks {
		Subscribe(event, cb)
	}

	var present []int
	for i, cb := range callbacks {
		valCb := reflect.ValueOf(cb)
		for _, x := range pubsub[event] {
			valX := reflect.ValueOf(x)
			if valCb.Pointer() == valX.Pointer() {
				present = append(present, i)
				break
			}
		}
	}
	if len(present) > 0 {
		t.Fatalf("callback %v are present", present)
	}

}

func TestMultipleCallbacks(t *testing.T) {
	event := "multiple callbacks"

	var counter int = 0

	Subscribe(event, func(eventName string, payload Payload) {
		counter = counter + 2
		fmt.Println("ran cb 1")
	})

	Subscribe(event, func(eventName string, payload Payload) {
		counter = counter + 3
		fmt.Println("ran cb 2")
	})

	Subscribe(event, func(eventName string, payload Payload) {
		counter = counter + 4
		fmt.Println("ran cb 3")
	})

	Publish(event, Payload{})

	if counter < 3 {
		t.Fatal("problem with multiple callbacks")
	}
}
