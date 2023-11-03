package data

//go:generate reform

//go:generate easyjson -all data.go

//reform:users
type User struct {
	ID   int64  `json:"Id" reform:"id,pk"`
	Name string `json:"Name" reform:"name"`
}

//reform:homeworks
type Homework struct {
	ID   int64  `json:"Id" reform:"id,pk"`
	Name string `json:"Name" reform:"name"`
}

//reform:contents
type Content struct {
	ID       int64   `json:"Id" reform:"id,pk"`
	Task     *string `json:"Task" reform:"task"`
	Answer   *string `json:"Answer" reform:"answer"`
	Solution *string `json:"Solution" reform:"solution"`
}

//reform:questions
type Question struct {
	ID        int64  `json:"Id" reform:"id,pk"`
	Name      string `json:"Name" reform:"name"`
	ContentID int64  `json:"Content_Id" reform:"content_id"`
}

//reform:homework_questions
type HomeworkQuestion struct {
	ID         int64 `json:"Id" reform:"id,pk"`
	HomeworkID int64 `json:"Homework_Id"  reform:"homework_id"`
	QuestionID int64 `json:"Question_Id" reform:"question_id"`
	Order      int64 `json:"Order" reform:"order"`
}

type Data struct {
	Homework  Homework   `json:"Homework"`
	Questions []Question `json:"Questions"`
}

type FrontMessage struct {
	Data Data `json:"Data"`
}
