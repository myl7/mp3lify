package mp3lify

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

var ffmpegOptions = []string{"-i", "-", "-vn", "-b:a", "192k", "-f", "mp3", "-"}

func ConvertAudio(r io.ReadCloser, w io.Writer) error {
	cmd := exec.Command("ffmpeg", ffmpegOptions...)

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	var inWg sync.WaitGroup
	inWg.Add(1)
	go func() {
		defer inWg.Done()
		defer in.Close()

		inBuf := bufio.NewWriter(in)
		_, _ = inBuf.ReadFrom(r)
	}()

	var outWg sync.WaitGroup
	outWg.Add(1)
	go func() {
		defer outWg.Done()

		outBuf := bufio.NewReader(out)
		_, _ = outBuf.WriteTo(w)
	}()

	inWg.Wait()

	if err := cmd.Wait(); err != nil {
		return err
	}

	outWg.Wait()

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

	res, err := http.Get(src)
	if err != nil {
		http.Error(w, "Proxy request failed", http.StatusBadRequest)
		return
	}

	if int(res.StatusCode/100) != 2 {
		http.Error(w, "Proxy request failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := ConvertAudio(res.Body, w); err != nil {
		log.Println(err)
		return
	}
}
