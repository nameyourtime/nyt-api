package pg

import (
	"database/sql"
	"github.com/lib/pq"
	"nameyourtime.com/api/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Create(user *models.User) (string, error) {
	stmt := `insert into user (name, email, password, status, refresh_token, refresh_token_exp) 
             values ($1, $2, $3, $4, $5, $6) returning id;`
	var id string
	err := m.DB.QueryRow(stmt, user.Name, user.Email, user.Password, user.Status,
		user.Token.RefreshToken, user.Token.RefreshTokenExp).Scan(&id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if "unique_violation" == err.Code.Name() {
				return "", models.ErrNonUniqueEmail
			}
		}
		return "", err
	}
	return id, nil
}

func (m *UserModel) Get(userID string) (*models.User, error) {
	stmt := `select id, name, email, password, status, created, refresh_token, refresh_token_exp 
             from flcrd.user where id = $1;`
	d := &models.User{}
	err := m.DB.QueryRow(stmt, userID).Scan(&d.ID, &d.Name, &d.Email, &d.Password, &d.Status, &d.Created,
		&d.Token.RefreshToken, &d.Token.RefreshTokenExp)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	}
	if err != nil {
		return nil, err
	}
	d.Created = d.Created.UTC()
	d.Token.RefreshTokenExp = d.Token.RefreshTokenExp.UTC()
	return d, nil
}

func (m *UserModel) GetByEmail(email string) (*models.User, error) {
	stmt := `select id, name, email, password, status, created, refresh_token, refresh_token_exp 
             from flcrd.user where email = $1;`
	d := &models.User{}
	err := m.DB.QueryRow(stmt, email).Scan(&d.ID, &d.Name, &d.Email, &d.Password, &d.Status, &d.Created,
		&d.Token.RefreshToken, &d.Token.RefreshTokenExp)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	}
	if err != nil {
		return nil, err
	}
	d.Created = d.Created.UTC()
	return d, nil
}
