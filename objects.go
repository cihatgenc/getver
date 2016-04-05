package main

import (
	"net/http"
)

type versionrow struct {
	Build       string `json:"build"`
	Description string `json:"description"`
	ReleaseDate string `json:"releasedate"`
}

type versionrows []versionrow

type prettyversionrow struct {
	MajorVersion string `json:"majorversion"`
	BuildVersion string `json:"buildversion"`
	BuildNumber  int64  `json:"buildnumber"`
	Description  string `json:"description"`
	ReleaseDate  string `json:"releasedate"`
}

type prettyversionrows []prettyversionrow

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
