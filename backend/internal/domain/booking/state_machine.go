package booking

import "fmt"

func CanTransition(from, to Status) bool {
	if from == to {
		return true
	}
	switch from {
	case Draft:
		return to == Requested || to == Cancelled
	case Requested:
		return to == SearchingDriver || to == Cancelled || to == Expired || to == Failed
	case SearchingDriver:
		return to == DriverAssigned || to == Cancelled || to == Expired || to == Failed
	case DriverAssigned:
		return to == DriverArriving || to == Cancelled || to == Failed
	case DriverArriving:
		return to == DriverAtPickup || to == Cancelled || to == NoShow || to == Failed
	case DriverAtPickup:
		return to == PassengerOnboard || to == NoShow || to == Failed
	case PassengerOnboard:
		return to == InProgress || to == Failed
	case InProgress:
		return to == Completed || to == Failed
	default:
		return false
	}
}

func Transition(from, to Status) error {
	if !CanTransition(from, to) {
		return fmt.Errorf("invalid booking transition: %s -> %s", from, to)
	}
	return nil
}
