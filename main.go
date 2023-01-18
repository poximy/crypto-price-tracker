package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/poximy/crypto-price-tracker/pages"
)

func main() {
	s := NewServer()
	s.MountMiddleware()
	s.MountHandlers()

	fmt.Println(pages.Green+pages.Bold+"\n    >> CRYPTO <<", pages.Reset)
	fmt.Print(pages.Dim, ">"+pages.Reset)
	fmt.Print(pages.Cyan+pages.Bold, "    port  "+pages.Reset)
	fmt.Print(pages.Cyan+pages.Italic, port())
	fmt.Print("\n\n\n", pages.Reset)
	err := http.ListenAndServe(port(), s.Router)
	if err != nil {
		panic(err)
	}

}

type Server struct {
	Router *chi.Mux
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.RequestLogger(&CustomLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags)}))
	s.Router.Use(middleware.Compress(5, "text/css", "application/javascript"))
}

func (s *Server) MountHandlers() {
	s.Router.Mount("/", pages.NewRoute())
}

func NewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func port() string {
	portNum := os.Getenv("PORT")
	if portNum == "" {
		portNum = "8080" // Default port if not specified
	}
	return ":" + portNum
}

type CustomLogFormatter struct {
	Logger middleware.LoggerInterface
}

func (l *CustomLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	useColor := true
	entry := &customLogEntry{
		CustomLogFormatter: l,
		request:            r,
		buf:                &bytes.Buffer{},
		useColor:           useColor,
	}

	return entry
}

type customLogEntry struct {
	*CustomLogFormatter
	request  *http.Request
	buf      *bytes.Buffer
	useColor bool
}

func (l *customLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	cW(l.buf, l.useColor, pages.Magenta+pages.Bold, "%s ", l.request.Method)

	switch {
	case status < 200:
		cW(l.buf, l.useColor, pages.Blue, "%03d", status)
	case status < 300:
		cW(l.buf, l.useColor, pages.Green, "%03d", status)
	case status < 400:
		cW(l.buf, l.useColor, pages.Cyan, "%03d", status)
	case status < 500:
		cW(l.buf, l.useColor, pages.Yellow, "%03d", status)
	default:
		cW(l.buf, l.useColor, pages.Red, "%03d", status)
	}

	cW(l.buf, l.useColor, pages.Cyan, " %s", l.request.RequestURI)
	cW(l.buf, l.useColor, pages.Dim+pages.Yellow, "  %s ", ByteCountSI(bytes))

	if elapsed < 500*time.Millisecond {
		cW(l.buf, l.useColor, pages.Dim+pages.Green, "%s", elapsed)
	} else if elapsed < 5*time.Second {
		cW(l.buf, l.useColor, pages.Dim+pages.Yellow, "%s", elapsed)
	} else {
		cW(l.buf, l.useColor, pages.Dim+pages.Red, "%s", elapsed)
	}

	l.Logger.Print(l.buf.String())
}

func (l *customLogEntry) Panic(v interface{}, stack []byte) {
	panic(v) //PrintPrettyStack(v)
}

// Use colors and format on writer
func cW(w io.Writer, useColor bool, color string, s string, args ...interface{}) {
	w.Write([]byte(color))
	fmt.Fprintf(w, s, args...)
	w.Write([]byte(pages.Reset))
}

// Format bytes size to B, kB, MB, TB, etc...
// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func ByteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
