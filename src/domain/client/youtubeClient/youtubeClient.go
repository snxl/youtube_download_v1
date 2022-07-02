package youtubeClient

import (
	"github.com/snxl/youtube-dowloader/src/domain/entity"
	"os"
)

type YoutubeClient interface {
	GetVideoInfos(id string) ([]entity.Video, error)
	FindByTag(id string, tagId int, file *os.File) error
}
