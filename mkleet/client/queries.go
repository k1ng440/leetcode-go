package client

const questionDataQuery = `query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    title
    titleSlug
    content
    isPaidOnly
    difficulty
    exampleTestcases
    categoryTitle
    codeSnippets {
      lang
      langSlug
      code
    }
    sampleTestCase
    judgerAvailable
    judgeType
    __typename
  }
}`
