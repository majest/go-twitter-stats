package main 


import (
	t "github.com/majest/go-twitter-stream-api"
	db "github.com/majest/go-user-service/db"
	cfg "github.com/majest/go-user-service/config"
	tweet "github.com/majest/go-twitter-stats/tweet"
	"log"
     "os/signal"
    "syscall"
    "time"
    "os"
)

var running = true

func cleanup() {
	log.Println("Interupted")
	running = false
}

func main() {

	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    go func() {
        <-c
        cleanup()
        os.Exit(1)
    }()


    cfg.Load("./config.json")
	db.InitSession()
	
	ch := make(chan t.Message)
	params := make(map[string]string, 1)

	credentials := t.NewCredentials("XjY7q0CYwRxSBzCpUeRDzQ", "214373359-jn77FNlrKEajR4Gpp9l5msb1KXCGXZ7QeJPtt5TF", "cuseCPmxY4taUEmouOhXIvR7MVSUWdRKjKHvHKgVvOk", "tO5hW1ye3myBnT78DspVbTKWFgadvKeU1EOiV3o5Tg")

	params["track"] = "BT"
	params["language"] = "en"

	go t.TwitterStream(ch, credentials, params)

	for ;running; {
		message := <- ch

		T := tweet.New()
		T.Text = message.Tweet.Text()
		T.UserId = message.Tweet.UserID()
		T.UserName = message.Tweet.UserName()
		T.Save()

        time.Sleep(0.1*1e9);
    }

}