package repository

import (
	"context"
	"database/sql"
	"fmt"

	"server/internal/models"

	"github.com/google/uuid"
)

type DonationRepository interface {
	Create(ctx context.Context, donation *models.Donation) error

	GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error)
	GetByCauseID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error)
	GetByPaymentID(ctx context.Context, id uuid.UUID) (*models.Donation, error)

	// Update(ctx context.Context, donation *models.Donation) error
	// Delete(ctx context.Context, id uuid.UUID) error
}

type donationRepository struct {
	db *sql.DB
}

func NewDonationRepository(db *sql.DB) DonationRepository {
	return &donationRepository{db: db}
}

func (d *donationRepository) Create(ctx context.Context, donation *models.Donation) error {
	query := `
		INSERT INTO donations (
			id, cause_id, user_id, name, phone, billing_address,
			pincode, amount, status, pan_number, payment_id, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := d.db.ExecContext(ctx, query,
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

func GetDonationByColumnID(d *donationRepository, ctx context.Context, ID uuid.UUID, column string) (*models.Donation, error) {
	query := fmt.Sprintf(`
		SELECT
			c.id, c.cause_id, c.user_id, c.name,
			c.phone, c.billing_address, c.pincode,
			c.amount, c.status, c.pan_number,
			c.payment_id, c.created_at
		FROM donations c
		WHERE c.%s = $1
		`, column)

	donation := models.Donation{}

	err := d.db.QueryRowContext(ctx, query, ID).Scan(
		&donation.ID,
		&donation.CauseID,
		&donation.UserID,
		&donation.Name,
		&donation.Phone,
		&donation.BillingAddress,
		&donation.Pincode,
		&donation.Amount,
		&donation.Status,
		&donation.PanNumber,
		&donation.PaymentID,
		&donation.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &donation, nil
}

func GetDonationsByColumnID(d *donationRepository, ctx context.Context, ID uuid.UUID, column string) ([]*models.Donation, error) {
	query := fmt.Sprintf(`
		SELECT
			c.id, c.cause_id, c.user_id, c.name,
			c.phone, c.billing_address, c.pincode,
			c.amount, c.status, c.pan_number,
			c.payment_id, c.created_at
		FROM donations c
		WHERE c.%s = $1
		`, column)

	result, err := d.db.QueryContext(ctx, query, ID)

	var donationsResult []*models.Donation = make([]*models.Donation, 0, 5)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("donations not found")
		}
		return nil, err
	}

	for result.Next() {
		donation := &models.Donation{}

		err = result.Scan(
			&donation.ID,
			&donation.CauseID,
			&donation.UserID,
			&donation.Name,
			&donation.Phone,
			&donation.BillingAddress,
			&donation.Pincode,
			&donation.Amount,
			&donation.Status,
			&donation.PanNumber,
			&donation.PaymentID,
			&donation.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		donationsResult = append(donationsResult, donation)
	}

	return donationsResult, nil
}

func (d *donationRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Donation, error) {
	return GetDonationByColumnID(d, ctx, id, "id")
}

func (d *donationRepository) GetByCauseID(ctx context.Context, id uuid.UUID) ([]*models.Donation, error) {
	return GetDonationsByColumnID(d, ctx, id, "cause_Id")
}

func (d *donationRepository) GetByPaymentID(ctx context.Context, id uuid.UUID) (*models.Donation, error) {
	return GetDonationByColumnID(d, ctx, id, "payment_id")
}

// // func (r *donationRepository) Update(ctx context.Context, donation *models.Donation) error { }
//
// func (c *donationRepository) Delete(ctx context.Context, id uuid.UUID) error {
// 	query := `UPDATE donations SET is_active = false WHERE id = $1`
//
// 	_, err := c.db.ExecContext(ctx, query, id)
//
// 	return err
// }
