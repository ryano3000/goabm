package shopper

// Agent struct represents the shopper agent used in the simualtion
type Agent struct {
	Status string
}

// Run captures the state change of the agent after every tick
func (a *Agent) Run() interface{} {
	switch a.Status {
	case "store":
		if buy() == 1 {
			a.Status = "buy"
			return a
		}
		a.Status = "dropped"
		return a
	case "buy":
		if buy() == 1 {
			a.Status = "bought"
			return a
		}
		a.Status = "dropped"
		return a
	default:
		if buy() == 1 {
			a.Status = "store"
			return a
		}
		a.Status = "dropped"
		return a
	}
}
