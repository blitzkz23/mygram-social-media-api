package service

import (
	"fmt"
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/repository/photo_repository"
)

type PhotoService interface {
	PostPhoto(userID uint, photoPayload *dto.PhotoRequest) (*dto.PhotoResponse, errs.MessageErr)
	GetAllPhotos() ([]*dto.GetPhotoResponse, errs.MessageErr)
	EditPhotoData(photoID uint, photoPayload *dto.PhotoRequest) (*dto.PhotoUpdateResponse, errs.MessageErr)
	DeletePhoto(photoID uint) *dto.DeletePhotoResponse
}

type photoService struct {
	photoRepository photo_repository.PhotoRepository
}

func NewPhotoService(photoRepository photo_repository.PhotoRepository) PhotoService {
	return &photoService{photoRepository: photoRepository}
}

func (p *photoService) PostPhoto(userID uint, photoPayload *dto.PhotoRequest) (*dto.PhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)
	if err != nil {
		return nil, err
	}

	payload := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoURL: photoPayload.PhotoURL,
		UserID:   userID,
	}

	photo, err := p.photoRepository.PostPhoto(payload)
	if err != nil {
		return nil, err
	}

	response := &dto.PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	return response, nil
}

func (p *photoService) GetAllPhotos() ([]*dto.GetPhotoResponse, errs.MessageErr) {
	photos, err := p.photoRepository.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	return photos, nil
}

func (p *photoService) EditPhotoData(photoID uint, photoPayload *dto.PhotoRequest) (*dto.PhotoUpdateResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(photoPayload)
	if err != nil {
		return nil, err
	}

	payload := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoURL: photoPayload.PhotoURL,
	}

	photo, err := p.photoRepository.EditPhotoData(photoID, payload)
	if err != nil {
		return nil, err
	}
	fmt.Println("Melihat photo di service: ", photo)

	response := &dto.PhotoUpdateResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}
	fmt.Println("Melihat response di service: ", response)

	return response, nil
}

func (p *photoService) DeletePhoto(photoID uint) *dto.DeletePhotoResponse {
	return nil
}
