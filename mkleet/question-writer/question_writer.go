package questionwriter

import (
	"fmt"
	"io/ioutil"
	"leetcode/mkleet/client"
	"os"
	"path"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type QuestionWriter struct {
	Path     string
	URL      string
	Slug     string
	Question *client.GetQuestionDataResponse
}

func (q *QuestionWriter) WriteMarkdown() error {
	p := path.Join(q.Path, "README.md")

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(q.Question.Data.Question.Content)
	if err != nil {
		return err
	}

	// add header and url to leetcode problem
	markdown = fmt.Sprintf("# [%s](%s)\n\n%s", q.Question.Data.Question.Title, q.URL, markdown)

	return q.writeFile(p, []byte(markdown))
}

func (q *QuestionWriter) writeFile(p string, data []byte) error {
	q.ensureDir()
	return ioutil.WriteFile(p, data, 0664)
}

func (q *QuestionWriter) ensureDir() error {
	return os.MkdirAll(q.Path, 0755)
}
