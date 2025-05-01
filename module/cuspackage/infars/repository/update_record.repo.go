package cuspackagerepository

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagedomain "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/domain"
)

func (repo *cusPackageRepo) UpdateMedicalRecord(ctx context.Context, entity *cuspackagedomain.MedicalRecord) error {
	dto := ToMedicalRecordDTO(entity)
	where := "id=:id"
	query := common.GenerateSQLQueries(common.UPDATE, TABLE_MEDICALRECORD, UPDATE_MEDICALRECORD, &where)
	if _, err := repo.db.NamedExec(query, dto); err != nil {
		return err
	}

	return nil
}
