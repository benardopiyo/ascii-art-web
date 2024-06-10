package asciimodifier

import (
	"ascii-art-web/ascii-art/banner"
	"fmt"
	"strings"
)

func GenerateArt(text, style string) (string, error) {
	var strArray []string
	var err error

	switch style {
	case "shadow":
		strArray, err = banner.Shadow()
	case "standard":
		strArray, err = banner.Standard()
	case "thinkertoy":
		strArray, err = banner.Thinkertoy()
	default:
		return "", fmt.Errorf("unknown banner style: %s", style)
	}

	if err != nil {
		return "", err
	}

	words := strings.Split(text, "\\n")
	return standardReader(words, strArray), nil
}

func standardReader(words []string, strArray []string) string {
	var str string
	for _, word := range words {
		if word == "" || word == "\n" {
			str = str + "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range word {
				startPoint := int(((char - 32) * 9) + 1)
				str = str + strArray[startPoint+i]
			}
			str = str + "\n"
		}
	}
	return str
}
