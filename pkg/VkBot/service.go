package VkBot

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"log"
	"math/rand"
	"os"
	"time"
)

func (b *Bot) GetLikes() {
	p := params.NewLikesGetListBuilder()

	p.Type("post")
	p.Count(1000)
	p.OwnerID(-191968702)
	p.ItemID(120)

	list, err := b.bot.LikesGetListExtended(p.Params)

	if err != nil {
		log.Fatal(err)
	}

	wholiked(list)

}

func wholiked(list api.LikesGetListExtendedResponse) {
	data, err := json.MarshalIndent(list.Items, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var res []User
	reader := bytes.NewReader(data)
	dec := json.NewDecoder(reader)
	err = dec.Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	var (
		usersFile  string
		id         int
		numWinners int
	)
	flag.StringVar(&usersFile, "user", "users.csv", "List of users that liked a specific post.")
	flag.IntVar(&id, "id", 0, "Id of a post.")
	flag.IntVar(&numWinners, "winners", 0, "Pick or no the winner.")
	flag.Parse()

	usernames := make([]string, 0, len(res))
	for _, users := range res {
		usernames = append(usernames, users.FirstName+"_"+users.LastName)
	}

	existingUsers := existing(usersFile)

	allUserNames := merge(usernames, existingUsers)

	err = writeUsers(usersFile, allUserNames)
	if err != nil {
		log.Fatal(err)
	}

	if numWinners == 0 {
		return
	}

	winners := pickWinners(existingUsers, numWinners)
	fmt.Println("The winners are: ")
	for _, username := range winners {
		fmt.Printf("\t%s\n", username)
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//Id        int    `json:"id"`
}

func existing(usersFile string) []string {
	f, err := os.Open(usersFile)
	if err != nil {
		return []string{}
	}

	defer f.Close()

	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	users := make([]string, 0, len(lines))
	for _, line := range lines {
		users = append(users, line[0])
	}
	return users
}

func merge(a, b []string) []string {
	uniq := make(map[string]struct{}, 0)
	for _, user := range a {
		uniq[user] = struct{}{}
	}

	for _, user := range b {
		uniq[user] = struct{}{}
	}

	res := make([]string, 0, len(uniq))
	for user := range uniq {
		res = append(res, user)
	}
	return res
}

func writeUsers(usersFile string, users []string) error {
	f, err := os.OpenFile(usersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)
	for _, user := range users {
		if err = w.Write([]string{user}); err != nil {
			return err
		}
	}
	w.Flush()

	if err = w.Error(); err != nil {
		return err
	}
	return nil
}

func pickWinners(users []string, numWinners int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(users))
	winners := perm[:numWinners]
	res := make([]string, 0, numWinners)
	for _, idx := range winners {
		res = append(res, users[idx])
	}
	return res
}
