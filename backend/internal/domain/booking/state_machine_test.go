package booking

import "testing"

func TestValidTransition(t *testing.T) {
	for _, tc := range []struct {
		from Status
		to   Status
	}{
		{Requested, SearchingDriver},
		{SearchingDriver, DriverAssigned},
		{DriverAssigned, DriverArriving},
		{DriverArriving, DriverAtPickup},
		{DriverAtPickup, PassengerOnboard},
		{PassengerOnboard, InProgress},
		{InProgress, Completed},
	} {
		if err := Transition(tc.from, tc.to); err != nil {
			t.Fatalf("expected %s -> %s to be valid: %v", tc.from, tc.to, err)
		}
	}
}

func TestInvalidTerminalTransition(t *testing.T) {
	if err := Transition(Completed, InProgress); err == nil {
		t.Fatal("expected terminal transition to fail")
	}
}

func TestCancellationWindow(t *testing.T) {
	if err := Transition(SearchingDriver, Cancelled); err != nil {
		t.Fatal("expected cancellation to be allowed")
	}
	if err := Transition(PassengerOnboard, Cancelled); err == nil {
		t.Fatal("expected cancellation to be rejected after boarding")
	}
}
