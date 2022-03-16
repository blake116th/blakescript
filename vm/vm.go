package vm

import "math/rand"

type HeapId int
type Heap map[HeapId] int

//Runtime environment
type Env struct {
	Heap Heap
}

func (env *Env) MakeHeapId() HeapId {
	retVal := HeapId(rand.Int())

	//naively deal with Id collision
	if _, contains := env.Heap[retVal]; contains {
		return env.MakeHeapId()
	}
	return retVal;
}