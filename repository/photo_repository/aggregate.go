package photo_repository

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
)

type PhotoWithUser struct {
	Photo entity.Photo
	User  entity.User
}

func (p PhotoWithUser) ToGetPhotoResponseDTO() dto.GetPhotoResponse {
	photoWithUserTemp := dto.GetPhotoResponse{
		ID:        p.Photo.ID,
		Title:     p.Photo.Title,
		Caption:   p.Photo.Caption,
		PhotoURL:  p.Photo.PhotoURL,
		UserID:    p.Photo.UserID,
		CreatedAt: p.Photo.CreatedAt,
		UpdatedAt: p.Photo.UpdatedAt,
		User: dto.EmbeddedUserResponse{
			Username: p.User.Username,
			Email:    p.User.Email,
		},
	}

	return photoWithUserTemp
}
