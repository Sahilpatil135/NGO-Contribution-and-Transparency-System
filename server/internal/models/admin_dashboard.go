package models

type AdminDashboardStats struct {
	TotalOrganizations int64 `json:"total_organizations"`
	TotalDonors        int64 `json:"total_donors"`
	TotalCauses        int64 `json:"total_causes"`
	TotalDonations     int64 `json:"total_donations"`
}

type AdminDashboardData struct {
	Stats             AdminDashboardStats `json:"stats"`
	OrganizationNames []string            `json:"organization_names"`
	DonorNames        []string            `json:"donor_names"`
}
