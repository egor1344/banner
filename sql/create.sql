create schema public;

comment on schema public is 'standard public schema';

alter schema public owner to postgres;

create table banners
(
    id serial not null
        constraint banners_pk
            primary key
);

alter table banners
    owner to postgres;

create table soc_dem_group
(
    id   serial not null
        constraint soc_dem_group_pk
            primary key,
    name text
);

alter table soc_dem_group
    owner to postgres;

create unique index soc_dem_group_id_uindex
    on soc_dem_group (id);

create table slot
(
    id        serial not null
        constraint slot_pk
            primary key,
    name      text,
    id_banner integer
        constraint slot_banners_id_fk
            references banners
            on update cascade on delete set null
);

alter table slot
    owner to postgres;

create unique index slot_id_uindex
    on slot (id);

create table statistic
(
    id          serial  not null
        constraint statistic_pk
            primary key,
    id_banner   integer not null
        constraint statistic_banners_id_fk
            references banners
            on update cascade on delete cascade,
    id_soc_dem  integer not null
        constraint statistic_soc_dem_group_id_fk
            references soc_dem_group
            on update cascade on delete cascade,
    count_click integer not null,
    count_views integer not null,
    id_slot     integer not null
        constraint statistic_slot_id_fk
            references slot
            on update cascade on delete cascade
);

alter table statistic
    owner to postgres;

create unique index statistic_id_uindex
    on statistic (id);

