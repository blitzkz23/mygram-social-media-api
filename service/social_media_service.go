package service

import (
	"mygram-social-media-api/dto"
	"mygram-social-media-api/pkg/errs"
	"mygram-social-media-api/repository/social_media_repository"
)

type SocialMediaService interface {
	AddSocialMedia(socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr)
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

func (s *socialMediaService) AddSocialMedia(socialMediaPayload *dto.SocialMediaRequest) (*dto.SocialMediaResponse, errs.MessageErr) {
	return nil, nil
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
