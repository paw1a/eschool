insert into "user" (id, email, password, name, surname)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', 'paw1a@yandex.ru', '123', 'Pavel', 'Shpakovsliy');
insert into "user" (id, email, password, name, surname)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', 'hanoys@mail.ru', 'qwerty', 'Timur', 'Musin');
insert into "user" (id, email, password, name, surname)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cc', 'emir@gmail.com', '12345', 'Emir', 'Shimshir');

insert into school (id, name, description, owner_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7034cc', 'name1', 'desc1', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into school (id, name, description, owner_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7034cd', 'name2', 'desc2', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');

insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7027cb', 'course1', '30e18bc1-4354-4937-9a3b-03cf0b7034cc',
        4, 1200, 'russian', 'draft');
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7027cc', 'course2', '30e18bc1-4354-4937-9a3b-03cf0b7034cc',
        2, 1500, 'english', 'published');
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7026cb', 'course3', '30e18bc1-4354-4937-9a3b-03cf0b7034cd',
        3, 12000, 'russian', 'draft');
insert into course (id, name, school_id, level, price, language, status)
values ('30e18bc1-4354-4937-9a4d-03cf0b7026cc', 'course4', '30e18bc1-4354-4937-9a3b-03cf0b7034cd',
        2, 0, 'english', 'published');

insert into review (id, text, course_id, user_id)
values ('30e18bc1-4354-4937-9a4d-03cf0b7021cc', 'review1 text',
        '30e18bc1-4354-4937-9a4d-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7027ca');
insert into review (id, text, course_id, user_id)
values ('30e18bc1-4354-4937-9a4d-03cf0b7021cd', 'review1 text',
        '30e18bc1-4354-4937-9a4d-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7027cb');

insert into certificate (id, name, score, grade, created_at, user_id, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027cb', 'cert1', 120, 'gold', now(),
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a4d-03cf0b7027cb');
insert into certificate (id, name, score, grade, created_at, user_id, course_id)
values ('30e18bc1-4352-4937-9a3b-03cf0b7027cc', 'cert1', 120, 'bronze', now(),
        '30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a4d-03cf0b7027cc');

insert into school_teacher (teacher_id, school_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027ca', '30e18bc1-4354-4937-9a3b-03cf0b7034cc');
insert into school_teacher (teacher_id, school_id)
values ('30e18bc1-4354-4937-9a3b-03cf0b7027cb', '30e18bc1-4354-4937-9a3b-03cf0b7034cc');
