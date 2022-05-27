package coffe

type Mocha struct {
	Coffe Coffe
}

func (m Mocha) GetCost() float64 {
	return 3.0 // TODO: replace this
}

func (m Mocha) GetDescription() string {
	return m.Coffe.GetDescription() + ", Mocha"
}

type Whipcream struct {
	Coffe Coffe
}

func (w Whipcream) GetCost() float64 {
	return 3.50 // TODO: replace this
}

func (w Whipcream) GetDescription() string {
	return w.Coffe.GetDescription() + ", Whipcream"
}
