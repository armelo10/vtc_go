package booking
import("errors";"time")
type ServiceType string
const(ServiceVTC ServiceType="VTC";ServiceTaxi ServiceType="TAXI")
type Status string
const(Draft Status="DRAFT";Requested Status="REQUESTED";SearchingDriver Status="SEARCHING_DRIVER";DriverAssigned Status="DRIVER_ASSIGNED";DriverArriving Status="DRIVER_ARRIVING";DriverAtPickup Status="DRIVER_AT_PICKUP";PassengerOnboard Status="PASSENGER_ONBOARD";InProgress Status="IN_PROGRESS";Completed Status="COMPLETED";Cancelled Status="CANCELLED";Expired Status="EXPIRED";NoShow Status="NO_SHOW";Failed Status="FAILED")
type Booking struct{ID,PassengerID,PickupAddress,DropoffAddress string;ServiceType ServiceType;Status Status;RequestedAt,ScheduledAt time.Time;EstimatedCents int64;Currency string}
func(b Booking)Validate()error{if b.PassengerID==""||b.PickupAddress==""||b.DropoffAddress==""{return errors.New("passenger, pickup and dropoff are required")};if b.ServiceType==ServiceVTC&&b.ScheduledAt.IsZero(){return errors.New("VTC booking requires a reservation time")};return nil}
