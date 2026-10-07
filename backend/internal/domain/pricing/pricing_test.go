package pricing

import "testing"

func TestEstimate(t *testing.T) {
	engine := Engine{BaseCentsPerKm: 150, MinuteCents: 50}
	quote := engine.Estimate(10.1, 12.2)
	if quote.AmountCents != 2300 {
		t.Fatalf("unexpected amount: %d", quote.AmountCents)
	}
	if quote.Currency != "EUR" || quote.Version == "" {
		t.Fatalf("unexpected quote metadata: %+v", quote)
	}
}
