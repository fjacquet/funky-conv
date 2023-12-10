package main

import (
	"os/exec"
	"testing"
)

func encodeVideo(input, output, codec, resolution string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", codec, "-s", resolution, output)
	return cmd.Run()
}

func convertVideo(input, output, format string) error {
	cmd := exec.Command("ffmpeg", "-i", input, output)
	return cmd.Run()
}

func scaleVideo(input, output, codec string, resolution string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-vf", "scale=1280:720", output)
	return cmd.Run()
}

func encodeTo720p(input, output string) error {

	cmd := exec.Command("ffmpeg", "-i", input, "-vf", "scale=1280:720", "-crf", "20", output)
	return cmd.Run()
}

func BenchmarkEncodeVideo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeVideo("test.mp4", "test.mkv", "libx264", "1920x1080")
	}
}

func BatchEncode(videos []Video) {
	for _, video := range videos {
		go encodeVideo(video.Input, video.Output, video.Codec, video.Resolution)
	}
}

func TestEncodeVideo(t *testing.T) {
	err := encodeVideo("test.mp4", "test.mkv", "libx264", "1920x1080")
	if err != nil {
		t.Errorf("Failed to encode video: %v", err)
	}
}

// func TestUploadEndpoint(t *testing.T) {
// 	e := httpexpect.New(t, "http://localhost:8080")
// 	file := e.POST("/upload").WithFile("video", "test.mp4")
// 	file.Expect().Status(http.StatusOK).JSON().Object().Value("message").Equal("Video uploaded successfully")
// }
