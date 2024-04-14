package errs

import "errors"

var (
	ErrLessonContentUrlEmpty   = errors.New("lesson content url is empty")
	ErrLessonTestQuestionEmpty = errors.New("lesson test question is empty")
	ErrLessonTestOptionsEmpty  = errors.New("lesson test options is empty")
	ErrLessonTestInvalidLevel  = errors.New("lesson test level is invalid value")
	ErrLessonTestInvalidMark   = errors.New("lesson test mark is invalid value")
)

var (
	ErrCourseNotEnoughLessons               = errors.New("course must have at least 1 theory and 1 practice lessons")
	ErrCourseLessonInvalidScore             = errors.New("course lesson score must be > 0")
	ErrCoursePracticeLessonEmptyTests       = errors.New("course practice lesson must contain at least 1 test")
	ErrCoursePracticeLessonEmptyTestTaskUrl = errors.New("course practice lesson test has no question")
	ErrCoursePracticeLessonEmptyTestOptions = errors.New("course practice lesson test has no options")
	ErrCoursePracticeLessonInvalidTestScore = errors.New("course practice lesson test score must be > 0")
	ErrCourseTheoryLessonEmptyUrl           = errors.New("course theory lesson url is empty")
	ErrCourseVideoLessonEmptyUrl            = errors.New("course video lesson url is empty")
	ErrCourseReadyState                     = errors.New("course must be in draft state to make it ready")
	ErrCoursePublishedState                 = errors.New("course must be in ready state to publish it")
)

var (
	ErrCertificateCourseNotPassed = errors.New("course is not passed to make a certificate")
)

var (
	ErrDuplicate         = errors.New("record already exists")
	ErrNotExist          = errors.New("record does not exist")
	ErrUpdateFailed      = errors.New("record update failed")
	ErrDeleteFailed      = errors.New("record delete failed")
	ErrPersistenceFailed = errors.New("persistence internal error")
	ErrEnumValueError    = errors.New("enum value is out of scope")
	ErrTransactionError  = errors.New("transaction error occurred")
)

var (
	ErrFilenameEmpty        = errors.New("validation filename is empty error")
	ErrFilepathEmpty        = errors.New("validation filepath is empty error")
	ErrFiletypeNotSupported = errors.New("validation filetype is not supported")
	ErrFileReaderEmpty      = errors.New("validation file reader is nil error")
	ErrSaveFileError        = errors.New("failed to save file to object storage")
)
