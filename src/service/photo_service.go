package service

import (
	"fmt"
	"mygram-social-media-api/src/dto"
	"mygram-social-media-api/src/entity"
	"mygram-social-media-api/src/pkg/errs"
	"mygram-social-media-api/src/pkg/helpers"
	"mygram-social-media-api/src/repository/photo_repository"
)

type PhotoService interface {
	PostPhoto(userID uint, photoPayload *dto.PhotoRequest) (*dto.PhotoResponse, errs.MessageErr)
	GetAllPhotos() ([]*dto.GetPhotoResponse, errs.MessageErr)
	EditPhotoData(photoID uint, photoPayload *dto.PhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr)
	DeletePhoto(photoID uint) (*dto.DeletePhotoResponse, errs.MessageErr)
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
	dto := make([]*dto.GetPhotoResponse, 0)
	photos, err := p.photoRepository.GetAllPhotos()
	if err != nil {
		return nil, err
	}

	for _, photo := range photos {
		dto = append(dto, photo.ToGetPhotoResponseDTO())
	}

	return dto, nil
}

func (p *photoService) EditPhotoData(photoID uint, photoPayload *dto.PhotoRequest) (*dto.UpdatePhotoResponse, errs.MessageErr) {
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

	response := &dto.UpdatePhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		UpdatedAt: photo.UpdatedAt,
	}

	return response, nil
}

func (p *photoService) DeletePhoto(photoID uint) (*dto.DeletePhotoResponse, errs.MessageErr) {
	_, err := p.photoRepository.GetPhotoByID(photoID)
	if err != nil {
		return nil, err
	}

	err = p.photoRepository.DeletePhoto(photoID)
	if err != nil {
		return nil, err
	}

	response := &dto.DeletePhotoResponse{
		Message: "Your photo has been deleted",
	}

	fmt.Println("Melihat response di service: ", response)
	return response, nil
}
