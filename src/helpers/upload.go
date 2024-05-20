package helpers

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
	uploadDir := "src/uploads"
	epochTime := time.Now().Unix()
	filePath := filepath.Join(uploadDir, fmt.Sprintf("%d_%s", epochTime, file.Filename))

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
			return "", err
		}
	}
	return filePath, nil
}

// func ProcessFileUpload(file *multipart.FileHeader, maxFileSize int64, validFileTypes []string) (string, error) {
// 	if err := SizeUploadValidation(file.Size, maxFileSize); err != nil {
// 		return "", fmt.Errorf("file is too large: %w", err)
// 	}

// 	fileHeader, err := file.Open()
// 	if err != nil {
// 		return "", fmt.Errorf("failed to open file: %w", err)
// 	}
// 	defer fileHeader.Close()

// 	buffer := make([]byte, 512)
// 	if _, err := fileHeader.Read(buffer); err != nil {
// 		return "", fmt.Errorf("failed to read file: %w", err)
// 	}

// 	if err := TypeUploadValidation(buffer, validFileTypes); err != nil {
// 		return "", fmt.Errorf("type of file is invalid: %w", err)
// 	}

// 	uploadResult, err := services.UploadCloudinary(c, file)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to save file: %w", err)
// 	}

// 	return uploadResult.URL, nil
// }
