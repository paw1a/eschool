package test

import "github.com/paw1a/eschool/internal/core/domain"

type TestStatBuilder struct {
	testStat domain.TestStat
}

func NewTestStatBuilder() *TestStatBuilder {
	return &TestStatBuilder{
		testStat: domain.TestStat{
			ID:     domain.NewID(),
			TestID: domain.NewID(),
			UserID: domain.NewID(),
			Score:  10,
		},
	}
}

func (b *TestStatBuilder) WithID(id domain.ID) *TestStatBuilder {
	b.testStat.ID = id
	return b
}

func (b *TestStatBuilder) WithTestID(testID domain.ID) *TestStatBuilder {
	b.testStat.TestID = testID
	return b
}

func (b *TestStatBuilder) WithUserID(userID domain.ID) *TestStatBuilder {
	b.testStat.UserID = userID
	return b
}

func (b *TestStatBuilder) WithScore(score int) *TestStatBuilder {
	b.testStat.Score = score
	return b
}

func (b *TestStatBuilder) Build() domain.TestStat {
	return b.testStat
}

type LessonStatBuilder struct {
	lessonStat domain.LessonStat
}

func NewLessonStatBuilder() *LessonStatBuilder {
	return &LessonStatBuilder{
		lessonStat: domain.LessonStat{
			ID:        domain.NewID(),
			LessonID:  domain.NewID(),
			UserID:    domain.NewID(),
			Score:     20,
			TestStats: nil,
		},
	}
}

func (b *LessonStatBuilder) WithID(id domain.ID) *LessonStatBuilder {
	b.lessonStat.ID = id
	return b
}

func (b *LessonStatBuilder) WithLessonID(lessonID domain.ID) *LessonStatBuilder {
	b.lessonStat.LessonID = lessonID
	return b
}

func (b *LessonStatBuilder) WithUserID(userID domain.ID) *LessonStatBuilder {
	b.lessonStat.UserID = userID
	return b
}

func (b *LessonStatBuilder) WithScore(score int) *LessonStatBuilder {
	b.lessonStat.Score = score
	return b
}

func (b *LessonStatBuilder) WithTestStats(testStats []domain.TestStat) *LessonStatBuilder {
	b.lessonStat.TestStats = testStats
	return b
}

func (b *LessonStatBuilder) AddTestStat(testStat domain.TestStat) *LessonStatBuilder {
	b.lessonStat.TestStats = append(b.lessonStat.TestStats, testStat)
	return b
}

func (b *LessonStatBuilder) Build() domain.LessonStat {
	return b.lessonStat
}
