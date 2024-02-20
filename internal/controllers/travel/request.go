package travel_controller

import "time"

type GetAllTravelRequest struct {
	Page      int        `form:"page" validate:"required,gt=0"`
	Limit     int        `form:"limit" validate:"required,gt=0"`
	StartDate *time.Time `form:"startDate" validate:"omitempty,date_format=2006-01-02"`
	EndDate   *time.Time `form:"endDate" validate:"omitempty,date_format=2006-01-02"`
}
