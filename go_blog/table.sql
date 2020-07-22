create table blog_tag (
    id serial primary key,
    name varchar(100),
    create_on timestamp not null,
    create_by varchar(100),
    modified_on timestamp not null,
    modified_by varchar(100),
    state boolean
);

create table blog_article (
    id serial primary key,
    tag_id integer,
    title varchar(100),
    descr varchar(255),
    content text,
    content_on timestamp not null,
    content_by varchar(100),
    modified_on timestamp not null,
    modified_by varchar(100),
    state boolean
);

create table blog_auth (
    id serial primary key,
    username varchar(50),
    password varchar(50)
);

insert into blog_auth (id, username, password) value(1,'root', '123456');