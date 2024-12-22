-- insert users
insert into public.user (id, email, password, name, surname, phone)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'owner@mail.ru', '123', 'Owner', 'Shpakovskiy', '+79992233444');
insert into public.user (id, email, password, name, surname, phone)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'teacher@mail.ru', '123', 'Teacher', 'Ivanov', '+79992233445');
insert into public.user (id, email, password, name, surname, city)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'student@mail.ru', '123', 'Student', 'Musin', 'Moscow');

-- insert schools
insert into school (id, name, description, owner_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'School 1', 'description', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');

-- insert courses
insert into course (id, name, description, school_id, level, price, language, image_url, status)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'Course 1', 'Course description',
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', 4, 0, 'russian',
        'https://cdn.stepik.net/media/cache/images/courses/63054/cover_foIuz1t/6bc976a3abd69e9e3e5163a5973a8ccf.jpg',
        'published');
insert into course (id, name, description, school_id, level, price, language, image_url, status)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'Course 2', 'Course description',
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', 2, 0, 'english',
        'https://cdn.stepik.net/media/cache/images/courses/104774/cover_3mkW6vV/f5ae6ccd6cc779eeaecfa9da9d567613.png',
        'published');
insert into course (id, name, description, school_id, level, price, language, image_url, status)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'Course 3', 'Course description',
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', 3, 10, 'russian',
        'https://cdn.stepik.net/media/cache/images/courses/154096/cover_LOGEM4B/5a9b962caa4a946a3d23bdfd49e0480a.png',
        'draft');

-- insert school teachers
insert into school_teacher (teacher_id, school_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');

-- insert student courses
insert into course_student (student_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');

-- insert teacher courses
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cc');
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7027cc');

-- insert lessons
insert into lesson (id, title, type, score, theory_url, video_url, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'Lesson 1', 'theory', 10,
        'http://theory.md', null, '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into lesson (id, title, type, score, theory_url, video_url, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'Lesson 2', 'video', 10,
        null, 'http://video.mp4', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into lesson (id, title, type, score, theory_url, video_url, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'Lesson 3', 'practice', 10,
        null, null, '30e18bc1-4354-4937-9a3b-03cf0b7027ca');

-- insert tests
insert into test (id, task_url, options, answer, score, level, lesson_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'url', E'opt1\nopt2\nopt3',
        'opt1', 12, 2, '30e18bc1-4354-4937-9a3b-03cf0b7027cc');
insert into test (id, task_url, options, answer, score, level, lesson_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'url', E'opt1\nopt2',
        'opt2', 24, 2, '30e18bc1-4354-4937-9a3b-03cf0b7027cc');
insert into test (id, task_url, options, answer, score, level, lesson_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'url', E'opt1',
        'opt1', 12, 2, '30e18bc1-4354-4937-9a3b-03cf0b7027cc');

-- insert test_stats
insert into test_stat (id, score, user_id, test_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 12, '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into test_stat (id, score, user_id, test_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 0, '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');
insert into test_stat (id, score, user_id, test_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 12, '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cc');

-- insert lesson_stats
insert into lesson_stat (id, score, user_id, lesson_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 0, '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into lesson_stat (id, score, user_id, lesson_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 10, '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');
insert into lesson_stat (id, score, user_id, lesson_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 10, '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cc');
