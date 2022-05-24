// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package ratdelivery

import (
	json "encoding/json"
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

func easyjsonB202bacaDecodeTmpEjsn(in *jlexer.Lexer, out *RatingResp) {
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
		case "newrating":
			out.NewMovieRating = string(in.String())
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
func easyjsonB202bacaEncodeTmpEjsn(out *jwriter.Writer, in RatingResp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"newrating\":"
		out.RawString(prefix[1:])
		out.String(string(in.NewMovieRating))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RatingResp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB202bacaEncodeTmpEjsn(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RatingResp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB202bacaDecodeTmpEjsn(l, v)
}
func easyjsonB202bacaDecodeTmpEjsn1(in *jlexer.Lexer, out *RatingReq) {
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
		case "movieId":
			out.MovieId = string(in.String())
		case "userId":
			out.UserId = string(in.String())
		case "rating":
			out.Rating = string(in.String())
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
func easyjsonB202bacaEncodeTmpEjsn1(out *jwriter.Writer, in RatingReq) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"movieId\":"
		out.RawString(prefix[1:])
		out.String(string(in.MovieId))
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.UserId))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.String(string(in.Rating))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RatingReq) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB202bacaEncodeTmpEjsn1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RatingReq) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB202bacaDecodeTmpEjsn1(l, v)
}
