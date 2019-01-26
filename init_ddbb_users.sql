-- alter user person_master with encrypted password 'QWERTYU';

      DROP TABLE IF EXISTS person_pile;
    CREATE TABLE person_pile(
        id serial primary key not null,
    status integer NOT NULL default 0,
CRM_result jsonb,
 person_id integer NOT NULL,
first_name varchar(200) NOT NULL,
 last_name varchar(200) NOT NULL,
     email varchar(200) NOT NULL,
     phone varchar(20) NOT NULL,
    report varchar(1000),
created_at timestamptz not null default now(),
updated_at timestamptz not null default now());

create index idx_person_pile
    on person_pile
 using GIN(CRM_result jsonb_path_ops);

create index idx_person_pile_status
    on person_pile (status);
