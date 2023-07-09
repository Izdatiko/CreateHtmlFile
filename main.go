package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func findCSSFile(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".css" {
			return filepath.Join(dir, file.Name()), nil
		}
	}

	return "", nil
}

func CreateHtmlFile(nameOfFile string, titleInHtml string, href string) (*os.File, error) {
	var (
		HtmlFile *os.File
		err      error
	)
	HtmlFile, err = os.Create(nameOfFile + ".html")
	if err != nil {
		return nil, err
	}
	write := func(text string) {
		HtmlFile.WriteString(text)
	}
	write("<html lang=\"en\">\n")
	write("<head>\n")
	write("\t<meta charset=\"UTF-8\">\n")
	write("\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n")
	write("\t<link rel=\"stylesheet\" href=\"" + href + "\">\n")
	write("\t<title>" + titleInHtml + "</title>\n")
	write("</head>\n")
	write("<body>\n\n")
	write("</body>\n")
	write("</html>")

	return HtmlFile, nil
}

func main() {
	dir := "./" // путь к директории
	cssFile, err := findCSSFile(dir)
	if err != nil {
		fmt.Printf("Ошибка при поиске .css файла: %s\n", err)
		return
	}

	if cssFile == "" {
		fmt.Println("В данной директории нет .css файла.")
	}

	file, err := CreateHtmlFile("index", "index", cssFile)
	if err != nil {
		fmt.Printf("Ошибка при создании HTML-файла: %s\n", err)
		return
	}
	file.Close()
}
