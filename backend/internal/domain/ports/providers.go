package ports
import "context"
type PaymentProvider interface{CreatePayment(context.Context,int64,string,string)(string,error);RefundPayment(context.Context,string,int64)error;HandleWebhook(context.Context,[]byte,string)error}
type RoutingProvider interface{Geocode(context.Context,string)(float64,float64,error);Estimate(context.Context,float64,float64,float64,float64)(float64,float64,error)}
type NotificationProvider interface{SendSMS(context.Context,string,string)error;SendEmail(context.Context,string,string,string)error;SendPush(context.Context,string,string,string)error}
