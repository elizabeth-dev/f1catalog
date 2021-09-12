package command

type F1TVService interface {
	Authenticate() (*string, *uint, error)
	GetPlaylistURL(contentId string, subToken string) (string, error)
}
