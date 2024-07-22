package scripts

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hrm1810884/works-hai-backend/config"
)

func ImplGenerateAIDrawing(drawingData map[string][]byte) (data []byte, err error) {
	inputPaths := map[string]string{}

	positions := map[string][]byte{
		"top":    drawingData["top"],
		"bottom": drawingData["bottom"],
		"right":  drawingData["right"],
		"left":   drawingData["left"],
	}

	for direction, posData := range positions {
		if posData != nil {
			filePath := "image/" + direction + ".png"
			err := SaveToLocal(posData, filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to save %s drawing: %w", direction, err)
			}
			inputPaths[direction] = filePath
		}
	}

	outputPath := "image/out.png"

	err = sendHTTPPostRequest(inputPaths, outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP POST request: %w", err)
	}
	data, err = ReadFromLocal(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read output of AI generation: %w", err)
	}
	return data, nil
}

func sendHTTPPostRequest(inputPaths map[string]string, outputPath string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	host := cfg.Python.Host
	endpoint := cfg.Python.Endpoint
	url := "http://" + host + "/" + endpoint

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for direction, path := range inputPaths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", path, err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile(direction, filepath.Base(path))
		if err != nil {
			return fmt.Errorf("failed to create form file for %s: %w", path, err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return fmt.Errorf("failed to copy file %s: %w", path, err)
		}
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy response body to output file: %w", err)
	}

	return nil
}

// func executePythonScript(scriptPath string, inputPaths []string, outputPath string) error {
// 	cmd := exec.Command("python3", append(append([]string{scriptPath}, inputPaths...), outputPath)...) //nolint: gosec
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }

func SaveToLocal(data []byte, filePath string) error {
	err := os.WriteFile(filePath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromLocal(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
