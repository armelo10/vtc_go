package httpapi

import ("encoding/json";"net/http"; pricingapp "github.com/armelo10/vtc_go/backend/internal/application/pricing")

type pricingHandler struct { service *pricingapp.Service }
type estimateRequest struct {
 PickupLat float64 `json:"pickup_lat"`
 PickupLng float64 `json:"pickup_lng"`
 DropoffLat float64 `json:"dropoff_lat"`
 DropoffLng float64 `json:"dropoff_lng"`
}
func newPricingHandler(s *pricingapp.Service)*pricingHandler{return &pricingHandler{service:s}}
func(h *pricingHandler)estimate(w http.ResponseWriter,r *http.Request){
 var req estimateRequest
 if err:=json.NewDecoder(r.Body).Decode(&req);err!=nil{writeJSON(w,400,map[string]string{"error":"invalid estimate"});return}
 if req.PickupLat < -90 || req.PickupLat > 90 || req.DropoffLat < -90 || req.DropoffLat > 90 || req.PickupLng < -180 || req.PickupLng > 180 || req.DropoffLng < -180 || req.DropoffLng > 180 {writeJSON(w,400,map[string]string{"error":"invalid coordinates"});return}
 q,km,min,err:=h.service.Estimate(r.Context(),req.PickupLat,req.PickupLng,req.DropoffLat,req.DropoffLng)
 if err!=nil{writeJSON(w,503,map[string]string{"error":"pricing unavailable"});return}
 writeJSON(w,200,map[string]any{"quote":q,"distance_km":km,"duration_min":min,"routing":"straight_line","warning":"provisional until a production routing provider is configured"})
}
