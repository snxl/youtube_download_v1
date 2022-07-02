package youtubeVideoInfos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kkdai/youtube/v2"
	"github.com/snxl/youtube-dowloader/src/domain/usecases/youtubeVideoInfosUseCase"
	"github.com/snxl/youtube-dowloader/src/infra/client/youtubeClient"
	"sync"
)

type YoutubeVideoInfoController struct {
	client youtube.Client
}

var controllerInstance *YoutubeVideoInfoController
var lock = &sync.Mutex{}

func NewYoutubeVideoInfoController() *YoutubeVideoInfoController {
	if controllerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if controllerInstance == nil {
			controllerInstance = &YoutubeVideoInfoController{
				client: youtube.Client{},
			}
		}
	}

	return controllerInstance
}

func (y *YoutubeVideoInfoController) Handle(ctx *fiber.Ctx) error {

	videoId := ctx.Params("videoId")

	youtubeVideoInfoUseCase := youtubeVideoInfosUseCase.NewYoutubeVideoInfo(youtubeClient.NewYoutubeClientImpl())
	Output, err := youtubeVideoInfoUseCase.Run(youtubeVideoInfosUseCase.Input{VideoId: videoId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed in proccessing")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": Output,
	})
}
