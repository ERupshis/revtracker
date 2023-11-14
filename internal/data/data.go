package data

//go:generate reform

//go:generate easyjson -all data.go

//reform:homeworks
type Homework struct {
	ID      int64  `json:"Id" reform:"id,pk"`
	Name    string `json:"Name" reform:"name"`
	Deleted bool   `json:"-" reform:"deleted"`
}

//reform:contents
type Content struct {
	ID       int64   `json:"-" reform:"id,pk"`
	Task     *string `json:"Task" reform:"task"`
	Answer   *string `json:"Answer" reform:"answer"`
	Solution *string `json:"Solution" reform:"solution"`
	Deleted  bool    `json:"-" reform:"deleted"`
}

//reform:questions
type Question struct {
	ID        int64   `json:"Id" reform:"id,pk"`
	Name      string  `json:"Name" reform:"name"`
	ContentID int64   `json:"-" reform:"content_id"`
	Content   Content `json:"Content"`
	Deleted   bool    `json:"-" reform:"deleted"`
}

//reform:homework_questions
type HomeworkQuestion struct {
	ID         int64 `json:"Id" reform:"id,pk"`
	HomeworkID int64 `json:"Homework_Id"  reform:"homework_id"`
	QuestionID int64 `json:"Question_Id" reform:"question_id"`
	Order      int64 `json:"Order" reform:"order"`
	Deleted    bool  `json:"-" reform:"deleted"`
}

type Data struct {
	Homework HomeworkData `json:"Homework"`
}

type HomeworkData struct {
	ID        int64      `json:"Id"`
	Name      string     `json:"Name"`
	Questions []Question `json:"Questions"`
}
