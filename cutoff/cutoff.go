package cutoff
// 用于gin的path截断

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

type PathRegexp struct {
	Path   string
	Regexp *regexp.Regexp
	Method string
}

func GenRouteRegexp(r *gin.Engine) []PathRegexp {
	pathRegexps := make([]PathRegexp, len(r.Routes()))
	for i, routeInfo := range r.Routes() {
		itemList := strings.Split(routeInfo.Path, "/")
		for idx, item := range itemList {
			if strings.HasPrefix(item, ":") {
				itemList[idx] = `\w+`
			}
		}
		pRegexp, _ := regexp.Compile(strings.Join(itemList, "/"))
		pathRegexps[i].Path = routeInfo.Path
		pathRegexps[i].Regexp = pRegexp
		pathRegexps[i].Method = routeInfo.Method
	}
	return pathRegexps
}

func FindMatchPath(method string, uri string, pathRegexps []PathRegexp) string {
	for _, item := range pathRegexps {
		if item.Method == method && item.Regexp.MatchString(uri) {
			return item.Path
		}
	}
	return "unknow"
}

