// cmd/futurevolatilesignal/main.go
package main

import (
"flag"
"log"
"os"

"futurevolatilesignal/internal/futurevolatilesignal"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := futurevolatilesignal.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
