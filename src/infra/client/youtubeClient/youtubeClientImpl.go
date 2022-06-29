package youtubeClient

import (
	"github.com/kkdai/youtube/v2"
	"github.com/snxl/youtube-dowloader/src/domain/entity"
)

type YoutubeClient struct {
	client youtube.Client
}

func NewYoutubeClientImpl(client youtube.Client) *YoutubeClient {
	return &YoutubeClient{
		client: client,
	}
}

func (y *YoutubeClient) GetVideoInfos(id string) ([]entity.Video, error) {
	video, err := y.client.GetVideo(id)
	if err != nil {
		return nil, err
	}

	videoArr := make([]entity.Video, 0)

	for _, element := range video.Formats {
		videoInfo := entity.Video{
			Quality:      element.Quality,
			QualityLabel: element.QualityLabel,
			AudioChannel: element.AudioChannels != 0,
			Title:        video.Title,
			Author:       video.Author,
			PublishDate:  video.PublishDate,
			Duration:     video.Duration,
		}
		videoArr = append(videoArr, videoInfo)
	}

	return videoArr, nil
}
