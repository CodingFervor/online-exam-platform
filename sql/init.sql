CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL, email VARCHAR(100) UNIQUE NOT NULL,
    role VARCHAR(20) DEFAULT 'student' CHECK (role IN ('admin','teacher','student')),
    status VARCHAR(20) DEFAULT 'active', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL, parent_id BIGINT REFERENCES categories(id)
);

CREATE TABLE questions (
    id BIGSERIAL PRIMARY KEY, title VARCHAR(500) NOT NULL, content TEXT,
    type VARCHAR(20) NOT NULL CHECK (type IN ('single_choice','multi_choice','true_false','fill_blank','essay')),
    options JSONB DEFAULT '[]', answer TEXT, explanation TEXT,
    score DECIMAL(5,2) DEFAULT 1.00, difficulty VARCHAR(10) DEFAULT 'medium' CHECK (difficulty IN ('easy','medium','hard')),
    category_id BIGINT REFERENCES categories(id), tags JSONB DEFAULT '[]',
    created_by BIGINT REFERENCES users(id), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE exams (
    id BIGSERIAL PRIMARY KEY, title VARCHAR(200) NOT NULL, description TEXT,
    duration INT NOT NULL DEFAULT 60, total_score DECIMAL(5,2) DEFAULT 100,
    pass_score DECIMAL(5,2) DEFAULT 60, question_ids JSONB DEFAULT '[]',
    shuffle BOOLEAN DEFAULT FALSE, max_attempts INT DEFAULT 1,
    start_time TIMESTAMP, end_time TIMESTAMP,
    status VARCHAR(20) DEFAULT 'draft' CHECK (status IN ('draft','published','archived')),
    created_by BIGINT REFERENCES users(id), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE exam_submissions (
    id BIGSERIAL PRIMARY KEY, exam_id BIGINT NOT NULL REFERENCES exams(id),
    user_id BIGINT NOT NULL REFERENCES users(id),
    start_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, submit_time TIMESTAMP,
    score DECIMAL(5,2), status VARCHAR(20) DEFAULT 'in_progress' CHECK (status IN ('in_progress','submitted','graded')),
    tab_switches INT DEFAULT 0, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE answers (
    id BIGSERIAL PRIMARY KEY, submission_id BIGINT NOT NULL REFERENCES exam_submissions(id),
    question_id BIGINT NOT NULL REFERENCES questions(id),
    answer TEXT, is_correct BOOLEAN, score DECIMAL(5,2) DEFAULT 0
);

CREATE TABLE certificates (
    id BIGSERIAL PRIMARY KEY, user_id BIGINT NOT NULL REFERENCES users(id),
    exam_id BIGINT NOT NULL REFERENCES exams(id),
    certificate_no VARCHAR(50) UNIQUE NOT NULL, score DECIMAL(5,2),
    issued_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cheating_logs (
    id BIGSERIAL PRIMARY KEY, submission_id BIGINT NOT NULL REFERENCES exam_submissions(id),
    user_id BIGINT NOT NULL REFERENCES users(id),
    type VARCHAR(20) NOT NULL CHECK (type IN ('tab_switch','copy_paste','multiple_login','time_anomaly')),
    detail TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_questions_category ON questions(category_id);
CREATE INDEX idx_submissions_exam ON exam_submissions(exam_id);
CREATE INDEX idx_submissions_user ON exam_submissions(user_id);
CREATE INDEX idx_answers_submission ON answers(submission_id);
CREATE INDEX idx_cheating_user ON cheating_logs(user_id);

INSERT INTO users (username, password, name, email, role) VALUES
('admin', '$2a$10$dummyhash', 'Admin', 'admin@exam.com', 'admin');
