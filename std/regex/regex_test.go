package regex

import (
	"fmt"
	"regexp"
	"testing"
)

const (
	keyword = "Golang"
)

type wrapKeywordEntity struct {
	original  string
	afterWrap string
}

// wrapKeywordWithBacktick 将关键字使用 ` 包起来
func wrapKeywordWithBacktick(originalSql, keyword string) string {
	var replace = "`" + keyword + "`"
	var whitespacePattern = fmt.Sprintf(`(?i)(\s+)(%s)(\s+)`, keyword)
	var parenthesesPattern = fmt.Sprintf(`(?i)(\s*\(\s*)(%s)(\s*\)\s*)`, keyword)
	whitespaceCompile := regexp.MustCompile(whitespacePattern)
	parenthesesCompile := regexp.MustCompile(parenthesesPattern)
	originalSql = whitespaceCompile.ReplaceAllString(originalSql, "$1"+replace+"$3")
	return parenthesesCompile.ReplaceAllString(originalSql, "$1"+replace+"$3")
}

var wrapKeywordEntities []wrapKeywordEntity

func init() {
	wrapKeywordEntities = []wrapKeywordEntity{
		{
			original:  "[golang]this is a (Golang) string!",
			afterWrap: "[golang]this is a (`Golang`) string!",
		},
		{
			original:  "[golang]this is a ( Golang) string!",
			afterWrap: "[golang]this is a ( `Golang`) string!",
		},
		{
			original:  "[golang]this is a (Golang ) string!",
			afterWrap: "[golang]this is a (`Golang` ) string!",
		},
		{
			original:  "this is a Golang string!",
			afterWrap: "this is a `Golang` string!",
		},
		{
			original:  "this is a golang string!",
			afterWrap: "this is a `Golang` string!",
		},
	}
}

func TestWrapKeywordWithBacktick(t *testing.T) {
	for _, entity := range wrapKeywordEntities {
		afterWrap := wrapKeywordWithBacktick(entity.original, keyword)
		if afterWrap != entity.afterWrap {
			t.Errorf("error wrap keyword: before = [%s], after = [%s], expected = [%s]\n", entity.original, afterWrap, entity.afterWrap)
		}
	}
}
