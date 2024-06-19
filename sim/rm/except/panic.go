package except

import (
  "os"
  "os/signal"
)

var Error error chan

func PanicHandler(){
  ch := make(chan os.Signal, 1)
  Error = make(chan error, 1)
}
