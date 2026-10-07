package postgres

import "testing"

func TestBookingRepositoryFixture(t *testing.T) {
	var repository *BookingRepository
	if repository != nil {
		t.Fatal("expected nil repository in fixture")
	}
}
