package scripts

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateAIDrawing(drawingData map[string][]byte) ([]byte, error) {
	inputPaths := []string{} //["top=hoge.png", "bottom=fuga.png"]

	positions := map[string][]byte{
		"top":    drawingData["top"],
		"bottom": drawingData["bottom"],
		"right":  drawingData["right"],
		"left":   drawingData["left"],
	}

	for direction, posData := range positions {
		if posData != nil {
			filepath := "image/" + direction + ".png"
			err := SaveToLocal(posData, filepath)
			if err != nil {
				return nil, fmt.Errorf("failed to download %s drawing: %w", direction, err)
			}
			inputPaths = append(inputPaths, fmt.Sprintf("%s=%s", direction, filepath))
		}
	}

	outputPath := "image/out.png"

	err := executePythonScript("./process_image.py", inputPaths, outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to execute Python script: %w", err)
	}
	data, err := ReadFromLocal(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read output of AI generation: %w", err)
	}
	return data, nil
}

func executePythonScript(scriptPath string, inputPaths []string, outputPath string) error {
	cmd := exec.Command("python3", append(append([]string{scriptPath}, inputPaths...), outputPath)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

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
