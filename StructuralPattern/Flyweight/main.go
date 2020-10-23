package main

import "fmt"

const (
	TerroristDressType = "tDress"
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

type dress interface {
	getColor() string
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if d.dressMap[dressType] != nil {
		d.dressMap[dressType] = newCounterTerroristDress()

		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}

func newTerroristDress() *terroristDress {
	return &terroristDress{
		color: "red",
	}
}

type counterTerroristDress struct {
	color string
}

func (t *counterTerroristDress) getColor() string {
	return t.color
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{
		color: "green",
	}
}

type player struct {
	dress dress
	playerType string
	lat int
	long int
}

func newPlayer(playerType, dressType string) *player{
	dress, _ := GetDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress: dress,
	}
}

func (p *player) newLocatiom(lat, long int)  {
	p.lat = lat
	p.long = long
}

func main()  {
	game := newGame()

	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	game.addCounterTerrorist(CounterTerrroristDressType)
	dressFactoryInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}