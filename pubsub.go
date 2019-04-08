package newstore_pubsub

import "sync"

// pubsub stores a slice of callbacks registered for each event name
var pubsub map[string][]CallbackFunc

type CallbackFunc func(eventName string, payload Payload)

type Payload map[string]interface{}

func Subscribe(eventName string, callbackFunc CallbackFunc) {
	if pubsub == nil {
		pubsub = make(map[string][]CallbackFunc)
	}
	if _, ok := pubsub[eventName]; !ok {
		pubsub[eventName] = []CallbackFunc{}
	}
	pubsub[eventName] = append(pubsub[eventName], callbackFunc)
}

func Publish(eventName string, payload Payload) {
	callbacks, ok := pubsub[eventName]
	if !ok {
		panic("no such event")
	}

	var wg sync.WaitGroup
	wg.Add(len(callbacks))

	for _, callback := range callbacks {
		go func(callback CallbackFunc) {
			callback(eventName, payload)
			wg.Done()
		}(callback)
	}

	wg.Wait()
}
