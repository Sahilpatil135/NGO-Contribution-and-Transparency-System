package repository

import (
	"context"
	"database/sql"

	"server/internal/models"
	// "github.com/google/uuid"
)

type DonationRepository interface {
	Create(ctx context.Context, donation *models.Donation) error

	// GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error)
	// GetByOrganizationID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	// GetByDomainID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	// GetByAidTypeID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	// GetAll(ctx context.Context) ([]*models.Donation, error)
	//
	// // Update(ctx context.Context, donation *models.Donation) error
	// Delete(ctx context.Context, id uuid.UUID) error
	//
	// GetDomains(ctx context.Context) ([]*models.DonationCategory, error)
	// GetAidTypes(ctx context.Context) ([]*models.DonationCategory, error)
	// GetDomainByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error)
	// GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error)
}

type donationRepository struct {
	db *sql.DB
}

func NewDonationRepository(db *sql.DB) DonationRepository {
	return &donationRepository{db: db}
}

func (c *donationRepository) Create(ctx context.Context, donation *models.Donation) error {
	query := `
		INSERT INTO donations (
			id, cause_id, user_id, name, phone, billing_address,
			pincode, amount, status, pan_number, payment_id, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := c.db.ExecContext(ctx, query,
		donation.ID,
		donation.CauseID,
		donation.UserID,
		donation.Name,
		donation.Phone,
		donation.BillingAddress,
		donation.Pincode,
		donation.Amount,
		donation.Status,
		donation.PanNumber,
		donation.PaymentID,
		donation.CreatedAt,
	)

	return err
}

// func (c *donationRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error) {
// 	query := `
// 		SELECT
// 			c.id, c.organization_id, c.title, c.description, c.collected_amount,
// 			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
// 			cd.id, cd.name, cd.description, cd.icon_url,
// 			ca.id, ca.name, ca.description, ca.icon_url
// 		FROM donations c
// 		LEFT JOIN donation_domains cd on c.domain_id = cd.id
// 		LEFT JOIN donation_aid_types ca on c.aid_type_id = ca.id
// 		WHERE c.id = $1
// 	`
//
// 	donation := &models.Donation{
// 		Domain:  models.DonationCategory{},
// 		AidType: models.DonationCategory{},
// 	}
//
// 	err := c.db.QueryRowContext(ctx, query, id).Scan(
// 		&donation.ID,
// 		&donation.OrganizationID,
// 		&donation.Title,
// 		&donation.Description,
// 		&donation.CollectedAmount,
// 		&donation.GoalAmount,
// 		&donation.Deadline,
// 		&donation.IsActive,
// 		&donation.CoverImageURL,
// 		&donation.CreatedAt,
// 		&donation.Domain.ID,
// 		&donation.Domain.Name,
// 		&donation.Domain.Description,
// 		&donation.Domain.IconURL,
// 		&donation.AidType.ID,
// 		&donation.AidType.Name,
// 		&donation.AidType.Description,
// 		&donation.AidType.IconURL,
// 	)
//
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("donation not found")
// 		}
// 		return nil, err
// 	}
//
// 	return donation, nil
// }
//
// func (c *donationRepository) GetByOrganizationID(ctx context.Context, organizationId uuid.UUID) ([]*models.Donation, error) {
// 	query := `
// 		SELECT
// 			c.id, c.organization_id, c.title, c.description, c.collected_amount,
// 			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
// 			cd.id, cd.name, cd.description, cd.icon_url,
// 			ca.id, ca.name, ca.description, ca.icon_url
// 		FROM donations c
// 		LEFT JOIN donation_domains cd on c.domain_id = cd.id
// 		LEFT JOIN donation_aid_types ca on c.aid_type_id = ca.id
// 		WHERE organization_id = $1
// 	`
//
// 	result, err := c.db.QueryContext(ctx, query, organizationId)
//
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
//
// 	var donationsResult []*models.Donation = make([]*models.Donation, 0, 5)
//
// 	for result.Next() {
// 		donation := &models.Donation{
// 			Domain:  models.DonationCategory{},
// 			AidType: models.DonationCategory{},
// 		}
//
// 		err = result.Scan(
// 			&donation.ID,
// 			&donation.OrganizationID,
// 			&donation.Title,
// 			&donation.Description,
// 			&donation.CollectedAmount,
// 			&donation.GoalAmount,
// 			&donation.Deadline,
// 			&donation.IsActive,
// 			&donation.CoverImageURL,
// 			&donation.CreatedAt,
// 			&donation.Domain.ID,
// 			&donation.Domain.Name,
// 			&donation.Domain.Description,
// 			&donation.Domain.IconURL,
// 			&donation.AidType.ID,
// 			&donation.AidType.Name,
// 			&donation.AidType.Description,
// 			&donation.AidType.IconURL,
// 		)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		donationsResult = append(donationsResult, donation)
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationRepository) GetByDomainID(ctx context.Context, domainID uuid.UUID) ([]*models.Donation, error) {
// 	query := `
// 		SELECT
// 			c.id, c.organization_id, c.title, c.description, c.collected_amount,
// 			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
// 			cd.id, cd.name, cd.description, cd.icon_url,
// 			ca.id, ca.name, ca.description, ca.icon_url
// 		FROM donations c
// 		LEFT JOIN donation_domains cd on c.domain_id = cd.id
// 		LEFT JOIN donation_aid_types ca on c.aid_type_id = ca.id
// 		WHERE domain_id = $1
// 	`
//
// 	result, err := c.db.QueryContext(ctx, query, domainID)
//
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
//
// 	var donationsResult []*models.Donation = make([]*models.Donation, 0, 5)
//
// 	for result.Next() {
// 		donation := &models.Donation{
// 			Domain:  models.DonationCategory{},
// 			AidType: models.DonationCategory{},
// 		}
//
// 		err = result.Scan(
// 			&donation.ID,
// 			&donation.OrganizationID,
// 			&donation.Title,
// 			&donation.Description,
// 			&donation.CollectedAmount,
// 			&donation.GoalAmount,
// 			&donation.Deadline,
// 			&donation.IsActive,
// 			&donation.CoverImageURL,
// 			&donation.CreatedAt,
// 			&donation.Domain.ID,
// 			&donation.Domain.Name,
// 			&donation.Domain.Description,
// 			&donation.Domain.IconURL,
// 			&donation.AidType.ID,
// 			&donation.AidType.Name,
// 			&donation.AidType.Description,
// 			&donation.AidType.IconURL,
// 		)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		donationsResult = append(donationsResult, donation)
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationRepository) GetByAidTypeID(ctx context.Context, aidTypeId uuid.UUID) ([]*models.Donation, error) {
// 	query := `
// 		SELECT
// 			c.id, c.organization_id, c.title, c.description, c.collected_amount,
// 			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
// 			cd.id, cd.name, cd.description, cd.icon_url,
// 			ca.id, ca.name, ca.description, ca.icon_url
// 		FROM donations c
// 		LEFT JOIN donation_domains cd on c.domain_id = cd.id
// 		LEFT JOIN donation_aid_types ca on c.aid_type_id = ca.id
// 		WHERE aid_type_id = $1
// 	`
//
// 	result, err := c.db.QueryContext(ctx, query, aidTypeId)
//
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
//
// 	var donationsResult []*models.Donation = make([]*models.Donation, 0, 5)
//
// 	for result.Next() {
// 		donation := &models.Donation{
// 			Domain:  models.DonationCategory{},
// 			AidType: models.DonationCategory{},
// 		}
//
// 		err = result.Scan(
// 			&donation.ID,
// 			&donation.OrganizationID,
// 			&donation.Title,
// 			&donation.Description,
// 			&donation.CollectedAmount,
// 			&donation.GoalAmount,
// 			&donation.Deadline,
// 			&donation.IsActive,
// 			&donation.CoverImageURL,
// 			&donation.CreatedAt,
// 			&donation.Domain.ID,
// 			&donation.Domain.Name,
// 			&donation.Domain.Description,
// 			&donation.Domain.IconURL,
// 			&donation.AidType.ID,
// 			&donation.AidType.Name,
// 			&donation.AidType.Description,
// 			&donation.AidType.IconURL,
// 		)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		donationsResult = append(donationsResult, donation)
// 	}
//
// 	return donationsResult, nil
// }
//
// func (c *donationRepository) GetAll(ctx context.Context) ([]*models.Donation, error) {
// 	query := `
// 		SELECT
// 			c.id, c.organization_id, c.title, c.description, c.collected_amount,
// 			c.goal_amount, c.deadline, c.is_active, c.cover_image_url, c.created_at,
// 			cd.id, cd.name, cd.description, cd.icon_url,
// 			ca.id, ca.name, ca.description, ca.icon_url
// 		FROM donations c
// 		LEFT JOIN donation_domains cd on c.domain_id = cd.id
// 		LEFT JOIN donation_aid_types ca on c.aid_type_id = ca.id
// 		WHERE is_active = true
// 	`
//
// 	result, err := c.db.QueryContext(ctx, query)
//
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
//
// 	var donationsResult []*models.Donation = make([]*models.Donation, 0, 5)
//
// 	for result.Next() {
// 		donation := &models.Donation{
// 			Domain:  models.DonationCategory{},
// 			AidType: models.DonationCategory{},
// 		}
//
// 		err = result.Scan(
// 			&donation.ID,
// 			&donation.OrganizationID,
// 			&donation.Title,
// 			&donation.Description,
// 			&donation.CollectedAmount,
// 			&donation.GoalAmount,
// 			&donation.Deadline,
// 			&donation.IsActive,
// 			&donation.CoverImageURL,
// 			&donation.CreatedAt,
// 			&donation.Domain.ID,
// 			&donation.Domain.Name,
// 			&donation.Domain.Description,
// 			&donation.Domain.IconURL,
// 			&donation.AidType.ID,
// 			&donation.AidType.Name,
// 			&donation.AidType.Description,
// 			&donation.AidType.IconURL,
// 		)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		donationsResult = append(donationsResult, donation)
// 	}
//
// 	return donationsResult, nil
// }
//
// // func (r *donationRepository) Update(ctx context.Context, donation *models.Donation) error { }
//
// func (c *donationRepository) Delete(ctx context.Context, id uuid.UUID) error {
// 	query := `UPDATE donations SET is_active = false WHERE id = $1`
//
// 	_, err := c.db.ExecContext(ctx, query, id)
//
// 	return err
// }
//
// func (c *donationRepository) GetDomains(ctx context.Context) ([]*models.DonationCategory, error) {
// 	query := `
// 		SELECT id, name, description, icon_url
// 		FROM donation_domains
// 	`
//
// 	result, err := c.db.QueryContext(ctx, query)
//
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
//
// 	var domainResults []*models.DonationCategory = make([]*models.DonationCategory, 0, 5)
//
// 	for result.Next() {
// 		domain := &models.DonationCategory{}
//
// 		err = result.Scan(
// 			&domain.ID,
// 			&domain.Name,
// 			&domain.Description,
// 			&domain.IconURL,
// 		)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		domainResults = append(domainResults, domain)
// 	}
//
// 	return domainResults, nil
// }
//
// func (c *donationRepository) GetAidTypes(ctx context.Context) ([]*models.DonationCategory, error) {
// 	query := `
// 		SELECT id, name, description, icon_url
// 		FROM donation_aid_types
// 	`
//
// 	result, err := c.db.QueryContext(ctx, query)
//
// 	if err != nil && err != sql.ErrNoRows {
// 		return nil, err
// 	}
//
// 	var aidTypeResults []*models.DonationCategory = make([]*models.DonationCategory, 0, 5)
//
// 	for result.Next() {
// 		aidType := &models.DonationCategory{}
//
// 		err = result.Scan(
// 			&aidType.ID,
// 			&aidType.Name,
// 			&aidType.Description,
// 			&aidType.IconURL,
// 		)
//
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		aidTypeResults = append(aidTypeResults, aidType)
// 	}
//
// 	return aidTypeResults, nil
// }
//
// func (c *donationRepository) GetDomainByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error) {
// 	query := `
// 		SELECT id, name, description, icon_url
// 		FROM donation_domains
// 		WHERE id = $1
// 	`
//
// 	domain := &models.DonationCategory{}
//
// 	err := c.db.QueryRowContext(ctx, query, id).Scan(
// 		&domain.ID,
// 		&domain.Name,
// 		&domain.Description,
// 		&domain.IconURL,
// 	)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return domain, err
// }
//
// func (c *donationRepository) GetAidTypeByID(ctx context.Context, id uuid.UUID) (*models.DonationCategory, error) {
// 	query := `
// 		SELECT id, name, description, icon_url
// 		FROM donation_aid_types
// 		WHERE id = $1
// 	`
//
// 	aidType := &models.DonationCategory{}
//
// 	err := c.db.QueryRowContext(ctx, query, id).Scan(
// 		&aidType.ID,
// 		&aidType.Name,
// 		&aidType.Description,
// 		&aidType.IconURL,
// 	)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return aidType, err
// }
