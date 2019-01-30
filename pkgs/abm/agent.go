package abm

// Agent represents the method signature required in the simulation
type Agent interface {
	Run() interface{}
}
