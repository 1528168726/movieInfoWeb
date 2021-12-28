package models

type Text1 struct {
	Id   string
	Name string
}

type Celebrity struct {
	Id          string `gorm:"primaryKey;not null;type:varchar(20)"`
	Picture     string `gorm:"type:varchar(200)"`
	Name        string `gorm:"type:varchar(100)"`
	Gender      string `gorm:"type:varchar(10)"`
	DateOfBirth string `gorm:"type:varchar(30)"`
	Info        string `gorm:"type:varchar(4000)"`
}

type Movie struct {
	Id          string  `gorm:"primaryKey;type:varchar(20)"`
	Picture     string  `gorm:"type:varchar(100)"`
	ChineseName string  `gorm:"type:varchar(80)"`
	OtherName   string  `gorm:"type:varchar(150)"`
	DirectorId  string  `gorm:"type:varchar(30)"`
	Year        int     `gorm:"type:int"`
	Country     string  `gorm:"type:varchar(100)"`
	Language    string  `gorm:"type:varchar(200)"`
	Length      int     `gorm:"type:int"`
	DoubanScore float32 `gorm:"type:float"`
	Summary     string  `gorm:"type:varchar(500)"`
}

type MovieType struct {
	MovieSubjectId string `gorm:"primaryKey;type:varchar(20)"`
	Type           string `gorm:"primaryKey;type:varchar(20)"`
}

type MovieTag struct {
	MovieSubjectId string `gorm:"primaryKey;type:varchar(20)"`
	Tag            string `gorm:"primaryKey;type:varchar(20)"`
}

type RelatedMovie struct {
	MovieSubjectId string `gorm:"primaryKey;type:varchar(20)"`
	RelatedMovieId string `gorm:"primaryKey;type:varchar(20)"`
}

type MovieCast struct {
	MovieSubjectId string `gorm:"primaryKey;type:varchar(20)"`
	CastId         string `gorm:"primaryKey;type:varchar(20)"`
}

type MovieWriter struct {
	MovieSubjectId string `gorm:"primaryKey;type:varchar(20)"`
	WriterId       string `gorm:"primaryKey;type:varchar(20)"`
}

type TopMovie struct {
	Ranking        int    `gorm:"primaryKey;type:int"`
	MovieSubjectId string `gorm:"type:varchar(20)"`
}

type Oscar struct {
	PrizeTitle       string `gorm:"primaryKey;not null;type:varchar(20)"`
	CelebrityName    string `gorm:"type:varchar(30)"`
	CelebrityId      string `gorm:"type:varchar(20)"`
	MovieChineseName string `gorm:"type:varchar(50)"`
	MovieSubjectId   string `gorm:"primaryKey;not null;type:varchar(20)"`
	Session          int    `gorm:"type:int"`
	Year             int    `gorm:"type:int"`
	ExistMovie       bool   `gorm:"type:bool"`
}

type NumCastsMovie struct {
	CastId string
	Num    int64
}

type NumDirectorsMovie struct {
	DirectorId string
	Num        int64
}

func (NumCastsMovie) TableName() string {
	return "num_casts_movie"
}

func (NumDirectorsMovie) TableName() string {
	return "num_directors_movie"
}
