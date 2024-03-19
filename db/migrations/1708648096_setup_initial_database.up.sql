BEGIN;

CREATE TABLE IF NOT EXISTS console(
   id uuid PRIMARY KEY,
   "name" VARCHAR (50) NOT NULL UNIQUE,
   "manufacturer" VARCHAR (50),
   release_date TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

INSERT INTO console (id, name, manufacturer, release_date)
VALUES ('0eee8295-9d2e-4c19-af43-1e5464c64eb6'::uuid, 'PlayStation 5', 'Sony', '2020-11-12'),
       ('f01a57ef-9d41-4c9c-908a-abc645594cef'::uuid, 'Xbox Series X', 'Microsoft', '2020-11-10'),
       ('e008b5b0-fb6b-44f9-bdb3-3126415e7125'::uuid, 'Nintendo Switch', 'Nintendo', '2017-03-03'),
       ('b930b4c9-7dc5-4f8e-9480-7479029fa9e6'::uuid, 'PlayStation 4', 'Sony', '2013-11-15'),
       ('8558bf1b-5fc8-415a-9179-f2dfb8c6d275'::uuid, 'Xbox One', 'Microsoft', '2013-11-22'),
       ('2d8b6353-0cb3-4f2e-8569-69230eeb6199'::uuid, 'Wii U', 'Nintendo', '2012-11-18'),
       ('2ccedd83-3dc2-4c1d-a868-9b4398344a90'::uuid, 'PlayStation 3', 'Sony', '2006-11-11'),
       ('b171ae30-2d02-4da2-98b4-33ad2c331669'::uuid, 'Xbox 360', 'Microsoft', '2005-11-22'),
       ('7fbfca72-586d-4e40-b2c0-fdcc174a4d90'::uuid, 'Wii', 'Nintendo', '2006-11-19');

COMMIT;
