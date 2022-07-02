package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/snxl/youtube-dowloader/src/application/controllers/youtubeVideoDownload"
	"github.com/snxl/youtube-dowloader/src/application/controllers/youtubeVideoInfos"
)

func youtubeRoutes(app *fiber.App) {
	app.Get("/video/:videoId/info", youtubeVideoInfos.NewYoutubeVideoInfoController().Handle)
	app.Get("/video/:videoId/download/:tagId", youtubeVideoDownload.NewYoutubeVideoDownloadController().Handle)
}
