package services

import (
	"backend/interfaces"
	"backend/models"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type AdminService struct {
	DB   *sql.DB
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
	if filter.PageSize == 0 {
		filter.PageSize = 10 // default 10 item per page
	}
	if filter.PageNumber == 0 {
		filter.PageNumber = 1 // default ke halaman pertama
	}

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
			query += item + " AND "
			mQuery += item + " AND "
		}
		query = strings.TrimSuffix(query, " AND ")
		mQuery = strings.TrimSuffix(mQuery, " AND ")
	}

	rows := s.DB.QueryRowContext(ctx, mQuery, sqlVars...)
	fmt.Println("mQuery:", mQuery)
	fmt.Println("countVars:", sqlVars)
	fmt.Println("query:", query)
	fmt.Println("sqlVars:", sqlVars)

	if err := rows.Scan(&totalSize); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to get total records"
		res.Data = nil
		return &res, nil
	}
	fmt.Println("FINAL SQL:", query)
	fmt.Println("FINAL VARS:", sqlVars)

	query += " LIMIT ? OFFSET ?"
	offset := (filter.PageNumber - 1) * filter.PageSize
	fmt.Println("Total records:", totalSize)
	fmt.Println("Offset:", offset)

	sqlVars = append(sqlVars, filter.PageSize, offset)

	result, err := s.DB.QueryContext(ctx, query, sqlVars...)
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "internal server error"
		res.Data = nil
		return &res, err
	}
	
	defer result.Close()
	fmt.Println("Query executed successfully")
	fmt.Println("Total records:", totalSize)
	fmt.Println("PageNumber:", filter.PageNumber)
	fmt.Println("PageSize:", filter.PageSize)

	for result.Next() {
		fmt.Println("Processing row")
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
		fmt.Printf("Scanned admin: %+v\n", admin)
		arr = append(arr, &admin)
	}
	if err := result.Err(); err != nil {
		fmt.Println("Rows error:", err)
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Error iterating result"
		res.Data = nil
		return &res, err
	}
	fmt.Println("All rows processed successfully")
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
		return &res, err
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
		data.Email, data.NoTelepon, data.IDRole)

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

func (s *AdminService) GetByID(ctx context.Context,admin *models.Admin, id uint) (*models.Response, error) {
	var res models.Response
	query := `SELECT * FROM admin WHERE id = ? AND deleted_at IS NULL`
	row := s.DB.QueryRowContext(ctx, query, id)

	if err := row.Scan(
		&admin.ID, &admin.Username, &admin.Password,
		&admin.NamaLengkap, &admin.Email, &admin.NoTelepon,
		&admin.IDRole, &admin.CreatedAt, &admin.UpdatedAt, &admin.DeletedAt,
	); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to scan admin data"
		res.Data = nil
		return &res, err
	}

	res.StatusCode = http.StatusOK
	res.Message = "Success"
	res.Data = admin
	return &res, nil
}

func (s *AdminService) Update(ctx context.Context, data *models.AdminUpdate) (*models.Response, error) {
	var res models.Response

	tx, err := s.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to begin transaction"
		res.Data = nil
		return &res, err
	}
	query := `	UPDATE admin SET username = ?, nama_lengkap = ?, email = ?, no_telepon = ?, id_role = ? WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to prepare statement"
		res.Data = nil
		return &res, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, data.Username, data.NamaLengkap, data.Email, data.NoTelepon, data.IDRole, data.ID)
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to execute update query"
		res.Data = nil
		return &res, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to get affected rows"
		res.Data = nil
		return &res, err
	}

	if err := tx.Commit(); err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Message = "Failed to commit transaction"
		res.Data = nil
		return &res, err
	}

	res.StatusCode = http.StatusOK
	res.Message = "Admin updated successfully"
	res.Data = rowsAffected
	return &res, nil
}

func (s *AdminService) Delete(id uint) error {
	query := `UPDATE admin SET deleted_at = NOW() WHERE id = ?`
	_, err := s.DB.Exec(query, id)
	return err
}
