package cmd

import "regexp"

var headerRegexStr = `([0-9][0-9]:[0-9][0-9]:[0-9][0-9].[0-9][0-9][0-9]) +(\w+): +([^ ]+)`
var reHeader *regexp.Regexp

var contextRegexStr = ` *=> (.+)`
var reContext *regexp.Regexp

var metaRegexStr = `(\w+):([^, ]+)`
var reMeta *regexp.Regexp

func init() {
	var err error
	reHeader, err = regexp.Compile(headerRegexStr)
	if err != nil {
		panic(err)
	}

	reContext, err = regexp.Compile(contextRegexStr)
	if err != nil {
		panic(err)
	}

	reMeta, err = regexp.Compile(metaRegexStr)
	if err != nil {
		panic(err)
	}
}
