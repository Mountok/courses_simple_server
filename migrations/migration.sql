create table subjects
(
    id serial not null
        constraint subjects_pk
            primary key,
    name text not null,
    description text not null
);
create table lessons
(
    id serial not null
        constraint lessons_pk
            primary key,
    name text not null,
    description text not null,
    positions integer not null,
    subject_id integer not null
);
create table lessons_contents
(
    id serial not null
        constraint lessons_contents_pk
            primary key,
    image text default 'empty',
    content text not null,
    subject_id integer not null,
    lesson_id integer not null
);

insert into public.subjects (name, description) values ('Базы Данных. SQL','Освойте основы разработки базы дынных на языке SQL.');
insert into public.subjects (name, description) values ('Веб разработка','Основы верстки сайтов на HTML & CSS.');

insert into public.lessons (name, description,subject_id,positions)
values ('Введение','Узнайте что такое sql. Реляционные и не ряляционные базы данных.',1,1);
insert into public.lessons (name, description,subject_id,positions)
values ('HTML','Основные теги HTML',2,1);
insert into public.lessons (name, description,subject_id,positions)
values ('Создаем базу данных','Создадим первую базу данных и таблицы.',1,2);
insert into public.lessons (name, description,subject_id,positions)
values ('Добавляем стили','Не много украсим нашу страницу добавив CSS',2,2);
insert into public.lessons (name, description,subject_id,positions)
values ('Псевдоэлементы','Не много украсим нашу страницу добавив CSS',2,4);
insert into public.lessons (name, description,subject_id,positions)
values ('Создание формы','Научимся создавать форму при помощи тега <form>',2,3);

insert into public.lessons_contents (content,subject_id,lesson_id) values ('sql - это...',1,1);
insert into public.lessons_contents (content,subject_id,lesson_id) values ('html - это язык...',2,2);
insert into public.lessons_contents (content,subject_id,lesson_id) values ('для создания бд надо...',1,3);
insert into public.lessons_contents (content,subject_id,lesson_id) values ('давате украсим нашу страничку...',2,4);
insert into public.lessons_contents (content,subject_id,lesson_id) values ('<form> - это для...',2,7);
insert into public.lessons_contents (content,subject_id,lesson_id) values ('псевдоэлементы...',2,6);



-- ДАЛЕЕ ЗАПРОСЫ ОНИ НЕ ВХОДЯ В МИГРАЦИЮ
-- select * from subjects;

-- update lessons set description = 'Что такое псевдоэлементы и как их использовать' where id = 6

-- select * from lessons;

-- select id, name, description  from lessons where subject_id = 2 order by positions;

-- select * from lessons_contents;

-- select id, image, content from lessons_contents where subject_id = 2 and lesson_id = 6;





