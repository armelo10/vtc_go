package pricingapp

import (
 "context"
 "testing"
 domainpricing "github.com/armelo10/vtc_go/backend/internal/domain/pricing"
)

type fakeRouting struct{}
func (fakeRouting) Estimate(context.Context,float64,float64,float64,float64)(float64,float64,error){return 10.1,12.2,nil}
func TestServiceEstimate(t *testing.T){
 s:=NewService(fakeRouting{},domainpricing.Engine{BaseCentsPerKm:150,MinuteCents:50})
 q,km,min,err:=s.Estimate(context.Background(),48.85,2.35,48.86,2.36)
 if err!=nil{t.Fatal(err)}
 if km!=10.1||min!=12.2||q.AmountCents!=2300{t.Fatalf("unexpected estimate: %+v distance=%v duration=%v",q,km,min)}
}
