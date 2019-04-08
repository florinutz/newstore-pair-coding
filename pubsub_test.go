package newstore_pubsub

import (
	"fmt"
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

	cb1 := func(eventName string, payload Payload) {
		fmt.Println("ran cb 1")
	}

	cb2 := func(eventName string, payload Payload) {
		fmt.Println("ran cb 2")
	}

	cb3 := func(eventName string, payload Payload) {
		fmt.Println("ran cb 3")
	}

	callbacks := []CallbackFunc{cb1, cb2, cb3}
	for _, cb := range callbacks {
		Subscribe(event, cb)
	}

	Publish("different callbacks", Payload{})
}

func TestMultipleCallbacks(t *testing.T) {
	event := "multiple callbacks"

	var counter int = 0

	Subscribe(event, func(eventName string, payload Payload) {
		counter = counter + 2
	})

	Subscribe(event, func(eventName string, payload Payload) {
		counter = counter + 3
	})

	Subscribe(event, func(eventName string, payload Payload) {
		counter = counter + 4
	})

	Publish(event, Payload{})

	if counter < 3 {
		t.Fatal("problem with multiple callbacks")
	}
}
