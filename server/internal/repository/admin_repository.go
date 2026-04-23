package repository

import (
	"context"
	"database/sql"

	"server/internal/models"
)

type AdminRepository interface {
	GetDashboardData(ctx context.Context) (*models.AdminDashboardData, error)
}

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetDashboardData(ctx context.Context) (*models.AdminDashboardData, error) {
	stats := models.AdminDashboardStats{}

	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM organizations`).Scan(&stats.TotalOrganizations); err != nil {
		return nil, err
	}
	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE role = 'user' AND is_active = true`).Scan(&stats.TotalDonors); err != nil {
		return nil, err
	}
	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM causes`).Scan(&stats.TotalCauses); err != nil {
		return nil, err
	}
	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM donations`).Scan(&stats.TotalDonations); err != nil {
		return nil, err
	}

	organizationNames := make([]string, 0)
	orgRows, err := r.db.QueryContext(ctx, `
		SELECT organization_name
		FROM organizations
		WHERE organization_name IS NOT NULL
		ORDER BY organization_name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer orgRows.Close()

	for orgRows.Next() {
		var name string
		if err := orgRows.Scan(&name); err != nil {
			return nil, err
		}
		organizationNames = append(organizationNames, name)
	}
	if err := orgRows.Err(); err != nil {
		return nil, err
	}

	donorNames := make([]string, 0)
	donorRows, err := r.db.QueryContext(ctx, `
		SELECT name
		FROM users
		WHERE role = 'user'
		  AND is_active = true
		  AND name IS NOT NULL
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, err
	}
	defer donorRows.Close()

	for donorRows.Next() {
		var name string
		if err := donorRows.Scan(&name); err != nil {
			return nil, err
		}
		donorNames = append(donorNames, name)
	}
	if err := donorRows.Err(); err != nil {
		return nil, err
	}

	return &models.AdminDashboardData{
		Stats:             stats,
		OrganizationNames: organizationNames,
		DonorNames:        donorNames,
	}, nil
}
