CREATE TABLE tutor 
(
	tutor_id UUID PRIMARY KEY,
	name VARCHAR,
	last_name VARCHAR,
	email VARCHAR,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at BIGINT DEFAULT 0
);

CREATE TABLE course
(
	course_id UUID PRIMARY KEY,
	name VARCHAR,
	student_number INT,
	tutor_id UUID,
	started_at TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at BIGINT DEFAULT 0,
    CONSTRAINT fk_tutor_tutor_id FOREIGN KEY(tutor_id) REFERENCES tutor(tutor_id);
);

CREATE TABLE enrollment 
(
	id UUID PRIMARY KEY,
	course_id UUID,
	student_id UUID,
	CONSTRAINT fk_course_course_id FOREIGN KEY(course_id) REFERENCES course(course_id),
	CONSTRAINT fk_student_student_id FOREIGN KEY(student_id) REFERENCES student(student_id)
);

CREATE TABLE student_group
(
	group_id UUID PRIMARY KEY,
	name VARCHAR,
	course_id UUID,
	student_count INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at BIGINT DEFAULT 0,
	CONSTRAINT fk_course_course_id FOREIGN KEY(course_id) REFERENCES course(course_id)
);