package api

// ResourceHandler can handle notifications for events that happen to a
// resource. The events are informational only, so you can't return an
// error.
//  * OnAdd is called when an object is added.
//  * OnUpdate is called when an object is modified. Note that oldObj is the
//      last known state of the object-- it is possible that several changes
//      were combined together, so you can't use this to see every single
//      change. OnUpdate is also called when a re-list happens, and it will
//      get called even if nothing changed. This is useful for periodically
//      evaluating or syncing something.
//  * OnDelete will get the final state of the item if it is known, otherwise
//      it will get an object of type DeletedFinalStateUnknown. This can
//      happen if the watch is closed and misses the delete event and we don't
//      notice the deletion until the subsequent re-list.
type ResourceHandler interface {
	OnAdd(obj interface{}) (interface{}, error)
	OnUpdate(oldObj, newObj interface{}) (interface{}, error)
	OnDelete(obj interface{}) error
}

// ResourceHandlerFuncs is an adaptor to let you easily specify as many or
// as few of the notification functions as you want while still implementing
// ResourceHandler.
type ResourceHandlerFuncs struct {
	AddFunc    func(obj interface{}) (interface{}, error)
	UpdateFunc func(oldObj, newObj interface{}) (interface{}, error)
	DeleteFunc func(obj interface{}) error
}

// OnAdd calls AddFunc if it's not nil.
func (r ResourceHandlerFuncs) OnAdd(obj interface{}) (interface{}, error) {
	if r.AddFunc != nil {
		return r.AddFunc(obj)
	}
	return nil, nil
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r ResourceHandlerFuncs) OnUpdate(oldObj, newObj interface{}) (interface{}, error) {
	if r.UpdateFunc != nil {
		return r.UpdateFunc(oldObj, newObj)
	}
	return nil, nil
}

// OnDelete calls DeleteFunc if it's not nil.
func (r ResourceHandlerFuncs) OnDelete(obj interface{}) error {
	if r.DeleteFunc != nil {
		return r.DeleteFunc(obj)
	}
	return nil
}