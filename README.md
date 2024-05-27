<img src="https://i.imgur.com/fa5DBLL.jpeg">

# Athánatos Discord Bot
A powerful Discord nuke bot written in Go

   * [Bot's authorization link](https://discord.com/api/oauth2/authorize?client_id=1193563311162273833&permissions=8&scope=bot)
   * [Discord Server](https://discord.gg/kAfuNzeUDx)

## Big thanks to [morg](https://github.com/00-Morg-00)
for code improvements

# .env file
There's many variables such as 

<pre>
{
    BOT_TOKEN: bot's token
    BOT_OWNER_ID: your id
	
    MASS_BAN: true or false

    WEBHOOK_URL: Webhook's URL
    AVATAR_URL: avatar url for webhook

    PREFERRED_LOCALE: check list of <a href="https://discord.com/developers/docs/reference#locales">locales</a>

    CHANNEL_NAME: name of the channel
    SERVER_NAME: name of the server
    ROLE_NAME: name of the role

    EMBED_TITLE: Embed's title
    EMBED_DESCRIPTION: Embed's description

}
</pre>
All variables have a string data type. Only the MASS_BAN variable has two possible values - true and false. Write them with a lowercase letter.

# Code Structure
All bot's functions are in core folder


# main.go & connector.go
main.go - starts the bot and ConnectAll handler from connector.go
<hr>
connector.go - runs bot commands from commands folder
<hr>

# sendhttp.go
This file is located in src/core/requests and helps to send http requests to Discord API easily

# smooth.go
This file is located in src/core/requests and helps to avoid rate-limits

# queue.go
This function is responsible for creating a nuke queue on the server

# Commands
<pre>
	The main command (.start) starts the nuking process. 
	.leave is for leave the server immediately.
	.ban_all is for ban everyone from the server.
        .overcharge is for leave every server immediately. It works if your ID is the value of BOT_OWNER_ID variable. In other cases it will not work.
	.bypass - use it when Athanatos is located on the server with Security and other anti nuke bots.
</pre>
# Installation guide
<pre>
	1. Clone or download the repository source code
	2. Install golang
	3. Go to src folder
	4. Change values in .env
	5. Run go build Athanatos and then ./Athanatos or double-click the executable named Athanatos
</pre>

# Bypass Anti Nuke bots
Yes, Athanatos can bypass Anti Nuke bots like Security, Wick and other. Use .bypass command for this.

# Where to host?
We recommend you to use <a href="https://fl0.com">fl0.com</a>, <a href="https://back4app.com">back4app.com</a>, <a href="https://koyeb.com">koyeb.com</a> and <a href="https://render.com">render.com</a>. They're free and there you can host Athanatos and other discord bots. More information about other hostings are <a href="https://github.com/DmitryScaletta/free-heroku-alternatives">here</a>

# Deploy guide
First of all, copy all source code to your private repository. Then create an account on <a href="https://railway.app">railway.app</a> via github. Use Dockerfile for quick deployment. <a href="https://railway.app">Railway.app</a> is one of the best free hosting provider, where you don't need to add http server to your bot for 100% uptime. 

# Dockerfile example
<pre>
# For deployment on railway.app
FROM golang:latest

WORKDIR /

COPY . .

RUN go build Athanatos

CMD [ "./Athanatos" ]
</pre>

<pre>
# For deployment on render.com and others
FROM golang:latest

WORKDIR /

COPY . .

RUN go build Athanatos

EXPOSE 8080

CMD [ "./Athanatos" ]
</pre>

If you want to deploy your fork on render.com, add code snippet bellow to main.go
<pre>
// imports
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

//starts http server
func main() {
	go func() {
		http.HandleFunc("/", getRoot)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Athanatos is at render.com now.. 🚀\n")
}
</pre>
