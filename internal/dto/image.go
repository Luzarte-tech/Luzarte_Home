package dto

type UploadImageRequest struct {
	PropertyID string `form:"property_id" binding:"required"`
}

type SetPrimaryImageRequest struct {
	IsPrimary bool `json:"is_primary"`
}