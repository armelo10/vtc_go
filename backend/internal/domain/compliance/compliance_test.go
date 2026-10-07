package compliance
import "testing"
func TestCheck(t *testing.T){if !Check(true,true,true).Eligible{t.Fatal("expected eligible")};r:=Check(false,true,true);if r.Eligible||len(r.Reasons)!=1||r.Reasons[0]!=ProfessionalCardExpired{t.Fatalf("unexpected: %+v",r)}}
