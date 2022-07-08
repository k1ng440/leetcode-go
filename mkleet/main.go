package main

import (
	"errors"
	"fmt"
	"leetcode/mkleet/client"
	questionwriter "leetcode/mkleet/question-writer"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid numbers of args")
		os.Exit(1)
	}
	workingDir, err := os.Getwd()

	slug, err := validateUrlAndGetSlug(os.Args[1])
	checkError(err)

	c := client.New()
	res, err := c.GetQuestionData(slug)
	checkError(err)

	q := &questionwriter.QuestionWriter{
		URL:      strings.TrimSpace(os.Args[1]),
		Slug:     slug,
		Path:     path.Join(workingDir, fmt.Sprintf("%s-%s", res.Data.Question.QuestionID, slug)),
		Question: res,
	}

	q.WriteMarkdown()
	q.WriteBoilerPlate()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateUrlAndGetSlug(str string) (string, error) {
	re := regexp.MustCompile(`\/problems\/(.*)\/`)

	str = strings.TrimSpace(str)
	parsedUrl, err := url.Parse(str)
	if err != nil {
		return "", err
	}

	if parsedUrl.Hostname() != "leetcode.com" && parsedUrl.Hostname() != "leetcode.cn" {
		return "", errors.New("invalid url provided. expected from leetcode.com or leetcode.cn TLD")
	}

	r := re.FindStringSubmatch(parsedUrl.Path)
	if len(r) != 2 {
		return "", errors.New("failed to parse slug from url")
	}

	return r[1], nil
}
