package youtubeClient

import "github.com/snxl/youtube-dowloader/src/domain/entity"

type YoutubeClient interface {
	GetVideoInfos(id string) ([]entity.Video, error)
}
