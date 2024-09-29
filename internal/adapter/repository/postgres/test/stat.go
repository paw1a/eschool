package test

import "github.com/paw1a/eschool/internal/core/domain"

// TestStatBuilder - билдер для TestStat в тестах
type TestStatBuilder struct {
	testStat domain.TestStat
}

// NewTestStatBuilder - создает новый билдер для TestStat
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

// WithID - устанавливает ID
func (b *TestStatBuilder) WithID(id domain.ID) *TestStatBuilder {
	b.testStat.ID = id
	return b
}

// WithTestID - устанавливает TestID
func (b *TestStatBuilder) WithTestID(testID domain.ID) *TestStatBuilder {
	b.testStat.TestID = testID
	return b
}

// WithUserID - устанавливает UserID
func (b *TestStatBuilder) WithUserID(userID domain.ID) *TestStatBuilder {
	b.testStat.UserID = userID
	return b
}

// WithScore - устанавливает Score
func (b *TestStatBuilder) WithScore(score int) *TestStatBuilder {
	b.testStat.Score = score
	return b
}

// Build - возвращает готовую TestStat
func (b *TestStatBuilder) Build() domain.TestStat {
	return b.testStat
}

// LessonStatBuilder - билдер для LessonStat в тестах
type LessonStatBuilder struct {
	lessonStat domain.LessonStat
}

// NewLessonStatBuilder - создает новый билдер для LessonStat
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

// WithID - устанавливает ID
func (b *LessonStatBuilder) WithID(id domain.ID) *LessonStatBuilder {
	b.lessonStat.ID = id
	return b
}

// WithLessonID - устанавливает LessonID
func (b *LessonStatBuilder) WithLessonID(lessonID domain.ID) *LessonStatBuilder {
	b.lessonStat.LessonID = lessonID
	return b
}

// WithUserID - устанавливает UserID
func (b *LessonStatBuilder) WithUserID(userID domain.ID) *LessonStatBuilder {
	b.lessonStat.UserID = userID
	return b
}

// WithScore - устанавливает Score
func (b *LessonStatBuilder) WithScore(score int) *LessonStatBuilder {
	b.lessonStat.Score = score
	return b
}

// WithTestStats - добавляет массив TestStats
func (b *LessonStatBuilder) WithTestStats(testStats []domain.TestStat) *LessonStatBuilder {
	b.lessonStat.TestStats = testStats
	return b
}

// AddTestStat - добавляет один TestStat в массив TestStats
func (b *LessonStatBuilder) AddTestStat(testStat domain.TestStat) *LessonStatBuilder {
	b.lessonStat.TestStats = append(b.lessonStat.TestStats, testStat)
	return b
}

// Build - возвращает готовую LessonStat
func (b *LessonStatBuilder) Build() domain.LessonStat {
	return b.lessonStat
}
