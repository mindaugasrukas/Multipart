package multipart

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type (
	// Multipart is a MIME multipart file
	Multipart struct {
		Files []string
	}
)

// New creates a new MIME multipart file
func New() *Multipart {
	return &Multipart{}
}

// AddFile adds a file to the multipart file
func (m *Multipart) AddFile(filename string) error {
	m.Files = append(m.Files, filename)
	return nil
}

// Write writes the MIME multipart file to the given file
func (m *Multipart) Write(out io.Writer) error {
	multipartWriter := multipart.NewWriter(out)
	defer multipartWriter.Close()

	// Write MIME Multipart headers
	fmt.Fprintf(out, "Content-Type: %s\n\n", multipartWriter.FormDataContentType())

	for _, file := range m.Files {
		partWriter, err := multipartWriter.CreateFormFile("file", file)
		if err != nil {
			return err
		}

		file, err := os.Open(file)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(partWriter, file)
		if err != nil {
			return err
		}
	}

	return nil
}
