package cuspackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

func (repo *cusPackageRepo) FindMedicalRecordByAppsId(ctx context.Context, appsId uuid.UUID) (*cuspackagedomain.MedicalRecord, error) {
	var dto MedicalRecordDTO
	where := "appointment_id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE_MEDICALRECORD, GET_MEDICALRECORD, &where)
	if err := repo.db.Get(&dto, query, appsId); err != nil {
		return nil, err
	}
	return dto.ToMedicalRecordEntity()
}
