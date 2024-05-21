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

// func ProcessFileUpload(c *fiber.Ctx, file *multipart.FileHeader) interface{} {
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":     "bad request",
// 			"statusCode": 400,
// 			"message":    "Failed to upload file",
// 		})
// 	}

// 	if err := SizeUploadValidation(file.Size, 2<<20); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":     "bad request",
// 			"statusCode": 400,
// 			"message":    "File is too large",
// 		})
// 	}

// 	fileHeader, err := file.Open()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"status":     "server error",
// 			"statusCode": 500,
// 			"message":    "Failed to open file",
// 		})
// 	}
// 	defer fileHeader.Close()

// 	buffer := make([]byte, 512)
// 	if _, err := fileHeader.Read(buffer); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"status":     "server error",
// 			"statusCode": 500,
// 			"message":    "Failed to read file",
// 		})
// 	}

// 	validFileTypes := []string{"image/png", "image/jpeg", "image/jpg"}
// 	if err := TypeUploadValidation(buffer, validFileTypes); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":     "bad request",
// 			"statusCode": 400,
// 			"message":    "Type of file is invalid",
// 		})
// 	}

// 	uploadResult, err := services.UploadCloudinary(c, file)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"status":     "server error",
// 			"statusCode": 500,
// 			"message":    "Failed to save file",
// 		})
// 	}

// 	return uploadResult.URL
// }
