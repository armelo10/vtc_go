package compliance

type Reason string

const (
	ProfessionalCardExpired Reason = "PROFESSIONAL_CARD_EXPIRED"
	VehicleDocumentExpired Reason = "VEHICLE_DOCUMENT_EXPIRED"
	NotRegistered Reason = "REVTC_REGISTRATION_MISSING"
)

type Result struct {
	Eligible bool
	Reasons  []Reason
}

func Check(proCardValid, vehicleDocsValid, revtcValid bool) Result {
	var reasons []Reason
	if !proCardValid { reasons = append(reasons, ProfessionalCardExpired) }
	if !vehicleDocsValid { reasons = append(reasons, VehicleDocumentExpired) }
	if !revtcValid { reasons = append(reasons, NotRegistered) }
	return Result{Eligible: len(reasons) == 0, Reasons: reasons}
}
