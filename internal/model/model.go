package model

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Role      string    `json:"role" db:"role"` // admin, teacher, student
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Category struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	ParentID *int64 `json:"parent_id" db:"parent_id"`
}

type Question struct {
	ID         int64     `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Content    string    `json:"content" db:"content"`
	Type       string    `json:"type" db:"type"` // single_choice, multi_choice, true_false, fill_blank, essay
	Options    string    `json:"options" db:"options"` // JSON
	Answer     string    `json:"answer" db:"answer"`
	Explanation string   `json:"explanation" db:"explanation"`
	Score      float64   `json:"score" db:"score"`
	Difficulty string    `json:"difficulty" db:"difficulty"` // easy, medium, hard
	CategoryID int64     `json:"category_id" db:"category_id"`
	Tags       string    `json:"tags" db:"tags"`
	CreatedBy  int64     `json:"created_by" db:"created_by"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type Exam struct {
	ID          int64      `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Duration    int        `json:"duration" db:"duration"` // minutes
	TotalScore  float64    `json:"total_score" db:"total_score"`
	PassScore   float64    `json:"pass_score" db:"pass_score"`
	QuestionIDs string     `json:"question_ids" db:"question_ids"` // JSON array
	Shuffle     bool       `json:"shuffle" db:"shuffle"`
	MaxAttempts int        `json:"max_attempts" db:"max_attempts"`
	StartTime   *time.Time `json:"start_time" db:"start_time"`
	EndTime     *time.Time `json:"end_time" db:"end_time"`
	Status      string     `json:"status" db:"status"` // draft, published, archived
	CreatedBy   int64      `json:"created_by" db:"created_by"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

type ExamSubmission struct {
	ID          int64     `json:"id" db:"id"`
	ExamID      int64     `json:"exam_id" db:"exam_id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	StartTime   time.Time `json:"start_time" db:"start_time"`
	SubmitTime  *time.Time `json:"submit_time" db:"submit_time"`
	Score       *float64  `json:"score" db:"score"`
	Status      string    `json:"status" db:"status"` // in_progress, submitted, graded
	TabSwitches int       `json:"tab_switches" db:"tab_switches"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Answer struct {
	ID           int64  `json:"id" db:"id"`
	SubmissionID int64  `json:"submission_id" db:"submission_id"`
	QuestionID   int64  `json:"question_id" db:"question_id"`
	Answer       string `json:"answer" db:"answer"`
	IsCorrect    *bool  `json:"is_correct" db:"is_correct"`
	Score        float64 `json:"score" db:"score"`
}

type Certificate struct {
	ID           int64     `json:"id" db:"id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	ExamID       int64     `json:"exam_id" db:"exam_id"`
	CertificateNo string   `json:"certificate_no" db:"certificate_no"`
	Score        float64   `json:"score" db:"score"`
	IssuedAt     time.Time `json:"issued_at" db:"issued_at"`
}

type CheatingLog struct {
	ID           int64     `json:"id" db:"id"`
	SubmissionID int64     `json:"submission_id" db:"submission_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Type         string    `json:"type" db:"type"` // tab_switch, copy_paste, multiple_login
	Detail       string    `json:"detail" db:"detail"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
