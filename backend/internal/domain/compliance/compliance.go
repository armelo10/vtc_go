package compliance
type Reason string
const(ProfessionalCardExpired Reason="PROFESSIONAL_CARD_EXPIRED";VehicleDocumentExpired Reason="VEHICLE_DOCUMENT_EXPIRED";NotRegistered Reason="REVTC_REGISTRATION_MISSING")
type Result struct{Eligible bool;Reasons []Reason}
func Check(card,vehicle,revtc bool)Result{var r []Reason;if !card{r=append(r,ProfessionalCardExpired)};if !vehicle{r=append(r,VehicleDocumentExpired)};if !revtc{r=append(r,NotRegistered)};return Result{Eligible:len(r)==0,Reasons:r}}
