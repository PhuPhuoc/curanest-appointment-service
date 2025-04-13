package appointmentrepository

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

func (repo *appointmentRepo) GetAppointment(ctx context.Context, filter *appointmentqueries.FilterGetAppointmentDTO) ([]appointmentdomain.Appointment, error) {
	var whereConditions []string
	var args []interface{}

	if filter != nil {
		if filter.ServiceId != nil && filter.ServiceId.String() != "" {
			whereConditions = append(whereConditions, "service_id = ?")
			args = append(args, filter.ServiceId.String())
		}
		if filter.CusPackageId != nil && filter.CusPackageId.String() != "" {
			whereConditions = append(whereConditions, "customized_package_id = ?")
			args = append(args, filter.CusPackageId.String())
		}
		if filter.NursingId != nil && filter.NursingId.String() != "" {
			whereConditions = append(whereConditions, "nursing_id = ?")
			args = append(args, filter.NursingId.String())
		}
		if filter.PatientId != nil && filter.PatientId.String() != "" {
			whereConditions = append(whereConditions, "patient_id = ?")
			args = append(args, filter.PatientId.String())
		}
		if filter.HadNurse != nil {
			if *filter.HadNurse {
				whereConditions = append(whereConditions, "nursing_id is not null")
			} else {
				whereConditions = append(whereConditions, "nursing_id is null")
			}
		}
		if filter.AppointmentStatus != nil && filter.AppointmentStatus.String() != "" {
			whereConditions = append(whereConditions, "status = ?")
			args = append(args, filter.AppointmentStatus.String())
		}
		if filter.EstDateFrom != nil && !filter.EstDateFrom.IsZero() {
			whereConditions = append(whereConditions, "est_date >= ?")
			args = append(args, filter.EstDateFrom.Format("2006-01-02"))
		}
		if filter.EstDateTo != nil && !filter.EstDateTo.IsZero() {
			whereConditions = append(whereConditions, "est_date <= ?")
			args = append(args, filter.EstDateTo.Format("2006-01-02"))
		}
	}

	orderBy := " order by est_date desc "

	limit := ""
	if filter.ApplyPaging != nil && filter.Paging != nil {
		limit = " limit ? offset ?"
		if *filter.ApplyPaging {
			args = append(args, filter.Paging.Size, (filter.Paging.Page-1)*filter.Paging.Size)
		}
	}

	var where string
	if len(whereConditions) > 0 {
		where = strings.Join(whereConditions, " AND ")
	}
	query := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_APPOINTMENT, GET_APPOINTMENT, &where)

	query += orderBy + limit

	fmt.Println("query: ", query)
	var dtos []AppointmentDTO
	if err := repo.db.SelectContext(ctx, &dtos, query, args...); err != nil {
		return nil, err
	}

	entities := make([]appointmentdomain.Appointment, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToAppointmentEntity()
		entities[i] = *entity
	}

	if filter.ApplyPaging != nil && filter.Paging != nil {
		totalRecord := 10
		totalPages := int(math.Ceil(float64(totalRecord) / float64(filter.Paging.Size)))
		filter.Paging.Total = totalPages
	}
	return entities, nil
}
