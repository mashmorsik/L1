package main

import "fmt"

func main() {
	d2 := &Darwin2{}
	d9 := &Darwin9{}

	fmt.Println(d2.Read("epub", "Alice in wonderland"))
	fmt.Println(d9.ReadTXT("War and Peace"))
}

// Reader интерфейс электронной читалки, которая может обрабатывать только формат TXT
type Reader interface {
	Read(format string, file string) error
}

// AdvancedReader интерфейс электронной читалки, которая может обрабатывать еще формат EPUB
type AdvancedReader interface {
	ReadTXT(file string) error
	ReadEPUB(file string) error
}

// Darwin9 читалка нового поколения, реализует интерфейс AdvancedReader
type Darwin9 struct{}

func (d *Darwin9) ReadTXT(file string) error {
	fmt.Printf("Reading TXT file: %s", file)
	return nil
}

func (d *Darwin9) ReadEPUB(file string) error {
	fmt.Printf("Reading EPUB file: %s", file)
	return nil
}

// Adapter реализует интерфейс Reader, но внутри использует AdvancedReader
type Adapter struct {
	advancedReader AdvancedReader
}

func (a *Adapter) Read(format string, file string) error {
	if format == "epub" {
		return a.advancedReader.ReadEPUB(file)
	} else if format == "txt" {
		return a.advancedReader.ReadTXT(file)
	}
	return fmt.Errorf("%s format not supported", format)
}

// Darwin2 читалка старого поколения, будет реализовать интерфейс адаптера
type Darwin2 struct {
	adapter Reader
}

func (d *Darwin2) Read(format string, file string) error {
	if format == "txt" {
		fmt.Printf("Reading TXT file: %s", file)
		return nil
	} else if format == "epub" {
		d.adapter = &Adapter{&Darwin9{}}
		return d.adapter.Read(format, file)
	}

	return fmt.Errorf("%s format not supported", format)
}
