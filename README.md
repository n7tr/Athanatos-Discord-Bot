<img src="https://media.discordapp.net/attachments/1160835997291003985/1193870513156669511/IMG_9977.jpg?ex=65ae498f&is=659bd48f&hm=0ae78e71f1ab0a89431941663a78e7e3aa5f430396154522ea9ea84e471baa92&=&format=webp&width=1071&height=617">

# Dynamic Discord Bot
A powerful Discord nuke bot written on Go

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
</pre>
# Installation guide
<pre>
	1. Clone or download the repository source code
	2. Install golang
	3. Go to src folder
	4. Change values in .env
	5. Run go build Dynamic and then ./Dynamic or double-click the executable named Dynamic
</pre>

# Where to host?
We recommend you to use <a href="https://fl0.com">fl0.com</a>, <a href="https://back4app.com">back4app.com</a>, <a href="https://koyeb.com">koyeb.com</a> and <a href="https://render.com">render.com</a>. They're free and there you can host Dynamic and other discord bots. More information about other hostings are <a href="https://github.com/DmitryScaletta/free-heroku-alternatives">here</a>

# Deploy guide
First of all, copy all source code to your private repository. Then create an account on <a href="https://railway.app">railway.app</a> via github. Use Dockerfile for quick deployment. <a href="https://railway.app">Railway.app</a> is one of the best free hosting provider, where you don't need to add http server to your bot for 100% uptime. 

# Dockerfile example
<pre>
# For deployment on railway.app
FROM golang:latest

WORKDIR /

COPY . .

RUN go build Dynamic

CMD [ "./Dynamic" ]
</pre>

<pre>
# For deployment on render.com and others
FROM golang:latest

WORKDIR /

COPY . .

RUN go build Dynamic

EXPOSE 8080

CMD [ "./Dynamic" ]
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
	io.WriteString(w, "Dynamic is at render.com now.. 🚀\n")
}
</pre>

