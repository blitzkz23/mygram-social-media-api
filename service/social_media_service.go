package service

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/entity"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/pkg/helpers"
	"mygram-social-media-api/repository/social_media_repository"
)

type SocialMediaService interface {
	AddSocialMedia(userID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr)
	GetAllSocialMedias() ([]*dto.GetSocialMediaResponse, errs.MessageErr)
	EditSocialMediaData(socialMediaID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(socialMediaID uint) errs.MessageErr
}

type socialMediaService struct {
	socialMediaRepository social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepository: socialMediaRepository}
}

func (s *socialMediaService) AddSocialMedia(userID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr) {
	if err := helpers.ValidateStruct(socialMediaPayload); err != nil {
		return nil, err
	}

	entity := entity.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaURL: socialMediaPayload.SocialMediaURL,
		UserID:         userID,
	}

	socialMedia, err := s.socialMediaRepository.AddSocialMedia(&entity)
	if err != nil {
		return nil, err
	}

	response := &dto.SocialMediaResponse{
		ID:             socialMedia.ID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		UserID:         socialMedia.UserID,
		CreatedAt:      socialMedia.CreatedAt,
	}

	return response, nil
}

func (s *socialMediaService) GetAllSocialMedias() ([]*dto.GetSocialMediaResponse, errs.MessageErr) {
	return nil, nil
}

func (s *socialMediaService) EditSocialMediaData(socialMediaID uint, socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr) {
	return nil, nil
}

func (s *socialMediaService) DeleteSocialMedia(socialMediaID uint) errs.MessageErr {
	return nil
}
