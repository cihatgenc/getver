package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	timeFormat     = "2006-01-02 15:04:05"
	buildnrurl     = "https://buildnumbers.wordpress.com/sqlserver/"
	builnrfilename = "mssqldb.html"
)

func Crawler() (fullist versionrows, err error) {

	doc, err := getDocument()
	if err != nil {
		log.Fatal(err)
	}

	var fetchedversions versionrows

	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		// s here is a tbody element
		if i < 1 {
			return // This is the first table body in the html which we dont want
		}
		s.Find("tr").Each(func(j int, s2 *goquery.Selection) {
			var fetchedversion versionrow
			s2.Find("td").Each(func(k int, s3 *goquery.Selection) {
				switch {
				case k == 0:
					fetchedversion.Build = s3.Text()
				case k == 1:
					fetchedversion.Description = s3.Text()
				case k == 2:
					fetchedversion.ReleaseDate = s3.Text()
				}
			})

			fetchedversions = append(fetchedversions, fetchedversion)
		})
	})

	return fetchedversions, nil
}

func CrawlerMajor(majorversion string) (versionlist prettyversionrows, err error) {
	allversions, err := Crawler()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var result prettyversionrows

	for _, allversionsrow := range allversions {
		var resultrow prettyversionrow

		resultrow.BuildVersion = allversionsrow.Build
		resultrow.Description = allversionsrow.Description

		rp := regexp.MustCompile("[0-9]+.[0-9]+")
		match := rp.FindString(allversionsrow.Build)
		resultrow.MajorVersion = match

		bldnr := strings.Replace(allversionsrow.Build, ".", "", -1)

		resultrow.BuildNumber, _ = strconv.ParseInt(bldnr, 10, 64)

		corrected := strings.Replace(allversionsrow.ReleaseDate, "Janury", "January", -1)
		layout := "2006, January 2"
		pdt, err := time.Parse(layout, corrected)
		if err != nil {
			resultrow.ReleaseDate = ""
		} else {
			resultrow.ReleaseDate = fmt.Sprintf("%s", pdt.Format(timeFormat))
		}

		if resultrow.MajorVersion == GetVersionByName(strings.ToUpper(majorversion)) {
			result = append(result, resultrow)
		}

	}

	return result, nil
}

func CrawlerMajorLatest(majorversion string) (latestversion prettyversionrows, err error) {
	majorversions, err := CrawlerMajor(majorversion)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	sort.Sort(sort.Reverse(prettyversionrows(majorversions)))

	result := majorversions[:1]
	return result, nil
}

func (a prettyversionrows) Len() int           { return len(a) }
func (a prettyversionrows) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a prettyversionrows) Less(i, j int) bool { return a[i].ReleaseDate < a[j].ReleaseDate }

func GetVersionByName(versionname string) (versionnumber string) {
	switch versionname {
	case "2016":
		return "13.0"
	case "2014":
		return "12.0"
	case "2012":
		return "11.0"
	case "2008R2":
		return "10.50"
	case "2008":
		return "10.00"
	case "2005":
		return "9.00"
	case "2000":
		return "8.00"
	default:
		return ""
	}
}

func getDocument() (document *goquery.Document, err error) {
	fqname := fmt.Sprintf("%s/%s", os.TempDir(), builnrfilename)
	var fi *os.File

	fileexist := fileExists(fqname)

	expired, err := isFileExpired(fqname, 4)

	if expired || !fileexist {
		err = downloadDocument(buildnrurl)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	fi, err = os.Open(fqname)
	if err != nil {
		log.Fatal(err)
	}

	defer fi.Close()

	doc, err := goquery.NewDocumentFromReader(fi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return doc, nil
}

func downloadDocument(url string) (err error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = writeFile(builnrfilename, contents)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func readFile(filename string) (reader io.Reader, err error) {
	b, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return b, nil
}

func writeFile(filename string, content []byte) (fqname string, err error) {
	fqname = fmt.Sprintf("%s/%s", os.TempDir(), filename)

	err = ioutil.WriteFile(fqname, content, 0644)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return fqname, nil
}

func isFileExpired(fqname string, expiry int) (result bool, err error) {
	fi, err := os.Stat(fqname)
	if err != nil {
		return true, err
	}

	filedatetime := fmt.Sprintf("%s", fi.ModTime().Format(timeFormat))
	currentdatetime := fmt.Sprintf("%s", time.Now().Format(timeFormat))

	a, err := time.Parse(timeFormat, filedatetime)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	b, err := time.Parse(timeFormat, currentdatetime)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	delta := b.Sub(a)

	if int(delta.Hours()) > expiry {
		return true, nil
	}

	return false, nil
}

func fileExists(fqname string) (result bool) {
	if _, err := os.Stat(fqname); os.IsNotExist(err) {
		return false
	}

	return true
}
