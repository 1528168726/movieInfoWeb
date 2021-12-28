package models

import "gorm.io/gorm"

func GetActorsList(searchKey map[string]string, begin int, num int) (*[]CelebrityCoverInfo, int) {
	var celeInfo []Celebrity
	result := DB.Select([]string{"picture", "name", "id"}).Where("id in (?)", DB.Select("cast_id").
		Limit(num).Offset(begin).Order("num desc").Model(&NumCastsMovie{})).Find(&celeInfo)
	var covers = make([]CelebrityCoverInfo, result.RowsAffected)
	for i := 0; i < int(result.RowsAffected); i++ {
		covers[i].Url = `/celebrity/` + string(celeInfo[i].Id)
		covers[i].Name = celeInfo[i].Name
		covers[i].ImgUrl = celeInfo[i].Picture
	}
	var allNum int
	row := DB.Table("num_casts_movie").Select("count(*)").Row()
	row.Scan(&allNum)

	return &covers, allNum
}

func GetDirectorsList(searchKey map[string]string, begin int, num int) (*[]CelebrityCoverInfo, int) {
	var celeInfo []Celebrity
	var allNum int

	result := DB.Select([]string{"picture", "name", "id"}).Where("id in (?)", DB.Select("director_id").
		Where("director_id is not null").Limit(num).Offset(begin).Order("num desc").Model(&NumDirectorsMovie{})).Find(&celeInfo)
	var covers = make([]CelebrityCoverInfo, result.RowsAffected)
	for i := 0; i < int(result.RowsAffected); i++ {
		covers[i].Url = `/celebrity/` + string(celeInfo[i].Id)
		covers[i].Name = celeInfo[i].Name
		covers[i].ImgUrl = celeInfo[i].Picture
	}

	row := DB.Table("num_directors_movie").Select("count(*)").Row()
	row.Scan(&allNum)

	return &covers, allNum
}

func GetMoviesList(searchKey map[string]string, begin int, num int) (*[]MoviesCoverInfo, int) {
	var movieInfo []Movie
	var curType string
	var allNum int
	if searchKey != nil {
		curType = searchKey["type"]
	}
	var result *gorm.DB
	if curType != "" {
		result = DB.Select([]string{"picture", "ChineseName", "id"}).
			Where("id in (?)", DB.Distinct("movie_subject_id").
				Where("type=(?)", curType).Model(&MovieType{})).
			Order("director_id is null,length<90").
			Limit(num).Offset(begin).Find(&movieInfo)
		row := DB.Model(&MovieType{}).Select("count(*)").Where("type=(?)", curType).Row()
		row.Scan(&allNum)

	} else {
		result = DB.Select([]string{"picture", "chinese_name", "id"}).
			Order("director_id is null,length<90").
			Limit(num).Offset(begin).
			Find(&movieInfo)
		row := DB.Table("movies").Select("count(*)").Row()
		row.Scan(&allNum)
	}
	var types []MovieType
	DB.Where("movie_subject_id in (?)", result.Select("id")).Find(&types)

	var covers = make([]MoviesCoverInfo, result.RowsAffected)
	for i := 0; i < int(result.RowsAffected); i++ {
		covers[i].Url = `/subject/` + string(movieInfo[i].Id)
		covers[i].Name = movieInfo[i].ChineseName
		covers[i].ImgUrl = movieInfo[i].Picture
		for j := 0; j < len(types); j++ {
			if types[j].MovieSubjectId == movieInfo[i].Id {
				covers[i].Type = append(covers[i].Type, types[j].Type)
			}
		}
	}
	return &covers, allNum
}

func GetMovie(id string) *MovieInfo {
	if id == "" {
		return nil
	}
	movieInfo := new(MovieInfo)
	var dbMovie Movie
	var dbType []MovieType
	var dbTag []MovieTag
	var dbRelated []Movie
	var dbDirector Celebrity
	var dbActors []Celebrity
	var dbWriter []Celebrity
	var result *gorm.DB
	result = DB.Where("id=(?)", id).First(&dbMovie)
	if result.RowsAffected == 0 {
		return nil
	}
	DB.Where("movie_subject_id=(?)", id).Find(&dbType)
	DB.Where("movie_subject_id=(?)", id).Find(&dbTag)
	DB.Select([]string{"picture", "ChineseName", "id"}).
		Where("id in (?)", DB.Select("related_movie_id").
			Where("movie_subject_id=(?)", id).Model(&RelatedMovie{})).
		Find(&dbRelated)
	if dbMovie.DirectorId != "" {
		DB.Select([]string{"picture", "name", "id"}).Where("id=(?)", dbMovie.DirectorId).First(&dbDirector)
	}
	DB.Select([]string{"picture", "name", "id"}).
		Where("id in (?)", DB.Distinct("cast_id").
			Where("movie_subject_id=(?)", id).Model(&MovieCast{})).
		Find(&dbActors)
	DB.Select([]string{"picture", "name", "id"}).
		Where("id in (?)", DB.Distinct("writer_id").
			Where("movie_subject_id=(?)", id).Model(&MovieWriter{})).
		Find(&dbWriter)

	movieInfo.Id = dbMovie.Id
	movieInfo.ImgUrl = dbMovie.Picture
	movieInfo.ChineseName = dbMovie.ChineseName
	movieInfo.OtherName = dbMovie.OtherName
	if dbMovie.DirectorId != "" {
		movieInfo.DirectorName = dbDirector.Name
		movieInfo.Director.Url = `/celebrity/` + string(dbDirector.Id)
		movieInfo.Director.Name = dbDirector.Name
		movieInfo.Director.ImgUrl = dbDirector.Picture
	}
	movieInfo.Year = dbMovie.Year
	movieInfo.Country = dbMovie.Country
	movieInfo.Language = dbMovie.Language
	movieInfo.Length = dbMovie.Length
	movieInfo.DoubanScore = dbMovie.DoubanScore
	movieInfo.Summary = dbMovie.Summary

	movieInfo.Type = make([]string, len(dbType))
	for i := 0; i < len(dbType); i++ {
		movieInfo.Type[i] = dbType[i].Type
	}
	movieInfo.Tag = make([]string, len(dbTag))
	for i := 0; i < len(dbTag); i++ {
		movieInfo.Tag[i] = dbTag[i].Tag
	}
	movieInfo.RelatedMovie = make([]MoviesCoverInfo, len(dbRelated))
	for i := 0; i < len(dbRelated); i++ {
		movieInfo.RelatedMovie[i].Url = `/subject/` + string(dbRelated[i].Id)
		movieInfo.RelatedMovie[i].Name = dbRelated[i].ChineseName
		movieInfo.RelatedMovie[i].ImgUrl = dbRelated[i].Picture
	}
	movieInfo.Actors = make([]CelebrityCoverInfo, len(dbActors))
	for i := 0; i < len(dbActors); i++ {
		movieInfo.Actors[i].Url = `/celebrity/` + string(dbActors[i].Id)
		movieInfo.Actors[i].Name = dbActors[i].Name
		movieInfo.Actors[i].ImgUrl = dbActors[i].Picture
	}
	movieInfo.Writers = make([]CelebrityCoverInfo, len(dbWriter))
	for i := 0; i < len(dbWriter); i++ {
		movieInfo.Writers[i].Url = `/celebrity/` + string(dbWriter[i].Id)
		movieInfo.Writers[i].Name = dbWriter[i].Name
		movieInfo.Writers[i].ImgUrl = dbWriter[i].Picture
	}
	return movieInfo
}

func GetCelebrity(id string) *CelebrityInfo {
	if id == "" {
		return nil
	}
	celebrityInfo := new(CelebrityInfo)
	var dbCelebrity Celebrity
	var dbRelatedMovie []Movie
	var result *gorm.DB
	result = DB.Where("id=(?)", id).First(&dbCelebrity)
	if result.RowsAffected == 0 {
		return nil
	}
	DB.Select([]string{"picture", "chinese_name", "id", "douban_score"}).
		Where("id in (?)", DB.Select("movie_subject_id").
			Where("cast_id=(?)", dbCelebrity.Id).Model(&MovieCast{})).
		Order("douban_score desc").Find(&dbRelatedMovie)

	celebrityInfo.Name = dbCelebrity.Name
	celebrityInfo.ImgUrl = dbCelebrity.Picture
	celebrityInfo.Gender = dbCelebrity.Gender
	celebrityInfo.DateOfBirth = dbCelebrity.DateOfBirth
	celebrityInfo.Info = dbCelebrity.Info

	celebrityInfo.RelatedMovie = make([]MoviesCoverInfo, len(dbRelatedMovie))
	for i := 0; i < len(dbRelatedMovie); i++ {
		celebrityInfo.RelatedMovie[i].Url = `/subject/` + string(dbRelatedMovie[i].Id)
		celebrityInfo.RelatedMovie[i].ImgUrl = dbRelatedMovie[i].Picture
		celebrityInfo.RelatedMovie[i].Name = dbRelatedMovie[i].ChineseName
	}
	return celebrityInfo
}

func GetTop250(begin int, num int) *[]Top250MoviesInfo {
	top250 := new([]Top250MoviesInfo)
	DB.Table("top_movies").Select([]string{"top_movies.ranking as ranking", "id", "picture as img_url", "chinese_name as name"}).
		Joins("inner join movies on movies.id=top_movies.movie_subject_id").
		Order("ranking").Limit(num).Offset(begin).
		Scan(&top250)

	for i := 0; i < len(*top250); i++ {
		(*top250)[i].Url = `/subject/` + string((*top250)[i].Id)
	}
	return top250
}

func SearchMovies(searchKey string) *[]CelebrityCoverInfo {
	var dbMovies []Movie
	if searchKey == "" {
		return nil
	}
	result := DB.Select([]string{"id", "picture", "chinese_name"}).
		Where("chinese_name like (?)", `%`+searchKey+`%`).
		Limit(50).Find(&dbMovies)
	if result.RowsAffected == 0 {
		return nil
	}
	movieInfo := make([]CelebrityCoverInfo, len(dbMovies))
	for i := 0; i < len(movieInfo); i++ {
		movieInfo[i].Url = `/subject/` + dbMovies[i].Id
		movieInfo[i].Name = dbMovies[i].ChineseName
		movieInfo[i].ImgUrl = dbMovies[i].Picture
	}
	return &movieInfo
}

func SearchCelebrity(searchKey string) *[]CelebrityCoverInfo {
	var celebrityInfo []CelebrityCoverInfo
	if searchKey == "" {
		return nil
	}
	result := DB.Model(&Celebrity{}).Select([]string{"id", "picture as img_url", "name as name"}).
		Where("name like (?)", `%`+searchKey+`%`).
		Limit(50).Scan(&celebrityInfo)
	if result.RowsAffected == 0 {
		return nil
	}
	for i := 0; i < len(celebrityInfo); i++ {
		celebrityInfo[i].Url = `/celebrity/` + celebrityInfo[i].Id
	}
	return &celebrityInfo
}

func GetOscarList(session int) *[]Oscar {
	if session <= 0 || session > 92 {
		return nil
	}
	var prizes []Oscar
	result := DB.Where("session =(?)", session).Find(&prizes)
	if result.RowsAffected == 0 {
		return nil
	}
	return &prizes
}
