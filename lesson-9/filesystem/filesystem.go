package filesystem

import "fmt"

type FileHandler interface {
	Open(string) string
	Read() string
	Close() string
}

type TextFile struct {
	path string
}

type CSVFile struct {
	path string
}

type LogFile struct {
	path string
}

func (t *TextFile) Open(path string) string {
	fmt.Printf("Text file is opening in %s \n", path)
	return "Opened"
}
func (t *TextFile) Read() string {
	fmt.Printf("Text file is reading in %s \n", t.path)
	return "The body of txt file"
}
func (t *TextFile) Close() string {
	fmt.Printf("Text file is closing... \n")
	return "Closed"
}

func (c *CSVFile) Open(path string) string {
	fmt.Printf("CSV file is opening in %s \n", path)
	return "Opened"
}
func (c *CSVFile) Read() string {
	fmt.Printf("CSV file is reading in %s \n", c.path)
	return "The body of CSV file"
}
func (c *CSVFile) Close() string {
	fmt.Printf("CSV file is closing... \n")
	return "Closed"
}

func (l *LogFile) Open(path string) string {
	fmt.Printf("LOG file is opening in %s \n", path)
	return "Opened"
}
func (l *LogFile) Read() string {
	fmt.Printf("LOG file is reading in %s \n", l.path)
	return "The body of LOG file"
}
func (l *LogFile) Close() string {
	fmt.Printf("LOG file is closing... \n")
	return "Closed"
}

func ReadAllTypeOfFiles(f []FileHandler) {
	for _, v := range f {
		v.Read()
	}
}
func CloseAllTypeOfFiles(f []FileHandler) {
	for _, v := range f {
		v.Close()
	}
}
