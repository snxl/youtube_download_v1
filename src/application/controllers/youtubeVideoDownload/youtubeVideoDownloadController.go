package youtubeVideoDownload

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kkdai/youtube/v2"
	"github.com/snxl/youtube-dowloader/src/domain/usecases/deleteFileUseCase"
	"github.com/snxl/youtube-dowloader/src/domain/usecases/youtubeVideoDownloadUseCase"
	"github.com/snxl/youtube-dowloader/src/infra/client/youtubeClient"
	"sync"
)

type YoutubeVideoDownloadController struct {
	client youtube.Client
}

var controllerInstance *YoutubeVideoDownloadController
var lock = &sync.Mutex{}

func NewYoutubeVideoDownloadController() *YoutubeVideoDownloadController {
	if controllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if controllerInstance != nil {
			controllerInstance = &YoutubeVideoDownloadController{
				client: youtube.Client{},
			}
		}
	}

	return controllerInstance
}

func (y *YoutubeVideoDownloadController) Handle(ctx *fiber.Ctx) error {

	videoId := ctx.Params("videoId")
	tagId, _ := ctx.ParamsInt("tagId")

	youtubeVideoDownload := youtubeVideoDownloadUseCase.NewYoutubeVideoDownload(youtubeClient.NewYoutubeClientImpl())
	output, err := youtubeVideoDownload.Run(youtubeVideoDownloadUseCase.Input{
		TagId:   tagId,
		VideoId: videoId,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to proccessing video")
	}

	defer deleteFileUseCase.NewDeleteFileUseCase().Run(deleteFileUseCase.Input{Path: output.Name()})

	return ctx.Download(output.Name())
}
