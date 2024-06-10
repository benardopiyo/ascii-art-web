package banner

func Standard() ([]string, error) {
	return readBannerFile("standard.txt")
}
