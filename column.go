package main

type Column struct {
	ID                        string
	status                    string
	amountOfFloors            int
	amountOfElevators         int
	amountOfElevatorPerColumn int
	servedFloorsList          []int
	isBasement                bool
	elevatorList              []Elevator
	callButtonList            []CallButton
}

var callButtonID int = 1
var elevatorID int = 1

func NewColumn(_id int, status string, _amountOfElevators int, _amountOfFloors int, _amountOfElevatorPerColumn int, _servedFloors []int, _isBasement bool) *Column {
	return &Column{_id, "online", _amountOfElevators, _amountOfFloors, 0, *createElevators(_amountOfFloors, _amountOfElevators), *createCallButtons(_amountOfFloors, _isBasement), _servedFloors}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator = findElevator(userPosition, direction, c)
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	if _isBasement {
		buttonFloor := -1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := *NewCallButton(callButtonID, "off", buttonFloor, "up")
			c.callButtonList = append(c.callButtonList, callButton)
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		for i := 0; i < _amountOfFloors; i++ {
			callButton := *NewCallButton(callButtonID, "off", buttonFloor, "down")
			c.callButtonList = append(c.callButtonList, callButton)
			buttonFloor++
			callButtonID++

		}
	}
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) *[]Elevator {
	var elevatorList []Elevator
	for i := 0; i < _amountOfElevators; i++ {
		elevator := *NewElevator(elevatorID, "idle", _amountOfFloors, 1)
		c.elevatorList = append(c.elevatorList, elevator)
		elevatorID
	}
	return &elevatorList
}

func (c *Column) requestElevator(userPosition int, direction string) *Elevator {
	elevator := findElevator(userPosition, direction)
	elevator.addNewRequest(userPosition)
	elevator.move()

	elevator.addNewRequest(1)
	elevator.move()
	return &elevator
}

func (c *Column) findElevator(_requestedFloor int, _direction string) *Elevator {
	var bestElevatorInformations = NewElevatorInformatons(6, 10000000, nil)
	if _requestedFloor == 1 {
		for elevator := range c.elevatorList {
			if 1 == elevator.currentFloor && elevator.status == "stopped" {
				bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
			} else if 1 == elevator.currentFloor && elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			} else {
				bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	} else {
		for elevator := range c.elevatorList {
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" {
				bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			} else {
				bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	}
	return bestElevatorInformations.bestElevator
}

func (c *Column) checkIfElevatorIsBetter(scoreToCheck int, NewElevator, BestElevatorInformations bestElevatorInformations, floor int) {
	if scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = scoreToCheck
		bestElevatorInformations.bestElevator = newElevator
		bestElevatorInformations.referenceGap = Math.Abs(newElevator.currentFloor - floor)
	} else if bestElevatorInformations.bestScore == scoreToCheck {
		gap = Math.Abs(newElevator.currentFloor - floor)
		if bestElevatorInformations.referenceGap > gap {
			bestElevatorInformations.bestElevator = newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return bestElevatorInformations
}
