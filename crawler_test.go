package main

import (
	"testing"
)

type testversions struct {
	versionname   string
	versionnumber string
}

var versions = []testversions{
	{"", ""},
	{"2050", ""},
	{"2016", "13.0"},
	{"2014", "12.0"},
	{"2012", "11.0"},
	{"2008R2", "10.50"},
	{"2008", "10.00"},
	{"2005", "9.00"},
	{"2000", "8.00"},
}

func TestGetVersionByName(t *testing.T) {
	for _, version := range versions {
		v := GetVersionByName(version.versionname)
		if v != version.versionnumber {
			t.Error(
				"For", version.versionname,
				"expected", version.versionnumber,
				"got", v,
			)
		}
	}
}
