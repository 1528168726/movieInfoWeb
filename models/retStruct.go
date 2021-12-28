package models

type MoviesCoverInfo struct {
	Id        string
	Url       string
	ImgUrl    string
	Name      string
	BriefInfo string
	Type      []string
}

type CelebrityCoverInfo struct {
	Id        string
	Url       string
	ImgUrl    string
	Name      string
	BriefInfo string
}

type MovieInfo struct {
	Id           string
	ImgUrl       string
	ChineseName  string
	OtherName    string
	DirectorName string
	Year         int
	Type         []string
	Country      string
	Language     string
	Length       int
	DoubanScore  float32
	Tag          []string
	Summary      string
	RelatedMovie []MoviesCoverInfo
	Director     CelebrityCoverInfo
	Actors       []CelebrityCoverInfo
	Writers      []CelebrityCoverInfo
}

type CelebrityInfo struct {
	Name         string
	ImgUrl       string
	Gender       string
	DateOfBirth  string
	Info         string
	RelatedMovie []MoviesCoverInfo
}

type Top250MoviesInfo struct {
	Id        string
	Url       string
	ImgUrl    string
	Name      string
	Ranking   int
	BriefInfo string
}
