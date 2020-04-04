package Facade

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Fprintf(outputWriter, " CarModel - SetModel\n")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}

func (c *CarEngine) SetEngine() {
	fmt.Fprintf(outputWriter, " CarEngine - SetEngine\n")
}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetBody() {
	fmt.Fprintf(outputWriter, " CarBody - SetBody\n")
}

type CarAccessories struct {
}

func NewCarAccessories() *CarAccessories {
	return &CarAccessories{}
}

func (c *CarAccessories) SetAccessories() {
	fmt.Fprintf(outputWriter, " CarAccessories - SetAccessories\n")
}

type CarFacade struct {
	accessories *CarAccessories
	body        *CarBody
	engine      *CarEngine
	model       *CarModel
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		accessories: NewCarAccessories(),
		body:        NewCarBody(),
		engine:      NewCarEngine(),
		model:       NewCarModel(),
	}
}

func (c *CarFacade) CreateCompleteCar() {
	fmt.Fprintf(outputWriter, "******** Creating a Car **********\n")
	c.model.SetModel()
	c.engine.SetEngine()
	c.body.SetBody()
	c.accessories.SetAccessories()
	fmt.Fprintf(outputWriter, "******** Car creation is completed. **********\n")
}
