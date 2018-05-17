package main

// Player ...
type Player struct{}

func (p *Player) Read() ([]float32, error) {
	return nil, nil
}

func (p *Player) Write(v []float32) error {
	return nil
}
