package taskrunner

import (
	"testing"
	"log"
	"github.com/pkg/errors"
	"time"
)

func TestRunner(t *testing.T) {


	size := 10
	d := func(dc dataChan) error{
		for i := 0;i<size ;i++  {

			dc <- i
			log.Printf("Dispatcher sent: %v",i)
		}

		return nil

	}

	e := func(dc dataChan) error {
		forloop:
		for {
			select {
			case d := <- dc:
				log.Printf("Executor recevied: %v",d)
			default:
				break forloop
			}

		}
		return errors.New("Executor error")

	}

	runner := NewRunner(size,false,d,e)

	go runner.StartAll()
	time.Sleep(1*time.Second)

}