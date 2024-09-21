CREATE OR REPLACE FUNCTION validate_course()
    RETURNS TRIGGER AS $$
DECLARE
    lesson_record RECORD;
    theory_count INT := 0;
    practice_count INT := 0;
    test_count INT;
BEGIN
    FOR lesson_record IN
        SELECT * FROM public.lesson WHERE course_id = NEW.id
        LOOP
            IF lesson_record.score <= 0 THEN
                RAISE EXCEPTION 'Lesson % has an invalid score. It must be greater than 0', lesson_record.id;
            END IF;

            CASE lesson_record.type
                WHEN 'practice' THEN
                    SELECT COUNT(*) INTO test_count
                    FROM public.test
                    WHERE lesson_id = lesson_record.id;

                    IF test_count = 0 THEN
                        RAISE EXCEPTION 'Practice lesson % must have at least one test', lesson_record.id;
                    END IF;
                    practice_count := practice_count + 1;
                WHEN 'theory' THEN
                    IF lesson_record.theory_url IS NULL OR LENGTH(lesson_record.theory_url) = 0 THEN
                        RAISE EXCEPTION 'Theory lesson % must have a valid theory URL', lesson_record.id;
                    END IF;
                    theory_count := theory_count + 1;
                WHEN 'video' THEN
                    IF lesson_record.video_url IS NULL OR LENGTH(lesson_record.video_url) = 0 THEN
                        RAISE EXCEPTION 'Video lesson % must have a valid video URL', lesson_record.id;
                    END IF;
                    theory_count := theory_count + 1;
                END CASE;
        END LOOP;

    IF theory_count = 0 THEN
        RAISE EXCEPTION 'Course must have at least one theory or video lesson';
    END IF;

    IF practice_count = 0 THEN
        RAISE EXCEPTION 'Course must have at least one practice lesson';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_validate_course
    BEFORE UPDATE ON public.course
    FOR EACH ROW EXECUTE FUNCTION validate_course();
