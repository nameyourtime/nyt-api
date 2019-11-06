package pg

import (
	"database/sql"
	"github.com/lib/pq"
	"nameyourtime.com/api/pkg/models"
)

type VerificationModel struct {
	DB *sql.DB
}

func (m *VerificationModel) Create(code models.VerificationCode) (string, error) {
	stmt := `insert into verification_code (user_id, code, code_exp) values ($1, $2, $3) returning code;`
	var c string
	err := m.DB.QueryRow(stmt, code.UserID, code.Code, code.CodeExp).Scan(&c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if "unique_violation" == err.Code.Name() {
				return "", models.ErrNonUniqueCode
			}
		}
		return "", err
	}
	return c, nil
}
