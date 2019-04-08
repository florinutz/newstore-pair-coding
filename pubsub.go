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
		return
	}

	var wg sync.WaitGroup

	for _, callback := range callbacks {
		wg.Add(1)
		go func() {
			callback(eventName, payload)
			wg.Done()
		}()
	}

	wg.Wait()
}
