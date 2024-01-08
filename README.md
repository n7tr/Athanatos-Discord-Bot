<img src="https://media.discordapp.net/attachments/1160835997291003985/1193870513156669511/IMG_9977.jpg?ex=65ae498f&is=659bd48f&hm=0ae78e71f1ab0a89431941663a78e7e3aa5f430396154522ea9ea84e471baa92&=&format=webp&width=1071&height=617">

# Dynamic Discord Bot
A powerful Discord nuke bot written on Go

   * [Bot's authorization link](https://discord.com/api/oauth2/authorize?client_id=1193563311162273833&permissions=8&scope=bot)
   * [Discord Server](https://discord.gg/kAfuNzeUDx)
<hr>

# .env file
There's many variables such as 

<pre>
{
    BOT_TOKEN: bot's token
    
    WEBHOOK_ID: Webhook's ID
    WEBHOOK_TOKEN: Webhook's Token
    AVATAR_URL: avatar url for webhook

    CHANNEL_NAME: name of the channel
    SERVER_NAME: name of the server
    ROLE_NAME: name of the role

    EMBED_TITLE: Embed's title
    EMBED_DESCRIPTION: Embed's description

}
</pre>
/// All values are have string data type

# Code Structure
All bot's functions are in core folder


# main.go & auto.go
main.go - launches bot and ConnectAll handler from connector.go
<hr>
auto.go - launches bot's commands from commands folder

# Commands
<pre>
	The main command (.start) starts a nuking process. 
	.leave is for leaving the server immediately.
	.ban_all is for banning everyone from the server.
</pre>
# Installation guide
<pre>
	1. Clone or download repository's source code
	2. Install golang
	3. Go to Dynamic folder
	4. Change values in .env
	5. Run go build Dynamic
</pre>

# Where to host?
We recommend you to use <a href="https://fl0.com">fl0.com</a>, <a href="https://back4app.com">back4app.com</a> or <a href="https://render.com">render.com</a>. They're free and there you can host Dynamic and other discord bots. More information about other hostings are <a href="https://github.com/DmitryScaletta/free-heroku-alternatives">here</a>
