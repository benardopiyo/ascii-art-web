package banner

func Shadow() ([]string, error) {
	return readBannerFile("shadow.txt")
}
