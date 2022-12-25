package proto

import "fmt"

// Logger is a simplified logger for making the code and output readable.
type Logger struct {
	debug bool
}

func (l Logger) Debug(in string) {
	if l.debug {
		fmt.Printf("DEBUG: %s\n", in)
	}
}

func (l Logger) Debugf(in string, args ...any) {
	l.Debug(fmt.Sprintf(in, args...))
}

func (l Logger) Error(in string) {
	fmt.Printf("ERROR: %s\n", in)
}

func (l Logger) Errorf(in string, args ...any) {
	l.Error(fmt.Sprintf(in, args...))
}

func (l Logger) Info(in string) {
	fmt.Printf("INFO: %s\n", in)
}

func (l Logger) Infof(in string, args ...any) {
	l.Info(fmt.Sprintf(in, args...))
}
