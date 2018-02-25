package main

//go:generate qtc -dir=templates

import (
	"github.com/gin-gonic/gin"

	"sync"

	"time"

	"bytes"

	"strconv"

	"./templates"

	"structs"

	"sort"

	"net/http"

)


//implementing sort interface in order to use Sort function. the order will be according to date, ASC 
type DateSort []*structs.User 

func (lst DateSort) Len() int           { return len(lst) }
func (lst DateSort) Swap(i, j int)      { lst[i], lst[j] = lst[j], lst[i] }
func (lst DateSort) Less(i, j int) bool { return lst[i].CreatedDate.Sub(lst[j].CreatedDate) < 0}

//global vars - 
//1. mutex lock - to make sure the user ID is thread safe
//2. DB - a map with id as the key, and User struct as the value
//3. IDS - will hold the next available Id
var Lock sync.Mutex
var DB = make(map[int]*structs.User)
var IDS int = 1

//the main app router with 3 handlers:
//1. list of users page handler GET - /
//2. adding/updating user handler POST - /add
//3. display user page GET - /get/*id
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		lst := make([]*structs.User, 0, len(DB))

		for  _, value := range DB {
		   lst = append(lst, value)
		}
		sort.Sort(DateSort(lst))

		p := &templates.MainPage{
			Users: lst}
		var buf bytes.Buffer
		templates.WritePageTemplate(&buf, p)
		c.Data(200, "text/html; charset=utf-8", buf.Bytes())
	})

	r.POST("/add", func(c *gin.Context) {
		first := c.PostForm("first")
		last := c.PostForm("last")
		id, err := strconv.Atoi(c.PostForm("id"))
		if err != nil || id == -1 {
			Lock.Lock()
			DB[IDS] = &structs.User{IDS,first,last,time.Now()}
			IDS++
			Lock.Unlock()
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			user, found := DB[id]
			if found {
				Lock.Lock()
				user.FirstName = first
				user.LastName = last
				Lock.Unlock()
				c.Redirect(http.StatusMovedPermanently, "/")
			} else {
				c.JSON(500,gin.H{"err":"id wasnt found"})
			}
		}
		
	})

	r.GET("/get/*id", func(c *gin.Context){
		var p *templates.UserPage

		//when using * in params instead of :, the / is added, so i get rid of it by slicing
		id, err := strconv.Atoi(c.Param("id")[1:])
		if err !=nil {
			p = &templates.UserPage{
				User: nil}
		} else {
			p = &templates.UserPage{
				User: DB[id]}
		}
		var buf bytes.Buffer
		templates.WritePageTemplate(&buf, p)
		c.Data(200, "text/html; charset=utf-8", buf.Bytes())
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}