import random
import uuid
import argparse
from faker import Faker
from faker.providers import phone_number
from dataclasses import dataclass

# Initialize Faker
fake = Faker()
fake.add_provider(phone_number)

# Define dataclasses for all tables
@dataclass
class User:
    id: uuid.UUID
    email: str
    password: str
    name: str
    surname: str
    phone: str
    city: str
    avatar: str

@dataclass
class School:
    id: uuid.UUID
    name: str
    description: str
    owner_id: uuid.UUID

@dataclass
class Course:
    id: uuid.UUID
    name: str
    school_id: uuid.UUID
    level: int
    price: int
    language: str
    status: str

@dataclass
class Lesson:
    id: uuid.UUID
    title: str
    type: str
    score: int
    theory_url: str
    video_url: str
    course_id: uuid.UUID

@dataclass
class Test:
    id: uuid.UUID
    task_url: str
    options: str
    answer: str
    score: int
    level: int
    lesson_id: uuid.UUID


def generate_data(user_count, school_count, course_count, lesson_count, test_count, output_file):
    # Generate Users
    sql_str = ''
    sql_str += "insert into public.user (id, email, password, name, surname, phone, city, avatar_url) values "
    users = []

    for i in range(0, user_count):
        user = User(id=uuid.uuid4(),
                    email=fake.email(),
                    password=fake.password(),
                    name=fake.last_name(),
                    surname=fake.first_name(),
                    phone=fake.phone_number(),
                    city=fake.city(),
                    avatar=fake.image_url())
        users.append(user)
        sql_str += f"('{user.id}', '{user.email}', '{user.password}', '{user.name}', "\
                   f"'{user.surname}', '{user.phone}', '{user.city}', '{user.avatar}'),\n"

    sql_str = sql_str.strip(',\n') + ';\n\n'
    output_file.write(sql_str)

    # Generate Schools
    sql_str = ''
    sql_str += "insert into public.school (id, name, description, owner_id) values "
    schools = []

    for i in range(0, school_count):
        school = School(id=uuid.uuid4(),
                        name=fake.company(),
                        description=fake.text(),
                        owner_id=random.choice(users).id)
        schools.append(school)
        sql_str += f"('{school.id}', '{school.name}', '{school.description}', '{school.owner_id}'),\n"

    sql_str = sql_str.strip(',\n') + ';\n\n'
    output_file.write(sql_str)

    # Generate Courses
    sql_str = ''
    sql_str += "insert into public.course (id, name, school_id, level, price, language, status) values "
    courses = []
    statuses = ['draft']

    for i in range(0, course_count):
        course = Course(id=uuid.uuid4(),
                        name=fake.job().replace('\'', ''),
                        school_id=random.choice(schools).id,
                        level=random.randint(1, 5),
                        price=random.randint(100, 1000),
                        language=fake.language_name(),
                        status=random.choice(statuses))
        courses.append(course)
        sql_str += f"('{course.id}', '{course.name}', '{course.school_id}', {course.level}, "\
                   f"{course.price}, '{course.language}', '{course.status}'),\n"

    sql_str = sql_str.strip(',\n') + ';\n\n'
    output_file.write(sql_str)

    # Generate Lessons
    sql_str = ''
    sql_str += "insert into public.lesson (id, title, type, score, theory_url, video_url, course_id) values "
    lessons = []
    lesson_types = ['theory', 'video', 'practice']

    # Ensure each course has at least one theory or video lesson and one practice lesson
    for course in courses:
        course_lessons = []
        # Add at least one theory or video lesson
        theory_type = random.choice(['theory', 'video'])
        theory_lesson = Lesson(id=uuid.uuid4(),
                               title=fake.sentence(),
                               type=theory_type,
                               score=random.randint(1, 100),
                               theory_url=fake.url() if theory_type == 'theory' else '',
                               video_url=fake.url() if theory_type == 'video' else '',
                               course_id=course.id)
        lessons.append(theory_lesson)
        course_lessons.append(theory_lesson)
        sql_str += f"('{theory_lesson.id}', '{theory_lesson.title}', '{theory_lesson.type}', {theory_lesson.score}, "\
                   f"'{theory_lesson.theory_url}', '{theory_lesson.video_url}', '{theory_lesson.course_id}'),\n"

        # Add at least one practice lesson
        practice_lesson = Lesson(id=uuid.uuid4(),
                                 title=fake.sentence(),
                                 type='practice',
                                 score=random.randint(1, 100),
                                 theory_url='',
                                 video_url='',
                                 course_id=course.id)
        lessons.append(practice_lesson)
        course_lessons.append(practice_lesson)
        sql_str += f"('{practice_lesson.id}', '{practice_lesson.title}', '{practice_lesson.type}', {practice_lesson.score}, "\
                   f"'{practice_lesson.theory_url}', '{practice_lesson.video_url}', '{practice_lesson.course_id}'),\n"

        # Generate additional lessons to fill up the lesson count
        for i in range(2, random.randint(2, lesson_count // course_count)):
            lesson_type = random.choice(lesson_types)
            lesson = Lesson(id=uuid.uuid4(),
                            title=fake.sentence(),
                            type=lesson_type,
                            score=random.randint(1, 100),
                            theory_url=fake.url() if lesson_type == 'theory' else '',
                            video_url=fake.url() if lesson_type == 'video' else '',
                            course_id=course.id)
            lessons.append(lesson)
            course_lessons.append(lesson)
            sql_str += f"('{lesson.id}', '{lesson.title}', '{lesson.type}', {lesson.score}, "\
                       f"'{lesson.theory_url}', '{lesson.video_url}', '{lesson.course_id}'),\n"

    sql_str = sql_str.strip(',\n') + ';\n\n'
    output_file.write(sql_str)

    # Generate Tests for practice lessons
    sql_str = ''
    sql_str += "insert into public.test (id, task_url, options, answer, score, level, lesson_id) values "
    tests = []

    practice_lessons = [lesson for lesson in lessons if lesson.type == 'practice']

    for lesson in practice_lessons:
        for i in range(0, test_count):
            test = Test(id=uuid.uuid4(),
                        task_url=fake.url(),
                        options=f"{fake.word()}, {fake.word()}, {fake.word()}",
                        answer=fake.word(),
                        score=random.randint(1, 100),
                        level=random.randint(1, 5),
                        lesson_id=lesson.id)
            tests.append(test)
            sql_str += f"('{test.id}', '{test.task_url}', '{test.options}', '{test.answer}', "\
                       f"{test.score}, {test.level}, '{test.lesson_id}'),\n"

    sql_str = sql_str.strip(',\n') + ';\n\n'
    output_file.write(sql_str)


if __name__ == "__main__":
    # Set up argument parser
    parser = argparse.ArgumentParser(description='Generate SQL data for database tables.')
    parser.add_argument('--users', type=int, default=10, help='Number of users to generate')
    parser.add_argument('--schools', type=int, default=5, help='Number of schools to generate')
    parser.add_argument('--courses', type=int, default=10, help='Number of courses to generate')
    parser.add_argument('--lessons', type=int, default=20, help='Number of lessons to generate')
    parser.add_argument('--tests', type=int, default=50, help='Number of tests to generate')
    parser.add_argument('--output', type=str, default='output.sql', help='Output SQL file')

    # Parse the arguments
    args = parser.parse_args()

    # Open output file
    with open(args.output, 'w') as file:
        generate_data(args.users, args.schools, args.courses, args.lessons, args.tests, file)