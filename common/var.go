package common

import (
	"errors"
)

var (
	ErrRecordNotFound                 = errors.New("record not found")
	ErrNoRecordsAreChanged            = errors.New("no records are changed")
	ErrEmailNotFound                  = errors.New("email not found")
	ErrWrongPassword                  = errors.New("wrong password")
	ErrNurseNotAvailable              = errors.New("this nurse are not available")
	ErrNursesNotAvailable             = errors.New("this nurses are not available")
	ErrAppointmentStatusIsNotUpcoming = errors.New("appointment status is not upcoming")
	ErrCustaskNotDoneAll              = errors.New("there are still pending custasks that haven't been completed")
)
