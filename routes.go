package main

var getverroutes = Routes{
	Route{
		"GetVersions",
		"GET",
		"/api/getver/v1/GetVersions",
		getVersions,
	},
	Route{
		"GetVersions",
		"GET",
		"/api/getver/v1/GetVersions/{majorversion}",
		getVersionsByMajor,
	},
	Route{
		"GetVersionLatest",
		"GET",
		"/api/getver/v1/GetVersionLatest/{majorversion}",
		getlatestVersionByMajor,
	},
}

var routes = append(getverroutes)
