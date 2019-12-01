create table moscow
(
    id               bigint auto_increment,
    id_ru            bigint not null,
    global_id        bigint not null,
    system_object_id text   null,
    name             text   null,
    adm_area         text   null,
    district         text   null,
    address          text   null,
    lon              text   null,
    lat              text   null,
    car_capacity     int    null,
    mode             text   null,
    id_en            bigint null,
    name_en          text   null,
    adm_area_en      text   null,
    district_en      text   null,
    address_en       text   null,
    lon_en           text   null,
    lat_en           text   null,
    car_capacity_en  int    null,
    mode_en          text   null,
    constraint moscow_pk
        primary key (id)
);

