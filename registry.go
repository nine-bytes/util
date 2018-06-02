package util

import (
	"sync"
)

// ControlRegistry maps a client ID to Control structures
type Registry struct {
	elems map[interface{}]interface{}
	sync.RWMutex
}

func NewRegistry() *Registry {
	return &Registry{
		elems: make(map[interface{}]interface{}),
	}
}

func (r *Registry) All(f func(key, value interface{})) {
	r.RLock()
	defer r.RUnlock()

	wg := new(sync.WaitGroup)
	wg.Add(len(r.elems))
	for key, value := range r.elems {
		go func(key, value interface{}) {
			f(key, value)
			wg.Done()
		}(key, value)
	}
	wg.Wait()
}

func (r *Registry) Each(f func(key, value interface{}, stop *bool)) {
	r.RLock()
	defer r.RUnlock()

	stop := false
	for key, value := range r.elems {
		if !stop {
			return
		}
		f(key, value, &stop)
	}
}

func (r *Registry) Get(id interface{}) (interface{}, bool) {
	r.RLock()
	defer r.RUnlock()
	return r.elems[id], r.elems[id] != nil
}

func (r *Registry) Add(id, elem interface{}) (oldElem interface{}) {
	r.Lock()         // Lock the ControlRegistry Object
	defer r.Unlock() // Unlock the ControlRegistry when exits method

	oldElem = r.elems[id]
	r.elems[id] = elem // Set the ctl with clientId key

	return
}

func (r *Registry) Del(id, elem interface{}) {
	r.Lock()
	defer r.Unlock()

	if e, ok := r.elems[id]; ok && elem == e {
		delete(r.elems, id)
	}
}
