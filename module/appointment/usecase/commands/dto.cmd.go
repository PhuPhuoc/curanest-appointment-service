package apppointmentcommands

import "github.com/google/uuid"

type DistanceMatrixResponse struct {
	Rows []Row `json:"rows"`
}

type Row struct {
	Elements []Element `json:"elements"`
}

type Element struct {
	Distance Distance `json:"distance"`
	Duration Duration `json:"duration"`
	Status   string   `json:"status"`
}

type Distance struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type Duration struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}

type RelativesData struct {
	RelativesID uuid.UUID `json:"relatives-id"`
}

type RelativesResponse struct {
	Data    RelativesData `json:"data"`
	Success bool          `json:"success"`
}

type PatientInfo struct {
	ID            string `json:"id"`
	FullName      string `json:"full-name"`
	Gender        bool   `json:"gender"`
	DOB           string `json:"dob"`
	PhoneNumber   string `json:"phone-number"`
	Address       string `json:"address"`
	Ward          string `json:"ward"`
	District      string `json:"district"`
	City          string `json:"city"`
	DescPathology string `json:"desc-pathology"`
	NoteForNurse  string `json:"note-for-nurse"`
}

type NurseProfile struct {
	NurseID          string  `json:"nurse-id"`
	NursePicture     string  `json:"nurse-picture"`
	NurseName        string  `json:"nurse-name"`
	City             string  `json:"city"`
	CurrentWorkPlace string  `json:"current-work-place"`
	EducationLevel   string  `json:"education-level"`
	Experience       string  `json:"experience"`
	Certificate      string  `json:"certificate"`
	Slogan           string  `json:"slogan"`
	Rate             float64 `json:"rate"`
}
