package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"text/template"
	"time"

	"github.com/gookit/color"
)

type a struct {
	Title string
	Items []items
}

type items struct {
	Title string
	Body  string
}

// FilesWalk recursively walks through files finding ones by extension
func FilesWalk(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// PrintStringArray prints all strings in the array
func PrintStringArray(files []string) {
	for _, str := range files {
		fmt.Println(str)
	}
}

// GetFileName fetches the component name of a given file path
func GetFileName(filePath string) string {
	s := strings.Split(filePath, "/")
	return s[len(s)-1]
}

// GetComponentName gets the component name
func GetComponentName(filePath string) string {
	data, _ := ReadFile(filePath)
	Logger(Concat("Attempting ", filePath))
	status, scriptData := GetScriptData(data)

	Logger(status)
	arrData := strings.Split(scriptData, "\n")
	componentName := ""
	level := 0

	for _, line := range arrData {
		if strings.Contains(line, "{") {
			level++
		} else if strings.Contains(line, "}") {
			level--
		}

		if strings.Contains(line, "name:") && level == 1 {
			componentName = line[len("name:")+1 : len(line)]
			componentName = strings.Trim(componentName, ",\"")
			componentName = strings.Trim(componentName, "'")
		}
	}

	if componentName == "" {
		componentName = GetFileName(filePath)
		componentName = RemoveExtension(componentName)
	}

	Logger("Success.")
	return componentName
}

// GetScriptData gets the script portion of a vue file
func GetScriptData(data string) (string, string) {
	startLine := strings.Index(data, "export default")
	endLine := strings.Index(data, "</script>")

	if startLine == -1 || endLine == -1 {
		return "NO_SCRIPT", ""
	}

	scriptData := data[startLine+len("export default") : endLine]
	return "OK", strings.Replace(scriptData, " ", "", -1)
}

// GetTemplateData gets the template portion of a vue file
func GetTemplateData(data string) (string, string) {
	startLine := strings.Index(data, "<template>")
	endLine := strings.LastIndex(data, "</template>")

	if startLine == -1 || endLine == -1 {
		return "NO_TEMPLATE", ""
	}

	templateData := data[startLine+len("<template>") : endLine]

	return "OK", strings.Replace(templateData, " ", "", -1)
}

// Concat merges two strings together
func Concat(a, b string) string {
	var str strings.Builder

	str.WriteString(a)
	str.WriteString(b)
	return str.String()
}

// RemoveExtension removes any file extensions from a name
func RemoveExtension(name string) string {
	s := strings.Split(name, ".")
	return s[0]
}

const templ = `{{.Title}}{{range .Items}}
{{.Title}}	{{.Body}}{{end}}
`

// PrintResults prints the component counter results
func PrintResults(c map[string]*ComponentStruct) {
	data := a{
		Title: "Results",
		Items: make([]items, len(c)),
	}

	i := 0
	for _, v := range c {
		title := color.Green.Sprintf("%s", v.name)

		if v.template == 0 || v.impt == 0 {
			title = color.Red.Sprintf("%s", v.name)
		}

		data.Items[i].Title = title
		data.Items[i].Body = fmt.Sprintf("|  %d call(s), %d import(s)", v.template, v.impt)
		i++
	}

	t := template.New("template")
	t, _ = t.Parse(templ)
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 8, ' ', 0)

	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
	w.Flush()
}

// GetCurrentTime gets the current time
func GetCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 3:4:5 PM")
}
