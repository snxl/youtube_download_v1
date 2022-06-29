package youtubeVideoInfos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kkdai/youtube/v2"
	"github.com/snxl/youtube-dowloader/src/domain/usecases/youtubeVideoInfos"
	"github.com/snxl/youtube-dowloader/src/infra/client/youtubeClient"
)

type YoutubeVideoInfoController struct {
	client youtube.Client
}

func NewYoutubeVideoInfoController() *YoutubeVideoInfoController {
	return &YoutubeVideoInfoController{
		client: youtube.Client{},
	}
}

func (y *YoutubeVideoInfoController) Handle(ctx *fiber.Ctx) error {

	videoId := ctx.Params("videoId")

	youtubeVideoInfoUseCase := youtubeVideoInfos.NewYoutubeVideoInfo(youtubeClient.NewYoutubeClientImpl(y.client))
	Output, err := youtubeVideoInfoUseCase.Run(youtubeVideoInfos.Input{VideoId: videoId})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed in proccessing")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": Output,
	})
}
