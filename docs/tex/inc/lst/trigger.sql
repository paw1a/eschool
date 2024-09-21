create or replace function add_course_review()
    returns trigger as
$$
declare
    course_reviews_count numeric;
begin
    select count(*) from review where course_id=new.course_id into course_reviews_count;
    if (tg_op = 'DELETE') then
        update course set rating=((course_reviews_count * rating) - old.rating) /
            (course_reviews_count - 1) where id=old.course_id;
    elsif (tg_op = 'INSERT') then
        update course set rating=((course_reviews_count * rating) + new.rating) /
            (course_reviews_count + 1) where id=new.course_id;
    elsif (tg_op = 'UPDATE') then
        update course set rating=((course_reviews_count * rating) - old.rating) /
            (course_reviews_count - 1) where id=new.course_id;
        update course set rating=((course_reviews_count * rating) + new.rating) /
            (course_reviews_count + 1) where id=new.course_id;
    end if;
    return null;
end
$$ language plpgsql;

create trigger update_course
    after insert or update or delete
    on review for each row execute function add_course_review();