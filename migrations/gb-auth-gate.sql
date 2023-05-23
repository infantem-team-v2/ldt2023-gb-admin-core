create schema personal;

create table personal.geography
(
    id      serial
        primary key,
    city    varchar(256) not null,
    country varchar(256) not null
);

create table personal."user"
(
    id              serial
        primary key,
    full_name       varchar(256) not null,
    email           varchar(128) not null
        unique,
    password        varchar(512) not null,
    geo_position_id integer
        constraint user_geography_id_fk
            references geography
            on delete set null,
    job_position    varchar(256)
);

comment on column personal."user".id is 'Identificator of buisness';

comment on column personal."user".full_name is 'Full name of user';

comment on column personal."user".email is 'Email of user';

comment on column personal."user".password is 'Hashed password';

comment on column personal."user".geo_position_id is 'Foreign key to possible geopositions';

comment on column personal."user".job_position is 'Position on current job or buisness';

create table personal.business
(
    id                serial
        unique,
    inn               varchar(12)  not null
        unique,
    name              varchar(128) not null,
    economic_activity varchar(128),
    website           varchar(256),
    user_id           integer      not null
        constraint business_user_id_fk
            references "user"
            on delete cascade
);

comment on column personal.business.inn is 'ИНН как-никак';

comment on column personal.business.name is 'Name of the company';

comment on column personal.business.economic_activity is 'Type of economic activity of business';

comment on column personal.business.website is 'Website of company';

create schema service;

create table service.auth
(
    id          serial
        primary key,
    name        varchar(54)  not null
        unique,
    public_key  varchar(256) not null
        unique,
    private_key varchar(512) not null
        unique,
    url         varchar(256)
);

