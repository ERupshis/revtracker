package utils

import (
	"testing"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/stretchr/testify/assert"
)

var (
	testEmpty = ""
	testStr   = "some_str"
)

func TestValidateHomeworkData(t *testing.T) {
	type args struct {
		data *data.HomeworkData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				data: &data.HomeworkData{
					Name: "some",
				},
			},
			wantErr: false,
		},
		{
			name: "empty name",
			args: args{
				data: &data.HomeworkData{
					Name: "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHomeworkData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHomeworkData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateContentData(t *testing.T) {
	type args struct {
		data *data.Content
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				data: &data.Content{
					Task:     &testStr,
					Answer:   &testStr,
					Solution: &testStr,
				},
			},
			wantErr: false,
		},
		{
			name: "task is empty",
			args: args{
				data: &data.Content{
					Task:     &testEmpty,
					Answer:   &testStr,
					Solution: &testStr,
				},
			},
			wantErr: true,
		},
		{
			name: "answer is empty",
			args: args{
				data: &data.Content{
					Task:     &testStr,
					Answer:   &testEmpty,
					Solution: &testStr,
				},
			},
			wantErr: true,
		},
		{
			name: "solution is empty",
			args: args{
				data: &data.Content{
					Task:     &testStr,
					Answer:   &testStr,
					Solution: &testEmpty,
				},
			},
			wantErr: true,
		},
		{
			name: "task is empty",
			args: args{
				data: &data.Content{
					Task:     nil,
					Answer:   &testStr,
					Solution: &testStr,
				},
			},
			wantErr: true,
		},
		{
			name: "answer is empty",
			args: args{
				data: &data.Content{
					Task:     &testStr,
					Answer:   nil,
					Solution: &testStr,
				},
			},
			wantErr: true,
		},
		{
			name: "solution is empty",
			args: args{
				data: &data.Content{
					Task:     &testStr,
					Answer:   &testStr,
					Solution: nil,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateContentData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateContentData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateQuestionsData(t *testing.T) {
	type args struct {
		data []data.Question
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "valid",
			args: args{
				data: []data.Question{
					{
						Name: "some_name",
						Content: data.Content{
							Task:     &testStr,
							Answer:   &testStr,
							Solution: &testStr,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "question doesn't have name",
			args: args{
				data: []data.Question{
					{
						Name: "",
						Content: data.Content{
							Task:     &testStr,
							Answer:   &testStr,
							Solution: &testStr,
						},
					},
				},
			},
			wantErr:    true,
			wantErrMsg: "invalid questions: need to check idx: [0]",
		},
		{
			name: "question doesn't have name",
			args: args{
				data: []data.Question{
					{
						Name: "",
						Content: data.Content{
							Task:     &testStr,
							Answer:   &testStr,
							Solution: &testStr,
						},
					},
					{
						Name: "some_name",
						Content: data.Content{
							Task:     &testEmpty,
							Answer:   &testStr,
							Solution: &testStr,
						},
					},
				},
			},
			wantErr:    true,
			wantErrMsg: "invalid questions: need to check idx: [0 1]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateQuestionsData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateQuestionsData() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				assert.Equal(t, tt.wantErrMsg, err.Error())
			}
		})
	}
}
