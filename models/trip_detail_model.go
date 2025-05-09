package models

type GetTripFinanceDetailsRequest struct {
	TripIDs        []int64
	DriverId       int64
	UserId         int64
	Limit          int
	Offset         int
	StartTime      int64
	EndTime        int64
	BypassMetaData bool
	PaymentMethod  int
	VehicleType    int
	DateType       int
	Year           string
}

func (params *GetTripFinanceDetailsRequest) Validate() string {
	if params.TripIDs == nil {
		return "trip_ids parse error, please check"
	}

	if len(params.TripIDs) > 100 {
		return "trip_ids size can not be more than 100"
	}

	if len(params.TripIDs) > 0 {
		// search by trip_ids
		return ""
	}

	if params.StartTime <= 0 {
		return "start_time is required"
	}
	if params.EndTime <= 0 {
		return "end_time is required"
	}

	if (params.EndTime - params.StartTime) > (24 * 3600 * 30) {
		return "end_time - start_time should be less than 30 days"
	}

	if params.DriverId > 0 {
		// search by driver_id
		return ""
	}

	if params.VehicleType >= 0 {
		// search by vehicle_type
		return ""
	}

	if params.PaymentMethod >= 0 {
		// search by payment_method
		return ""
	}

	// The 'end_trip_time' parameter must be combined with at least one other parameter to perform a search
	return "end_trip_time must be combined with at least one other parameter to perform a search"
}
