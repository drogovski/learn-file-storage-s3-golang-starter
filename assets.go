package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func getAssetPath(mediaType string) (string, error) {
	ext := mediaTypeToExt(mediaType)

	assetName := make([]byte, 32)
	_, err := rand.Read(assetName)
	if err != nil {
		return "", err
	}

	assetNameBase64 := base64.RawURLEncoding.EncodeToString(assetName)
	return fmt.Sprintf("%s%s", assetNameBase64, ext), nil
}

func (cfg apiConfig) getObjectURL(key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", cfg.s3Bucket, cfg.s3Region, key)
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

func mediaTypeToExt(mediaType string) string {
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return ".bin"
	}
	return "." + parts[1]
}

func getVideoAspectRatioBasedDirectory(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	commandOutput := bytes.Buffer{}
	cmd.Stdout = &commandOutput
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("ffprobe error: %v", err)
	}

	videoInfo := Video{}
	err = json.Unmarshal(commandOutput.Bytes(), &videoInfo)
	if err != nil {
		return "", fmt.Errorf("could not parse ffprobe output: %v", err)
	}

	if len(videoInfo.Streams) == 0 {
		return "", errors.New("no video streams found")
	}

	switch *videoInfo.Streams[0].DisplayAspectRatio {
	case "16:9":
		return "landscape", nil
	case "9:16":
		return "portrait", nil
	default:
		return "other", nil
	}
}

func createFullKey(prefix, key string) string {
	return fmt.Sprintf("%s/%s", prefix, key)
}
