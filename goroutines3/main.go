package main

import (
	"fmt"
	"time"
)

var (
	defaultTags = []string{"SystemUser", "User", "NewUser", "System"}
)

type Tag struct {
	Name, Type string
}

type User struct {
	Id, Name, LastName, Status string
	Tags []*Tag
}

type Post struct {
	Title string
	Status string
	UserId string
}

func main(){
	blocking()
}

func blocking(){
	user := &User{}
	done := make(chan bool)//unbuffered channel

	go func(){
				fmt.Println("[Second-GoRoutine] Start building User")
				buildingUser(user)
				fmt.Println("[Second-GoRoutine] Finish building User")
				done <- true

				fmt.Println("[Second-GoRoutine] Set default User tags")
				setDefaultTags(user)
	}()

	fmt.Println("[Main-GoRoutine] Start importing posts")
	posts := importingPosts()
	fmt.Println("[Main-GoRoutine] Finish importing posts")
	fmt.Println("[Main-GoRoutine] -----Waiting ---------")
	<- done

	mergeUserPosts(user, posts)
	fmt.Println("Done!!")
	fmt.Printf("User %v\n", user)
	for _, post := range posts{
		fmt.Printf("Post %v\n", post)
	}
}

func mergeUserPosts(user *User, posts []*Post){
	fmt.Println("[Main-GoRoutine] Start merging user posts")
	for _,post := range posts{
		post.UserId = user.Id
	}
	fmt.Println("[Main-GoRoutine] Finished merging user posts")
}

func importingPosts() []*Post{
	time.Sleep(1 * time.Second)
	titles := []string{"Post 1", "Random Post", "Second Post"}
	posts := []*Post{}
	for _,title := range titles{
		posts = append(posts, &Post{Title: title, Status: "draft"})
	}
	return posts
}

func buildingUser(user *User){
	time.Sleep(2 * time.Second)
	user.Name = "John"
	user.LastName = "Doe"
	user.Status = "active"
	user.Id = "1"
}

func setDefaultTags(user *User){
	time.Sleep(1 * time.Second)
	for _,tagName := range defaultTags{
		user.Tags = append(user.Tags, &Tag{Name: tagName, Type: "System"})
	}
}