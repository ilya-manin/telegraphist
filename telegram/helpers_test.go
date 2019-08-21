package telegram

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHandleFileToUpload(t *testing.T) {
	t.Run("with nil", func(t *testing.T) {
		list := make(files)

		err := handleFileToUpload(nil, &list)

		if err == nil {
			t.Errorf("should return an error")

			expected := "File should not be nil, should be *os.File or string"
			if err.Error() != expected {
				t.Errorf("should return proper error, \"%s\" != \"%s\"", err.Error(), expected)
			}
		}

		if len(list) != 0 {
			t.Errorf("should not add a file to list")
		}
	})

	t.Run("with a local file", func(t *testing.T) {
		tempFile, _ := ioutil.TempFile("", "file_to_upload")
		filePath := tempFile.Name()
		defer os.Remove(filePath)

		var file interface{} = tempFile
		list := make(files)

		err := handleFileToUpload(&file, &list)

		if err != nil {
			t.Errorf("should not return an error")
		}

		if len(list) == 0 {
			t.Errorf("should add a file to list")
		}

		fileName := filepath.Base(filePath)
		if file != fmt.Sprintf("attach://%s", fileName) {
			t.Errorf("should replace the file with filename, \"%s\" != \"%s\"", file, fileName)
		}
	})

	t.Run("with a reference to file", func(t *testing.T) {
		var file interface{} = "foo_bar_baz"
		list := make(files)

		err := handleFileToUpload(&file, &list)

		if err != nil {
			t.Errorf("should not return an error")
		}

		if len(list) != 0 {
			t.Errorf("should not add a file to list")
		}

		if file != "foo_bar_baz" {
			t.Errorf("should not replace the file reference")
		}
	})

	t.Run("with unknown type", func(t *testing.T) {
		var file interface{} = 123
		list := make(files)

		err := handleFileToUpload(&file, &list)

		if err == nil {
			t.Errorf("should return an error")

			expected := "Unknown type for file, should be *os.File or string"
			if err.Error() != expected {
				t.Errorf("should return proper error, \"%s\" != \"%s\"", err.Error(), expected)
			}
		}

		if len(list) != 0 {
			t.Errorf("should not add a file to list")
		}
	})
}

func TestBuildRequestBody(t *testing.T) {
	t.Run("with nil file list", func(t *testing.T) {
		params := struct {
			Foo string `json:"foo"`
			Bar int    `json:"bar"`
			baz bool
		}{
			Foo: "1",
			Bar: 2,
			baz: true,
		}

		body, contentType, err := buildRequestBody(params, nil)

		if err != nil {
			t.Errorf("should not return an error")
		}

		expected := "application/json"
		if contentType != contentType {
			t.Errorf("should return proper content type, \"%s\" != \"%s\"", contentType, expected)
		}

		expectedJSON := []byte("{\"foo\":\"1\",\"bar\":2}\n")
		gotJSON, _ := ioutil.ReadAll(body)

		if !bytes.Equal(gotJSON, expectedJSON) {
			t.Errorf("%s != %s", gotJSON, expectedJSON)
		}
	})

	t.Run("with empty file list", func(t *testing.T) {
		params := struct {
			Foo string `json:"foo"`
			Bar int    `json:"bar"`
			baz bool
		}{
			Foo: "1",
			Bar: 2,
			baz: true,
		}

		list := make(files)

		body, contentType, err := buildRequestBody(params, &list)

		if err != nil {
			t.Errorf("should not return an error")
		}

		expectedContentType := "application/json"
		if contentType != expectedContentType {
			t.Errorf("should return proper content type, \"%s\" != \"%s\"", contentType, expectedContentType)
		}

		expectedJSON := []byte("{\"foo\":\"1\",\"bar\":2}\n")
		gotJSON, _ := ioutil.ReadAll(body)

		if !bytes.Equal(gotJSON, expectedJSON) {
			t.Errorf("%s != %s", gotJSON, expectedJSON)
		}
	})

	t.Run("with files", func(t *testing.T) {
		params := struct {
			Foo bool `json:"foo"`
			Bar int  `json:"bar"`
		}{
			Foo: true,
			Bar: 2,
		}

		list := make(files)
		tempFile, _ := ioutil.TempFile("", "file_to_upload")
		defer os.Remove(tempFile.Name())
		list.add("file_1", tempFile)

		body, contentType, err := buildRequestBody(params, &list)

		if err != nil {
			t.Errorf("should not return an error")
		}

		expectedContentType := "multipart/form-data"
		if !strings.Contains(contentType, expectedContentType) {
			t.Errorf("should return proper content type, \"%s\" != \"%s\"", contentType, expectedContentType)
		}

		gotJSON, _ := ioutil.ReadAll(body)

		expectedData := []string{
			"Content-Disposition: form-data; name=\"foo\"\r\n\r\ntrue",
			"Content-Disposition: form-data; name=\"bar\"\r\n\r\n2",
			"Content-Disposition: form-data; name=\"file_1\"; filename=\"file_1\"\r\nContent-Type: application/octet-stream\r\n",
		}

		for _, expectedItem := range expectedData {
			if !strings.Contains(string(gotJSON), expectedItem) {
				t.Errorf("should contain all data, but \"%s\" is absent", expectedItem)
			}
		}
	})
}
