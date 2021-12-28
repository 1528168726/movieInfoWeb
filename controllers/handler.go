package controllers

import (
	"databaseHW/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const pageDataNum = 50

func HandlerIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func HandlerActors(c *gin.Context) {
	var err error
	curPage := 0
	if page := c.Query("page"); page != "" {
		curPage, err = strconv.Atoi(page)
		if err != nil {
			curPage = 1
		}
	} else {
		curPage = 1
	}
	celebrities, allNum := models.GetActorsList(nil, (curPage-1)*pageDataNum, pageDataNum)
	if len(*celebrities) <= 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	c.HTML(http.StatusOK, "actors.html", gin.H{
		"curPage":     curPage,
		"allPage":     (allNum-1)/pageDataNum + 1,
		"celebrities": celebrities,
		"nextPage":    curPage + 1,
	})
}

func HandlerCelebrity(c *gin.Context) {
	id := c.Param("id")
	celebrity := models.GetCelebrity(id)
	if celebrity == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	c.HTML(http.StatusOK, "celebrity.html", gin.H{
		"celebrity": celebrity,
	})
}

func HandlerDirectors(c *gin.Context) {
	var err error
	curPage := 0
	if page := c.Query("page"); page != "" {
		curPage, err = strconv.Atoi(page)
		if err != nil {
			curPage = 1
		}
	} else {
		curPage = 1
	}
	celebrities, allNum := models.GetDirectorsList(nil, (curPage-1)*pageDataNum, pageDataNum)
	if len(*celebrities) <= 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	c.HTML(http.StatusOK, "directors.html", gin.H{
		"curPage":     curPage,
		"allPage":     (allNum-1)/pageDataNum + 1,
		"celebrities": celebrities,
	})
}

func HandlerMovies(c *gin.Context) {
	var err error
	curPage := 0
	curType := c.Query("type")
	if page := c.Query("page"); page != "" {
		curPage, err = strconv.Atoi(page)
		if err != nil {
			curPage = 1
		}
	} else {
		curPage = 1
	}
	movies, allNum := models.GetMoviesList(map[string]string{"type": curType}, (curPage-1)*pageDataNum, pageDataNum)
	if len(*movies) <= 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	c.HTML(http.StatusOK, "movies.html", gin.H{
		"curPage": curPage,
		"allPage": (allNum-1)/pageDataNum + 1,
		"movies":  movies,
		"curType": curType,
	})
}

func HandlerSearch(c *gin.Context) {
	searchKey := c.Query("searchKey")
	if searchKey == "" {
		c.HTML(http.StatusOK, "search.html", gin.H{
			"searchKey": "你什么都没有搜索",
		})
		return
	}
	movies := models.SearchMovies(searchKey)
	celebrities := models.SearchCelebrity(searchKey)
	c.HTML(http.StatusOK, "search.html", gin.H{
		"searchKey":         searchKey,
		"searchMovies":      movies,
		"searchCelebrities": celebrities,
	})

}

func HandlerSubject(c *gin.Context) {
	id := c.Param("id")
	movie := models.GetMovie(id)
	if movie == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	c.HTML(http.StatusOK, "subject.html", gin.H{
		"movie": movie,
	})
}

func HandlerTop250(c *gin.Context) {
	var err error
	curPage := 1
	if page := c.Query("page"); page != "" {
		curPage, err = strconv.Atoi(page)
		if err != nil {
			curPage = 1
		}
	} else {
		curPage = 1
	}
	movies := models.GetTop250((curPage-1)*pageDataNum, pageDataNum)
	if len(*movies) <= 0 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	c.HTML(http.StatusOK, "top250.html", gin.H{
		"movies":  movies,
		"curPage": curPage,
	})
}

func HandlerOscar(c *gin.Context) {
	session, err := strconv.Atoi(c.Param("session"))
	if err != nil || session <= 0 || session > 92 {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "index error"})
		c.Error(errors.New("index error"))
		return
	}
	prizes := models.GetOscarList(session)
	if prizes == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "can't get data"})
		c.Error(errors.New("can't get data"))
		return
	}
	forList := make([]int, 92)
	for i := 0; i < len(forList); i++ {
		forList[i] = 92 - i
	}
	c.HTML(http.StatusOK, "oscar.html", gin.H{
		"prizes":  prizes,
		"session": session,
		"forList": forList,
	})
}
