package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)
var (
	stopChannel chan bool
	conf *config
)
func main(){
	conf = parse_config()
	dg, err := discordgo.New("Bot "+ conf.Token)
	if err!=nil {
		fmt.Println("error !gg!", err)
		return
	}
	stopChannel = make(chan bool)
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err!=nil {
		fmt.Print("open Web_socket Error gg!!", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	go queue_way()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	dg.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if (m.Author.ID==s.State.User.ID) {
		return
	}
	msg_string := string(m.Content)
	bot_func := strings.Split(msg_string, " ")
	switch bot_func[0] {
		case conf.PREFIX+"play":
			if len(bot_func)==1{
				err_msg(s, m.ChannelID, "Try !play [song/url]")
			}
			var song_name string
			song_name = ""
			for i := 1; i<len(bot_func);i++{
				song_name+=bot_func[i]+" "
			}
			song_name = strings.TrimSuffix(song_name, " ")
			go straem_to_discord(s, m, song_name)
		case conf.PREFIX+"stop":
			go stop_stream(s, m)
		case conf.PREFIX+"pause":
			go pause_stream(s, m)
		case conf.PREFIX+"skip":
			go skip_music(s,m)
		case conf.PREFIX+"gachi":
			go gachi(s,m.ChannelID)
	}
	if (m.Content=="Снюс"){
		s.ChannelMessageSend(m.ChannelID, "Чесвин")
	}
	if (m.Content=="Чесвин"){
		s.ChannelMessageSend(m.ChannelID, "Снюс")
	}
}
