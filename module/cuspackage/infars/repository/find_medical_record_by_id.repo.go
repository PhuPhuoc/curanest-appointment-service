package cuspackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
	"github.com/google/uuid"
)

func (repo *cusPackageRepo) FindMedicalRecordById(ctx context.Context, mecicalRecordId uuid.UUID) (*cuspackagedomain.MedicalRecord, error) {
	var dto MedicalRecordDTO
	where := "id=?"
	query := common.GenerateSQLQueries(common.FIND, TABLE_MEDICALRECORD, GET_MEDICALRECORD, &where)
	if err := repo.db.Get(&dto, query, mecicalRecordId); err != nil {
		return nil, err
	}
	return dto.ToMedicalRecordEntity()
}
