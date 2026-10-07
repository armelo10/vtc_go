package pricing
import "math"
type Quote struct{AmountCents int64;Currency,Version string}
type Engine struct{BaseCentsPerKm,MinuteCents int64}
func(e Engine)Estimate(distanceKm,durationMin float64)Quote{return Quote{AmountCents:e.BaseCentsPerKm*int64(math.Ceil(distanceKm))+e.MinuteCents*int64(math.Ceil(durationMin)),Currency:"EUR",Version:"mvp-1"}}
