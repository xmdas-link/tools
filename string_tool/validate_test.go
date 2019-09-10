package string_tool

import (
	"regexp"
	"testing"
)

func TestOnlyNumberWithLowercase(t *testing.T) {

	if !OnlyNumberWithLowercase("1231abaer") {
		t.Errorf("[1231abaer] should pass")
	}

	if !OnlyNumberWithLowercase("1231231") {
		t.Errorf("[1231231] should pass")
	}

	if !OnlyNumberWithLowercase("ab1231231") {
		t.Errorf("[ab1231231] should pass")
	}

	if OnlyNumberWithLowercase("!ad1231") {
		t.Errorf("[!ad1231] should not pass")
	}

	if OnlyNumberWithLowercase("!ad1231&") {
		t.Errorf("[!ad1231&] should not pass")
	}

	if OnlyNumberWithLowercase(" ad1231") {
		t.Errorf("[ ad1231] should not pass")
	}

	if OnlyNumberWithLowercase("AB1231") {
		t.Errorf("[AB1231] should not pass")
	}

	if OnlyNumberWithLowercase("") {
		t.Errorf("[] should not pass")
	}

}

func TestUpdate(t *testing.T) {
	reg := regexp.MustCompile(`^update_[\w]+`)
	shouldPass := []string{
		"update_123123.s1l",
		"update_adfadf",
	}

	for _, s := range shouldPass {
		if !reg.MatchString(s) {
			t.Errorf("[%s] should pass", s)
		}
	}

	shouldNotPass := []string{
		"update_",
		"1update_afdadfa",
	}

	for _, s := range shouldNotPass {
		if reg.MatchString(s) {
			t.Errorf("[%s] should not pass", s)
		}
	}

}
