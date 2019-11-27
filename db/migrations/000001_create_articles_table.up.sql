create table articles(
    article_id serial primary key,
    title text not null,
    auther text not null,
    link text not null,
    description text not null,
    published varchar(10),
    at_created timestamp default current_timestamp,
    at_updated timestamp
);
