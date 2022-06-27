package main

import "math"

type Battery struct {
	ID                        int
	status                    string
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnList                []Column
	floorRequstsButtonList    []FloorRequestButton
}

var columnID int = 1
var floorRequestButtonID int = 1

func NewBattery(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) *Battery {
	battery = &Battery{_id, "online", _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn, []column{}, []FloorRequestButton{}}
	battery.createColumns(_amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn)
	battery.createFloorRequestButton(_amountOfFloors)
	if _amountOfBasements > 0 {
		battery.createBasementFloorRequestButtons(_amountOfBasements)
		battery.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	return battery
}

func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	var servedFloors []int
	floor := -1
	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(floor)
		floor--
	}
	coulmn := *NewColumn(columnID, "online", _amountOfBasements, _amountOfEle_amountOfElevatorPerColumn)
	battery.columnList = append(battery.columnList, coulmn)
	columnID++
}

func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorsPerColumn int) {
	amountOfElevatorPerColumn := int(math.Ceil(float64(_amountOfFloors) / float64(_amountOfColumns)))
	floor := -1
	for i := 0; i < _amountOfColumns; i++ {
		servedFloors := []int{}
		for k := 0; k < amountOfFloorsPerColumn; k++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := *NewColumn(columnID, _amountOfFloors, _amountOfElevatorsPerColumn, servedFloors, false)
		b.columnList = append(b.columnList, column)
		coulmnID++
	}
}

func (b *Battery) createFloorRequestButton(_amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < _amountOfFloors; i++ {
		FloorRequestButton := *NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "up")
		b.floorRequstsButtonList = append(b.floorRequstsButtonList, floorRequestButton)
		buttonFloor++
		floorRequestButtonID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	buttonFloor := 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButton := *NewFloorRequestButton(floorRequestButtonID, "off", buttonFloor, "down")
		b.floorRequstsButtonList = append(b.floorRequstsButtonList, floorRequestButton)
		buttonFloor--
		floorRequestButtonID++
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for column := range b.columnList {
		if column.servedFloorsList.Contains(_requestedFloor) {
			return column
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator = column.findElevator(1, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()

	return &column, &elevator
}
