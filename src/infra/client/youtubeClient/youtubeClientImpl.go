package youtubeClient

import (
	"github.com/kkdai/youtube/v2"
	"github.com/snxl/youtube-dowloader/src/domain/entity"
	"io"
	"os"
)

type YoutubeClient struct {
	client youtube.Client
}

func NewYoutubeClientImpl() *YoutubeClient {
	return &YoutubeClient{
		client: youtube.Client{},
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
			Id:           element.ItagNo,
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

func (y *YoutubeClient) FindByTag(id string, tagId int, file *os.File) error {

	video, err := y.client.GetVideo(id)
	if err != nil {
		return err
	}

	format := video.Formats.FindByItag(tagId)
	stream, _, err := y.client.GetStream(video, format)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, stream)
	if err != nil {
		return err
	}

	return nil

}
