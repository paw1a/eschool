-- insert users
insert into public.user (id, email, password, name, surname, phone)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'paw1a@yandex.ru', '123', 'Pavel', 'Shpakovsliy', '+79992233444');
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
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7027ca', 'course1', '30e18bc1-4354-4937-9a3b-03cf0b7034cc',
        4, 1200, 'russian', 'draft');
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7027cb', 'course2', '30e18bc1-4354-4937-9a3b-03cf0b7034cc',
        2, 1500, 'english', 'published');
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7026cc', 'course3', '30e18bc1-4354-4937-9a3b-03cf0b7034cd',
        3, 12000, 'russian', 'ready');
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7026cd', 'course4', '30e18bc1-4354-4937-9a3b-03cf0b7034cd',
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

-- insert certificates
insert into certificate (id, name, score, grade, created_at, user_id, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027ca', 'course 1 cert', 120, 'gold', now(),
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a4d-03cf0b7027ca');
insert into certificate (id, name, score, grade, created_at, user_id, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027cb', 'course 2 cert', 50, 'bronze', now(),
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a4d-03cf0b7027cb');

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
