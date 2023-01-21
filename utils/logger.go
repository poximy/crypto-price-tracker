package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func StartMessage(port string) {
	fmt.Println(Green+Bold+">> CRYPTO <<", Reset)
	fmt.Print(Cyan+Bold, "I'm alive at http://localhost", Reset)
	fmt.Print(Cyan+Italic, port, Reset, "\n")
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
	cW(l.buf, l.useColor, Magenta+Bold, "%s ", l.request.Method)

	switch {
	case status < 200:
		cW(l.buf, l.useColor, Blue, "%03d", status)
	case status < 300:
		cW(l.buf, l.useColor, Green, "%03d", status)
	case status < 400:
		cW(l.buf, l.useColor, Cyan, "%03d", status)
	case status < 500:
		cW(l.buf, l.useColor, Yellow, "%03d", status)
	default:
		cW(l.buf, l.useColor, Red, "%03d", status)
	}

	cW(l.buf, l.useColor, Cyan, " %s", l.request.RequestURI)
	cW(l.buf, l.useColor, Dim+Yellow, "  %s ", byteCountSI(bytes))

	if elapsed < 500*time.Millisecond {
		cW(l.buf, l.useColor, Dim+Green, "%s", elapsed)
	} else if elapsed < 5*time.Second {
		cW(l.buf, l.useColor, Dim+Yellow, "%s", elapsed)
	} else {
		cW(l.buf, l.useColor, Dim+Red, "%s", elapsed)
	}

	l.Logger.Print(l.buf.String())
}

func (l *customLogEntry) Panic(v interface{}, stack []byte) {
	panic(v) // PrintPrettyStack(v)
}

// Use colors and format on writer
func cW(w io.Writer, useColor bool, color string, s string, args ...interface{}) {
	w.Write([]byte(color))
	fmt.Fprintf(w, s, args...)
	w.Write([]byte(Reset))
}

// Format bytes size to B, kB, MB, TB, etc...
// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func byteCountSI(b int) string {
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
