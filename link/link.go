package link

import (
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

func GetCompleteLink(currentLink string, targetLink string) (string, error) {

	case1, _ := regexp.MatchString("^/[a-zA-Z0-9_]+.*", targetLink)

	case2, _ := regexp.MatchString("^//[a-zA-Z0-9_]+.*", targetLink)

	case3, _ := regexp.MatchString("^(http|https).*", targetLink)

	case4, _ := regexp.MatchString(`^\.\./.*`, targetLink)

	u, err := url.Parse(currentLink)

	if err != nil {

		return "", err
	}

	switch true {

	case case1:

		return u.Scheme + "://" + u.Host + targetLink, nil

	case case2:

		return u.Scheme + "://" + strings.Replace(targetLink, "//", "", 1), nil

	case case3:

		return targetLink, nil

	case case4:

		return u.Scheme + "://" + strings.Replace(filepath.Join(u.Host+u.Path, targetLink), "\\", "/", -1), nil

	}

	return "", nil
}
