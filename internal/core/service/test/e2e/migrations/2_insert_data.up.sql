-- insert users
insert into public.user (id, email, password, name, surname, phone)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'paw1a@yandex.ru', '123', 'Pavel', 'Shpakovskiy', '+79992233444');
insert into public.user (id, email, password, name, surname, city)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'hanoys@mail.ru', 'qwerty', 'Timur', 'Musin', 'Moscow');
insert into public.user (id, email, password, name, surname, phone)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'emir@gmail.com', '12345', 'Emir', 'Shimshir', '+79992233555');

-- insert schools
insert into school (id, name, description, owner_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7034cc', 'school1', 'desc1', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into school (id, name, description, owner_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7034cd', 'school2', 'desc2', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');

-- insert courses
insert into course (id, name, description, school_id, level, price, image_url, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7027ca', 'course1', 'Course description', '30e18bc1-4354-4937-9a3b-03cf0b7034cc',
        4, 1200, 'https://cdn.stepik.net/media/cache/images/courses/63054/cover_foIuz1t/6bc976a3abd69e9e3e5163a5973a8ccf.jpg',
        'russian', 'draft');
insert into course (id, name, description, school_id, level, price, image_url, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7027cb', 'course2', 'Course description', '30e18bc1-4354-4937-9a3b-03cf0b7034cc',
        2, 1500, 'https://cdn.stepik.net/media/cache/images/courses/63054/cover_foIuz1t/6bc976a3abd69e9e3e5163a5973a8ccf.jpg',
        'english', 'published');
insert into course (id, name, description, school_id, level, price, image_url, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7026cc', 'course3', 'Course description', '30e18bc1-4354-4937-9a3b-03cf0b7034cd',
        3, 12000, 'https://cdn.stepik.net/media/cache/images/courses/104774/cover_3mkW6vV/f5ae6ccd6cc779eeaecfa9da9d567613.png',
        'russian', 'ready');
insert into course (id, name, description, school_id, level, price, image_url, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7026cd', 'course4', 'Course description', '30e18bc1-4354-4937-9a3b-03cf0b7034cd',
        2, 0, 'english', 'published');

-- insert reviews
insert into review (id, text, course_id, user_id)
values ('30e18bc1-4354-4937-9a4d-03cf0b7021ca', 'review1 text',
        '30e18bc1-4354-4937-9a4d-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into review (id, text, course_id, user_id)
values ('30e18bc1-4354-4937-9a4d-03cf0b7021cb', 'review2 text',
        '30e18bc1-4354-4937-9a4d-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');
insert into review (id, text, course_id, user_id)
values ('30e18bc1-4354-4937-9a4d-03cf0b7021cc', 'review3 text',
        '30e18bc1-4354-4937-9a4d-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');

-- insert school teachers
insert into school_teacher (teacher_id, school_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7034cc');
insert into school_teacher (teacher_id, school_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7034cc');

-- insert student courses
insert into course_student (student_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a4d-03cf0b7027ca');
insert into course_student (student_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a4d-03cf0b7027cb');

-- insert teacher courses
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a4d-03cf0b7027ca');
insert into course_teacher (teacher_id, course_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a4d-03cf0b7027cb');

-- insert lessons
insert into lesson (id, title, type, score, theory_url, video_url, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7022ca', 'lesson1', 'theory', 10,
        'url', null, '30e18bc1-4354-4937-9a4d-03cf0b7027ca');
insert into lesson (id, title, type, score, theory_url, video_url, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7022cb', 'lesson2', 'video', 10,
        null, 'url', '30e18bc1-4354-4937-9a4d-03cf0b7027ca');
insert into lesson (id, title, type, score, theory_url, video_url, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7022cc', 'lesson3', 'practice', 10,
        null, null, '30e18bc1-4354-4937-9a4d-03cf0b7027ca');

-- insert tests
insert into test (id, task_url, options, answer, score, level, lesson_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027ca', 'url', E'opt1\nopt2\nopt3',
        'opt1', 12, 2, '30e18bc1-4352-4937-9a3b-03cf0b7022cc');
insert into test (id, task_url, options, answer, score, level, lesson_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027cb', 'url', E'opt1\nopt2',
        'opt2', 12, 2, '30e18bc1-4352-4937-9a3b-03cf0b7022cc');
insert into test (id, task_url, options, answer, score, level, lesson_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027cc', 'url', E'opt1',
        'opt1', 12, 2, '30e18bc1-4352-4937-9a3b-03cf0b7022cc');
