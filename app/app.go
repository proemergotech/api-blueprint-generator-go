package app

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"path/filepath"
)

const start = "###### `# Generated docs start`"
const end = "###### `# Generated docs end`"

var (
	structRegex   = regexp.MustCompile(`(?m:^\s*type\s+(\S+)\s+struct\s*{(?s:(.*?))}\s*?$)`)
	tags          = "`([^`]*)`"
	fieldRegex    = regexp.MustCompile(`(?m:^\s*(\S+)(?:\s*$|\s*(\S+)\s*(?:$|` + tags + `\s*(?:$|//\s*?(\S.*?)?\s*$))))`)
	typeRegex     = regexp.MustCompile(`^(\*?)((?:\[\])?)(?:(interface{})|(\S+))$`)
	jsonTag       = regexp.MustCompile(`json:"(\w+)`)
	numberPattern = regexp.MustCompile(`^u?(?:int|float)\d*$`)
)

func Run(sourcePathPattern string, targetPath string) {
	sourcePaths, err := filepath.Glob(sourcePathPattern)
	check(err)

	apiMdFile, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE,0666)
	check(err)
	defer apiMdFile.Close()

	origData, err := ioutil.ReadAll(apiMdFile)
	check(err)

	startIndex := bytes.Index(origData, []byte(start))
	endIndex := bytes.LastIndex(origData, []byte(end))

	tail := []byte{}
	if startIndex == -1 || endIndex < startIndex {
		startIndex = len(origData)
	} else {
		tail = origData[endIndex + len(end):]
	}

	err = apiMdFile.Truncate(int64(startIndex))
	check(err)

	apiMdFile.Seek(0, 2)

	apiMdFile.WriteString(start + "\n\n")

	for _, sourcePath := range sourcePaths {
		dat, err := ioutil.ReadFile(sourcePath)
		check(err)

		for _, s := range structRegex.FindAllSubmatch(dat, -1) {
			sName := string(s[1])
			sBody := string(s[2])

			apiMdFile.WriteString("### " + sName + "\n")

			for _, f := range fieldRegex.FindAllStringSubmatch(sBody, -1) {
				fName := f[1]
				fType := f[2]
				fTags := f[3]
				fExample := f[4]
				if len(fType) == 0 {
					apiMdFile.WriteString("+ Include " + fName + "\n")
				} else {
					if len(fTags) > 0 {
						jsonMatch := jsonTag.FindStringSubmatch(fTags)
						if len(jsonMatch) > 0 && len(jsonMatch[1]) > 0 {
							fName = jsonMatch[1]
						}
					}

					typeMatches := typeRegex.FindStringSubmatch(fType)

					fOptional := len(typeMatches[1]) > 0
					fSlice := len(typeMatches[2]) > 0
					fInterface := len(typeMatches[3]) > 0
					fTypeName := typeMatches[4]

					apiMdFile.WriteString("+ " + fName)

					if len(fExample) > 0 {
						apiMdFile.WriteString(": `" + fExample + "`")
					}

					apiMdFile.WriteString(" (")

					if fSlice {
						apiMdFile.WriteString("array[")
					}

					if fInterface {
						apiMdFile.WriteString("enum")
					} else if numberPattern.MatchString(fTypeName) {
						apiMdFile.WriteString("number")
					} else {
						switch fTypeName {
						case "string":
							apiMdFile.WriteString("string")
						case "bool":
							apiMdFile.WriteString("bool")
						default:
							apiMdFile.WriteString(fTypeName)
						}
					}

					if fSlice {
						apiMdFile.WriteString("]")
					}

					if fOptional {
						apiMdFile.WriteString(", optional")
					} else {
						apiMdFile.WriteString(", required")
					}

					apiMdFile.WriteString(")\n")
				}
			}
			log.Print("Written: " + sName)

			apiMdFile.WriteString("\n")
		}
	}

	apiMdFile.WriteString(end)

	apiMdFile.Write(tail)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
