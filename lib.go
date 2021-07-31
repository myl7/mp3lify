package mp3lify

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os/exec"
)

var ffmpegOptions = []string{"-vn", "-b:a", "192k", "-f", "mp3", "-"}

func ConvertAudio(input string, w io.Writer) error {
	options := []string{"-i", input}
	options = append(options, ffmpegOptions...)
	cmd := exec.Command("ffmpeg", options...)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	buf := bufio.NewReader(out)
	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func Listen(addr string) error {
	http.HandleFunc("/", handle)
	return http.ListenAndServe(addr, nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	src := q.Get("src")
	if src == "" {
		http.Error(w, "Field src not found", http.StatusBadRequest)
		return
	}

	src, err := CleanUrl(src)
	if err != nil {
		http.Error(w, "Field src is invalid URL", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := ConvertAudio(src, w); err != nil {
		log.Println(err)
		return
	}
}
