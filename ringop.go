package goring

// TODO: Add Not Sync -> then Sync
// interface of ring
type RingOp interface {
	Cap() int
	Len() int
	Push(t interface{})
	PushSlice(s []interface{})
	Pop() interface{}
	At(i int) (interface{}, bool)
	Range(func(i int, t interface{}) bool)
	SetCap(c int) error
}
