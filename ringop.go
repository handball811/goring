package goring

// TODO: Add Sync Ring
// interface of ring
type RingOp interface {
	// Capacity of this ring
	Cap() int
	//Current Length of this ring
	Len() int

	Push(t interface{})
	PushSlice(s []interface{})
	Pop() (interface{}, bool)
	PopSlice(s []interface{}) int

	At(i int) (interface{}, bool)
	Range(c func(i int, t interface{}) bool)
	SetCap(c int) bool
	Clean()

	CloneTo(op RingOp)
}
