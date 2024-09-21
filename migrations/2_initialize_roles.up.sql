create role root with
    noinherit
    superuser
    createrole
    login
    password 'root';

create role guest with
    noinherit
    login
    password 'guest';

create role authenticated with
    noinherit
    login
    password 'authenticated';

grant all on all tables in schema public to root;

grant select on table
    public.school,
    public.course,
    public.school_teacher,
    public.user,
    public.review to guest;
grant insert on table public.user to guest;

grant select on all tables in schema public to authenticated;
grant insert, update, delete on table
    public.review,
    public.user,
    public.course,
    public.school,
    public.certificate,
    public.lesson,
    public.test,
    public.lesson_stat,
    public.test_stat,
    public.course_student,
    public.course_teacher,
    public.school_teacher to authenticated;
