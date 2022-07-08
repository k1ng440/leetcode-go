package client

type GetQuestionDataResponse struct {
	Data struct {
		Question struct {
			QuestionID       string      `json:"questionId"`
			Title            string      `json:"title"`
			TitleSlug        string      `json:"titleSlug"`
			Content          string      `json:"content"`
			Difficulty       string      `json:"difficulty"`
			ExampleTestcases string      `json:"exampleTestcases"`
			CategoryTitle    string      `json:"categoryTitle"`
			CompanyTagStats  interface{} `json:"companyTagStats"`
			CodeSnippets     []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
				Typename string `json:"__typename"`
			} `json:"codeSnippets"`
			Status            string        `json:"status"`
			SampleTestCase    string        `json:"sampleTestCase"`
			MetaData          string        `json:"metaData"`
			JudgerAvailable   bool          `json:"judgerAvailable"`
			JudgeType         string        `json:"judgeType"`
			MysqlSchemas      []interface{} `json:"mysqlSchemas"`
			EnableRunCode     bool          `json:"enableRunCode"`
			EnableTestMode    bool          `json:"enableTestMode"`
			LibraryURL        interface{}   `json:"libraryUrl"`
			ChallengeQuestion struct {
				ID                       string `json:"id"`
				Date                     string `json:"date"`
				IncompleteChallengeCount int    `json:"incompleteChallengeCount"`
				StreakCount              int    `json:"streakCount"`
				Type                     string `json:"type"`
				Typename                 string `json:"__typename"`
			} `json:"challengeQuestion"`
			Typename string `json:"__typename"`
		} `json:"question"`
	} `json:"data"`
}
