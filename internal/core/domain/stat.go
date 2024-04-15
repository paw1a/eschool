package domain

type LessonStat struct {
	ID        ID
	LessonID  ID
	UserID    ID
	Score     int
	TestStats []TestStat
}

type TestStat struct {
	ID     ID
	TestID ID
	UserID ID
	Score  int
}
