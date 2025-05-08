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

func validateTripIds(tripIds []int64) string {
	if len(tripIds) > 100 {
		return "trip_ids size can not be more than 100"
	}
	return ""
}

func validateTimeRanges(startTime int64, endTime int64) string {
	if startTime <= 0 {
		return "start_time is required"
	}
	if endTime <= 0 {
		return "end_time is required"
	}

	if (endTime - startTime) > (24 * 3600 * 30) {
		return "end_time - start_time should be less than 30 days"
	}

	return ""
}

func (params *GetTripFinanceDetailsRequest) Validate() string {
	if len(params.TripIDs) > 0 {
		return validateTripIds(params.TripIDs)
	}
	// start_time and end_time are required
	return validateTimeRanges(params.StartTime, params.EndTime)
}

func (params *GetTripFinanceDetailsRequest) ValidateForExportCSV() string {
	if len(params.TripIDs) > 0 {
		// if trip_ids exists, only validate trip_ids
		return validateTripIds(params.TripIDs)
	}
	// start_time and end_time are required
	return validateTimeRanges(params.StartTime, params.EndTime)
}
