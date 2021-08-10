package mapper

import (
	"github.com/xdorro/golang-fiber-base-project/app/entity/dto"
	"github.com/xdorro/golang-fiber-base-project/app/entity/response"
)

func SearchBannerMapper(banners *[]dto.SearchBannerDTO) *[]response.SearchBannerResponse {
	result := make([]response.SearchBannerResponse, 0)

	for _, movie := range *banners {
		mapper := SearchBanner(&movie)
		result = append(result, *mapper)
	}

	return &result
}

func SearchBanner(banner *dto.SearchBannerDTO) *response.SearchBannerResponse {
	return &response.SearchBannerResponse{
		BannerId: banner.BannerId,
		Status:   banner.Status,
		Image:    banner.Image,
		Movie: response.SearchMovieResponse{
			MovieId:     banner.MovieId,
			Name:        banner.Name,
			OriginName:  banner.OriginName,
			Slug:        banner.Slug,
			ReleaseDate: banner.ReleaseDate,
		},
	}
}
