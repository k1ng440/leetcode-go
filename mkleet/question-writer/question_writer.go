package questionwriter

import (
	"fmt"
	"io/ioutil"
	"leetcode/mkleet/client"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type QuestionWriter struct {
	Path     string
	URL      string
	Slug     string
	Question *client.GetQuestionDataResponse
}

var (
	regexpStartingNumber = regexp.MustCompile(`^[0-9]+`)
	packageNameReplacer  = strings.NewReplacer(" ", "-", "_", "-", ".", "_")
)

func (q *QuestionWriter) WriteMarkdown() error {
	p := path.Join(q.Path, "README.md")

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(q.Question.Data.Question.Content)
	if err != nil {
		return err
	}

	// add header and url to leetcode problem
	markdown = fmt.Sprintf("# [%s](%s)\n\n%s", q.Question.Data.Question.Title, q.URL, markdown)

	// Hack for https://github.com/JohannesKaufmann/html-to-markdown/issues/47
	replacer := strings.NewReplacer("<strong>", "**", "</strong>", "**")
	markdown = replacer.Replace(markdown)

	return q.writeFile(p, []byte(markdown))
}

func (q *QuestionWriter) WriteBoilerPlate() error {
	codeSnippet := ""
	for _, snippet := range q.Question.Data.Question.CodeSnippets {
		if snippet.Lang == "Go" {
			codeSnippet = snippet.Code
		}
	}

	p := path.Join(q.Path, fmt.Sprintf("%s-%s.go", q.Question.Data.Question.QuestionID, q.Slug))

	codeSnippet = fmt.Sprintf("package %s\n\n%s", q.makePackageName(), codeSnippet)
	return q.writeFile(p, []byte(codeSnippet))
}

func (q *QuestionWriter) makePackageName() string {
	haveFirst := false
	res := make([]rune, 0)

	for _, r := range q.Slug {
		if !haveFirst {
			if unicode.IsLetter(r) {
				res = append(res, unicode.ToLower(r))
				haveFirst = true
			}

			continue
		}

		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			res = append(res, unicode.ToLower(r))
		}
	}

	return string(res)
}

func (q *QuestionWriter) writeFile(p string, data []byte) error {
	q.ensureDir()
	return ioutil.WriteFile(p, data, 0664)
}

func (q *QuestionWriter) ensureDir() error {
	return os.MkdirAll(q.Path, 0755)
}
