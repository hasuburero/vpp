package except

import (
  "os"
  "os/signal"
)

var Error error chan

func PanicHandler(){
  ch := make(chan os.Signal, 1)
  Error = make(chan error, 1)

  signal.Notify(ch, os.Interrupt)
  go func(){
    for sig := range ch{
      fmt.Println()
    }
  }
}
