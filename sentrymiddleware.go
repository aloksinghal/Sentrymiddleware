package sentrymiddleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"errors"
	raven "github.com/getsentry/raven-go"
)

// SentryRecovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.
// Also, reports the error to Sentry on DSN provided as input
type Middleware struct{}

type SentryRecovery struct {
	Logger           *log.Logger
	PrintStack       bool
	StackAll         bool
	StackSize        int              
}

// NewSentryRecovery returns a new instance of Sentry Recovery
func NewSentryRecovery() *SentryRecovery {
	return &SentryRecovery{
		Logger:     log.New(os.Stdout, "[negroni] ", 0),
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
	}
}

func (m Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	defer func() {
		if recovery := recover(); recovery != nil {
			recoveryStr := fmt.Sprint(recovery)
			packet := raven.NewPacket(recoveryStr, raven.NewException(errors.New(recoveryStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(request))
			raven.Capture(packet, nil)
			response.WriteHeader(http.StatusInternalServerError)
		}
	}()
	next(response, request)
}