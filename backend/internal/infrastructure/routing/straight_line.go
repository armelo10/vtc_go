package routing

import ("context";"math")
type StraightLineProvider struct { AverageSpeedKmh float64 }
func NewStraightLineProvider(speed float64)*StraightLineProvider{if speed<=0{speed=30};return &StraightLineProvider{AverageSpeedKmh:speed}}
func(p *StraightLineProvider)Estimate(_ context.Context,a,b,c,d float64)(float64,float64,error){
 const r=6371.0088
 x,y:=a*math.Pi/180,c*math.Pi/180
 dy:=y-x; dx:=(d-b)*math.Pi/180
 q:=math.Sin(dy/2)*math.Sin(dy/2)+math.Cos(x)*math.Cos(y)*math.Sin(dx/2)*math.Sin(dx/2)
 km:=r*2*math.Atan2(math.Sqrt(q),math.Sqrt(1-q))
 return km,km/p.AverageSpeedKmh*60,nil
}
