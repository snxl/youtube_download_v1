package youtubeVideoInfos

import (
	"github.com/snxl/youtube-dowloader/src/domain/client/youtubeClient"
	"github.com/snxl/youtube-dowloader/src/domain/entity"
)

type YoutubeVideoInfo struct {
	youtubeClient youtubeClient.YoutubeClient
}

func NewYoutubeVideoInfo(youtubeClient youtubeClient.YoutubeClient) *YoutubeVideoInfo {
	return &YoutubeVideoInfo{
		youtubeClient: youtubeClient,
	}
}

func (t *YoutubeVideoInfo) Run(input Input) ([]entity.Video, error) {

	infos, err := t.youtubeClient.GetVideoInfos(input.VideoId)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
