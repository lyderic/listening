package main

import(
  "flag"
  "fmt"
  "log"
  "net"
  "os"
  "strconv"
)

func init() {
  log.SetFlags(log.Lshortfile)
}

func main() {

  flag.Usage = usage
  flag.Parse()
  args := flag.Args()
  if len(args) < 1 {
    usage()
  }

  port,err := strconv.ParseUint(args[0], 10, 16)
  if err != nil {
    log.Fatal(err)
  }

  user := os.Getenv("USER")
  if user != "root" && port < 1024 {
    log.Fatal("Only root can look up ports < 1024!")
  }

  available := false

  ln,err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  defer ln.Close()
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("Port %d is available.\n", port)
    available = true
  }

  if available {
    os.Exit(1)
  } else {
    os.Exit(0)
  }
}

func usage() {
  fmt.Println("Usage: listening <port>")
  flag.PrintDefaults()
  os.Exit(23)
}
