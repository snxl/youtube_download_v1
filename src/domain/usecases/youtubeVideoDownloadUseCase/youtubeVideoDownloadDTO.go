package youtubeVideoDownloadUseCase

type Input struct {
	TagId   int
	VideoId string
}

type Output struct {
	Path string
}
