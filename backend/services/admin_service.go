package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"net/http"
	"strings"
)

type AdminService struct {
	DB *sql.DB
	_cfg any
}

var _ interfaces.AdminInterface = &AdminService{}
func NewAdminService(db *sql.DB) *AdminService {
	return &AdminService{DB: db}
}

func (s *AdminService) GetAll(ctx context.Context, filter *models.AdminFilter) (*models.Response, error) {
	var res models.Response
	var totalSize int64
	arr := []*models.Admin{}

	query := `select * from admin`

	mQuery := `select count(id) as sample from admin`

	var sqlFilters []string
	var sqlVars []any


	if filter.Username != "" {
		sqlFilters = append(sqlFilters, "username LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.Username+"%")
	}
	if filter.NamaLengkap != "" {
		sqlFilters = append(sqlFilters, "nama_lengkap LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.NamaLengkap+"%")
	}
	if filter.Email != "" {
		sqlFilters = append(sqlFilters, "email LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.Email+"%")
	}
	if filter.NoTelepon != "" {
		sqlFilters = append(sqlFilters, "no_telepon LIKE ?")
		sqlVars = append(sqlVars, "%"+filter.NoTelepon+"%")
	}
	if filter.IDRole != 0 {
		sqlFilters = append(sqlFilters, "id_role = ?")
		sqlVars = append(sqlVars, filter.IDRole)
	}

	if len(sqlFilters) > 0 {
		query += " WHERE "
		mQuery += " WHERE "
		for _, item := range sqlFilters {
			query += item + " AND"
			mQuery += item + " AND"
		}
		query = strings.TrimSuffix(query, " AND")
		mQuery = strings.TrimSuffix(mQuery, " AND")
	}

	rows := s.DB.QueryRowContext(ctx, mQuery, sqlVars...)
	if err := rows.Scan(&totalSize); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to get total records"
		res.Data = nil
		return &res, nil
	}

	// Setelah query filter ditambahkan
	query += " LIMIT ? OFFSET ?"
	offset := (filter.PageNumber - 1) * filter.PageSize
	sqlVars = append(sqlVars, filter.PageSize, offset)


	result, err := s.DB.QueryContext(ctx, query, sqlVars...)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "internal server error"
		res.Data = nil
		return &res, err
	}
	defer result.Close()
	
	for result.Next() {
		var admin models.Admin
		if err := result.Scan(
			&admin.ID, &admin.Username, &admin.Password,
			&admin.NamaLengkap, &admin.Email, &admin.NoTelepon,
			&admin.IDRole, &admin.CreatedAt, &admin.UpdatedAt,
		); err != nil {
			res.StatusCode = http.StatusInternalServerError
			res.Message = "Failed to scan admin data"
			res.Data = nil
			return &res, err
		}
		arr = append(arr, &admin)
	}
	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = arr
	return &res, nil
}
	
func (s *AdminService) Create(ctx context.Context, data *models.AdminCreate) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		res.Data = nil
		return nil, nil
	}

	query := `
		INSERT INTO admin (username, password, nama_lengkap, email, no_telepon, id_role)
		VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, nil
	}
	defer stmt.Close()


	result, err := s.DB.ExecContext(ctx, query,
		data.Username, data.Password, data.NamaLengkap,
		data.Email, data.NoTelepon, data.IDRole,)

	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute query"
		res.Data = nil
		return &res, nil
	}

	lastID, _ := result.LastInsertId()

	if err := tx.Commit(); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to commit transaction"
		res.Data = nil
		return &res, nil
	}

	res.StatusCode = http.StatusCreated
	res.Message = "Admin created successfully"
	res.Data = lastID
	return &res, nil
}

func (s *AdminService) GetByID(id uint) (*models.Admin, error) {
	query := `
		SELECT id, username, password, nama_lengkap, email, no_telepon, id_role, created_at, updated_at
		FROM admin WHERE id = ?`
	row := s.DB.QueryRow(query, id)

	var result models.Admin
	err := row.Scan(
		&result.ID, &result.Username, &result.Password,
		&result.NamaLengkap, &result.Email, &result.NoTelepon,
		&result.IDRole, &result.CreatedAt, &result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *AdminService) Update(data *models.Admin) error {
	query := `
		UPDATE admin SET username = ?, password = ?, nama_lengkap = ?, email = ?, no_telepon = ?, id_role = ?, updated_at = ?
		WHERE id = ?`
	_, err := s.DB.Exec(query,
		data.Username, data.Password, data.NamaLengkap,
		data.Email, data.NoTelepon, data.IDRole,
		data.UpdatedAt, data.ID,
	)
	return err
}

func (s *AdminService) Delete(id uint) error {
	query := `UPDATE admin SET deleted_at = NOW() WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
