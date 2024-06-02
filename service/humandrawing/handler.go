package humandrawing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadHandler(c echo.Context) error {
	log.Println("Received upload request")

	// 最大ファイルサイズを10MBに制限
	err := c.Request().ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error parsing multipart form: %v", err))
	}

	file, handler, err := c.Request().FormFile("image")
	if err != nil {
		log.Printf("Error retrieving the file: %v", err)
		return c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving the file: %v", err))
	}
	defer file.Close()

	log.Printf("Uploaded File: %s, Size: %d, MIME: %s", handler.Filename, handler.Size, handler.Header.Get("Content-Type"))

	err = SaveImage(file, handler.Filename)
	if err != nil {
		log.Printf("Error processing the image: %v", err)
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error processing the image: %v", err))
	}

	return c.String(http.StatusCreated, fmt.Sprintf("File uploaded and converted successfully: %s", handler.Filename))
}
