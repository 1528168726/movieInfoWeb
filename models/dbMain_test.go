package models

import (
	"databaseHW/config"
	"fmt"
	"testing"
)

func init() {
	info, err := config.GetDbInfo("dbConfig.json")
	if err != nil {
		println("init db: GetDbInfo false", err.Error())
		return
	}
	err = DbInit(&info)
	if err != nil {
		println("init db: DbInit false", err.Error())
		return
	}
	println("init DB succeed")
}

func TestGetActorsList(t *testing.T) {
	covers, _ := GetActorsList(nil, 0, 50)
	//fmt.Println(covers)
	if len(*covers) > 0 {
		t.Log("ok")
	} else {
		t.Error("no data")
	}

}

func TestGetDirectorsList(t *testing.T) {
	covers, _ := GetDirectorsList(nil, 0, 50)
	fmt.Println(covers)
	if len(*covers) > 0 {
		t.Log("ok")
	} else {
		t.Error("no data")
	}

}

func TestGetMoviesList(t *testing.T) {
	covers, _ := GetMoviesList(nil, 0, 50)
	fmt.Println(*covers)
	if len(*covers) > 0 {
		t.Log("ok")
	} else {
		t.Error("no data")
	}

}
