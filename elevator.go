package main

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	door                  Door
	floorRequestsList     []int
	direction             string
	overweight            bool
	completedRequestsList []int
}

func NewElevator(_id int, _status string, _amountOfFloors int, _currentFloor string) *Elevator {
	elevator := &Elevator(_id, _status, _amountOfFloors, 1, _currentFloor, "up", *NewDoor(1, "opened"), []int{}, []int{})
	return elevator
}

func (e *Elevator) move() {
	i := 0
	for i < len(e.floorRequestsList) != 0 {
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if currentFloor > destination {
			e.direction = "down"
			e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		e.operateDoors()
		e.completedRequestsList = append(e.completedRequestsList, floorRequestsList[0])
		e.floorRequestsList = e.floorRequestList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if direction == "up" {
		floorRequestsList.Sort()
	} else {
		floorRequestsList.Sort()
	}
}

func (e *Elevator) operateDoors() {
	e.door.status = "opened"
}

func (e *Elevator) addNewRequest(_requestedFloor int) {
	if !floorRequestsList.Contains(_requestedFloor) {
		floorRequestsList.Add(_requestedFloor)
	}

	if currentFloor < _requestedFloor {
		e.direction = "up"
	}
	if currentFloor > _requestedFloor {
		e.direction = "down"
	}
}
