package compliance

type Reason string

const (
	ProfessionalCardExpired Reason = "PROFESSIONAL_CARD_EXPIRED"
	VehicleDocumentExpired  Reason = "VEHICLE_DOCUMENT_EXPIRED"
	NotRegistered           Reason = "REVTC_REGISTRATION_MISSING"
)

type Result struct {
	Eligible bool
	Reasons  []Reason
}

func Check(card, vehicle, revtc bool) Result {
	var reasons []Reason
	if !card {
		reasons = append(reasons, ProfessionalCardExpired)
	}
	if !vehicle {
		reasons = append(reasons, VehicleDocumentExpired)
	}
	if !revtc {
		reasons = append(reasons, NotRegistered)
	}
	return Result{Eligible: len(reasons) == 0, Reasons: reasons}
}
