package csvreader

import (
	"io"
	"log"
)

func Close(c io.Closer, f func(err error)) {
	f(c.Close())
}

func WarnOnError(err error) {
	if err != nil {
		log.Printf("Failed to close; %s", err.Error())
	}
}
