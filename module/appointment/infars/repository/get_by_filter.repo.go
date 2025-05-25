package appointmentrepository

import (
	"context"
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

func (repo *appointmentRepo) GetAppointment(ctx context.Context, filter *appointmentqueries.FilterGetAppointmentDTO) ([]appointmentdomain.Appointment, error) {
	var whereConditions []string
	var args []interface{}

	if filter != nil {
		if filter.Id != nil && filter.Id.String() != "" {
			whereConditions = append(whereConditions, "id = ?")
			args = append(args, filter.Id.String())
		}
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
	var argsLimit []interface{}
	if filter.ApplyPaging != nil && filter.Paging != nil {
		limit = " limit ? offset ?"
		if *filter.ApplyPaging {
			argsLimit = append(argsLimit, filter.Paging.Size, (filter.Paging.Page-1)*filter.Paging.Size)
		}
	}

	var where string
	if len(whereConditions) > 0 {
		where = strings.Join(whereConditions, " AND ")
	}

	queryGetData := common.GenerateSQLQueries(common.SELECT_WITHOUT_COUNT, TABLE_APPOINTMENT, GET_APPOINTMENT, &where)
	queryGetData += orderBy + limit
	queryGetCount := common.GenerateSQLQueries(common.SELECT_COUNT, TABLE_APPOINTMENT, GET_APPOINTMENT, &where)

	errchan := make(chan error, 2)
	countchan := make(chan int, 1)
	datachan := make(chan []appointmentdomain.Appointment, 1)
	argsOfGetData := append(args, argsLimit...)

	var wg sync.WaitGroup
	wg.Add(2)
	go repo.getCount(ctx, queryGetCount, args, errchan, countchan, &wg)
	go repo.getData(ctx, queryGetData, argsOfGetData, errchan, datachan, &wg)

	var once sync.Once // Make sure to close the channel only once.
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errchan <- fmt.Errorf("panic in goroutine: %v", r)
			}
		}()

		wg.Wait()
		once.Do(func() {
			close(errchan)
			close(countchan)
			close(datachan)
		})
	}()

	var totalRecord int
	var appointments []appointmentdomain.Appointment

	receivedCount := 0
	expectedCount := 2

	for {
		select {
		case err, ok := <-errchan:
			if ok {
				return nil, err
			}
		case count, ok := <-countchan:
			if ok {
				totalRecord = count
				receivedCount++
			}
		case data, ok := <-datachan:
			if ok {
				appointments = data
				receivedCount++
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("operation timed out: %w", ctx.Err())
		}

		if receivedCount == expectedCount {
			break
		}
	}

	if filter.ApplyPaging != nil && filter.Paging != nil {
		totalPages := int(math.Ceil(float64(totalRecord) / float64(filter.Paging.Size)))
		filter.Paging.Total = totalPages
	}
	return appointments, nil
}

func (repo *appointmentRepo) getData(
	ctx context.Context,
	queryStr string,
	args []interface{},
	errchan chan<- error,
	datachan chan<- []appointmentdomain.Appointment,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var dtos []AppointmentDTO
	if err := repo.db.SelectContext(ctx, &dtos, queryStr, args...); err != nil {
		errchan <- err
		return
	}

	entities := make([]appointmentdomain.Appointment, len(dtos))
	for i := range dtos {
		entity, _ := dtos[i].ToAppointmentEntity()
		entities[i] = *entity
	}

	datachan <- entities
}

func (repo *appointmentRepo) getCount(
	ctx context.Context,
	queryStr string,
	args []interface{},
	errchan chan<- error,
	countchan chan<- int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	var total int
	if err := repo.db.GetContext(ctx, &total, queryStr, args...); err != nil {
		errchan <- err
		return
	}

	countchan <- total
}
