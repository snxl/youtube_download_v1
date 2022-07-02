package youtubeVideoDownloadUseCase

import (
	"github.com/snxl/youtube-dowloader/src/domain/client/youtubeClient"
	"io/ioutil"
	"os"
)

type YoutubeVideoDownloadImpl struct {
	youtubeClient youtubeClient.YoutubeClient
}

func NewYoutubeVideoDownload(client youtubeClient.YoutubeClient) *YoutubeVideoDownloadImpl {
	return &YoutubeVideoDownloadImpl{
		youtubeClient: client,
	}
}

func (y *YoutubeVideoDownloadImpl) Run(input Input) (*os.File, error) {

	file, err := ioutil.TempFile("tmp", "video.*.mp4")
	if err != nil {
		return nil, err
	}

	err = y.youtubeClient.FindByTag(input.VideoId, input.TagId, file)
	if err != nil {
		return nil, err
	}

	return file, nil
}
