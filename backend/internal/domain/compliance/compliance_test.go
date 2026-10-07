package compliance

import "testing"

func TestCheck(t *testing.T) {
	if !Check(true, true, true).Eligible {
		t.Fatal("expected eligible")
	}
	result := Check(false, true, true)
	if result.Eligible || len(result.Reasons) != 1 || result.Reasons[0] != ProfessionalCardExpired {
		t.Fatalf("unexpected: %+v", result)
	}
}
