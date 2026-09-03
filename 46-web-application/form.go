package main

import (
	"fmt"
	"net/url"
	"regexp"
	"unicode/utf8"
)

type formError map[string][]string

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (e formError) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e formError) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}

type Form struct {
	url.Values
	Errors formError
}

func NewForm(data url.Values) *Form {

	return &Form{
		Values: data,
		Errors: formError(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) *Form {

	for _, field := range fields {
		value := f.Get(field)
		if value == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}

	return f
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MaxLength(field string, d int) *Form {
	value := f.Get(field)
	if value == "" {
		return f
	}

	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", d))
	}

	return f
}

func (f *Form) MinLength(field string, d int) *Form {
	value := f.Get(field)
	if value == "" {
		return f
	}

	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", d))
	}

	return f
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) *Form {
	value := f.Get(field)
	if value == "" {
		return f
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
	return f
}
