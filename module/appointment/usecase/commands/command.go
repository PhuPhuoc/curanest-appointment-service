package apppointmentcommands

type Commands struct{}

type Builder interface {
	BuildAppointmentCmdRepo() AppointmentCommandRepo
}

func NewAppointmentCmdWithBuilder(b Builder) Commands {
	return Commands{}
}

type AppointmentCommandRepo interface{}
