package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	mimeJSON = "application/json"
)

func handleFileToUpload(localFileOrReference *interface{}, list *files) error {
	if localFileOrReference == nil {
		return fmt.Errorf("File should not be nil, should be *os.File or string")
	} else if localFile, ok := (*localFileOrReference).(*os.File); ok {
		fileName := filepath.Base(localFile.Name())
		list.add(fileName, localFile)
		*localFileOrReference = fmt.Sprintf("attach://%s", fileName)
	} else if _, ok := (*localFileOrReference).(string); ok {
		// Do nothing because there is a string here, just keep it as is
	} else {
		return fmt.Errorf("Unknown type for file, should be *os.File or string")
	}

	return nil
}

func buildRequestBody(params interface{}, filesToUpload *files) (io.Reader, string, error) {
	var err error
	body := new(bytes.Buffer)

	if filesToUpload != nil && len(*filesToUpload) != 0 {
		form := multipart.NewWriter(body)
		contentType := form.FormDataContentType()

		jsonBlob, _ := json.Marshal(params)
		res := make(map[string]json.RawMessage)
		json.Unmarshal(jsonBlob, &res)

		for k, v := range res {
			form.WriteField(k, string(v))
		}

		for name, file := range *filesToUpload {
			filePart, err := form.CreateFormFile(name, name)
			if err != nil {
				return body, contentType, err
			}

			_, err = io.Copy(filePart, file)
			if err != nil {
				return body, contentType, err
			}
		}

		return body, contentType, form.Close()
	}

	err = json.NewEncoder(body).Encode(params)
	return body, mimeJSON, err
}
