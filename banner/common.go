package banner

import (
	"fmt"
	"os"
	"strings"
)

func Thinkertoy() ([]string, error) {
	return readBannerFile("banner/thinkertoy.txt")
}

func Standard() ([]string, error) {
	return readBannerFile("banner/standard.txt")
}

func Shadow() ([]string, error) {
	return readBannerFile("banner/shadow.txt")
}

func readBannerFile(fileName string) ([]string, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var strArray []string
	if fileName == "thinkertoy.txt" {
		strArray = strings.Split(string(fileContent), "\r\n")
	} else {
		strArray = strings.Split(string(fileContent), "\n")
	}

	if len(strArray) != 856 {
		return nil, fmt.Errorf("error: the file is modified")
	}
	return strArray, nil
}
