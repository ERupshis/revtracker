// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package data

import (
	json "encoding/json"
	models "github.com/erupshis/revtracker.git/db/models"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson794297d0DecodeGithubComErupshisRevtrackerGitInternalData(in *jlexer.Lexer, out *FrontMessage) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Data":
			(out.Data).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson794297d0EncodeGithubComErupshisRevtrackerGitInternalData(out *jwriter.Writer, in FrontMessage) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Data\":"
		out.RawString(prefix[1:])
		(in.Data).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FrontMessage) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComErupshisRevtrackerGitInternalData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FrontMessage) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComErupshisRevtrackerGitInternalData(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FrontMessage) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComErupshisRevtrackerGitInternalData(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FrontMessage) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComErupshisRevtrackerGitInternalData(l, v)
}
func easyjson794297d0DecodeGithubComErupshisRevtrackerGitInternalData1(in *jlexer.Lexer, out *Data) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Homework":
			(out.Homework).UnmarshalEasyJSON(in)
		case "Questions":
			if in.IsNull() {
				in.Skip()
				out.Questions = nil
			} else {
				in.Delim('[')
				if out.Questions == nil {
					if !in.IsDelim(']') {
						out.Questions = make([]models.Question, 0, 2)
					} else {
						out.Questions = []models.Question{}
					}
				} else {
					out.Questions = (out.Questions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 models.Question
					(v1).UnmarshalEasyJSON(in)
					out.Questions = append(out.Questions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson794297d0EncodeGithubComErupshisRevtrackerGitInternalData1(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Homework\":"
		out.RawString(prefix[1:])
		(in.Homework).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Questions\":"
		out.RawString(prefix)
		if in.Questions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Questions {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Data) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComErupshisRevtrackerGitInternalData1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Data) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComErupshisRevtrackerGitInternalData1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Data) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComErupshisRevtrackerGitInternalData1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Data) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComErupshisRevtrackerGitInternalData1(l, v)
}
