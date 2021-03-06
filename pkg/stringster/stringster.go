package stringster

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func RenameFile(str string) string {
	str = ToLowerWithNoSpace(str)
	str = RemoveAccent(str)
	str = RemoveSpecialChars(str)

	return str
}

func ToLowerWithNoSpace(str string) string {
	return strings.ToLower(strings.Replace(str, " ", "_", -1))
}

func RemoveSpecialChars(str string) string {
	re, err := regexp.Compile(`[^\w/.-]`)
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(str, "")
}

func RemoveAccent(str string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	str, _, _ = transform.String(t, str)
	return str
}

func GetExtensionFromString(str string) string {
	return filepath.Ext(str)
}

func IsExtension(file string, extensions ...string) bool {
	fext := GetExtensionFromString(file)
	for _, ext := range extensions {
		if fext != ext {
			continue
		}
		return true
	}
	return len(extensions) == 0
}
