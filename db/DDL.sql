CREATE database "car_rentals";

CREATE TABLE "booking_types" (
  "id" SERIAL NOT NULL,
  "name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

CREATE TABLE "cars" (
  "id" SERIAL NOT NULL,
  "name" varchar NOT NULL,
  "stock" int4 NOT NULL DEFAULT 1,
  "daily_rent_price" numeric(10,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

CREATE TABLE "drivers" (
  "id" SERIAL NOT NULL,
  "name" varchar NOT NULL,
  "nik" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "daily_cost" numeric(10,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

CREATE TABLE "memberships" (
  "id" SERIAL NOT NULL,
  "name" varchar NOT NULL,
  "discount" numeric(10,2) NOT NULL,
  "discount_type" varchar NOT NULL DEFAULT ('percentage'::charactervarying),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

CREATE TABLE "customers" (
  "id" SERIAL NOT NULL,
  "name" varchar NOT NULL,
  "nik" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "membership_id" int4,
  "membership_name" varchar,
  "membership_applied_at" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

CREATE TABLE "bookings" (
  "id" SERIAL NOT NULL,
  "customer_id" int4 NOT NULL,
  "customer_name" varchar NOT NULL,
  "car_id" int4 NOT NULL,
  "car_name" varchar NOT NULL,
  "car_daily_price" numeric(10,2) NOT NULL,
  "start_rent" date NOT NULL,
  "end_rent" date NOT NULL,
  "total_cost" numeric(10,2) NOT NULL,
  "finished" int4 NOT NULL DEFAULT 0,
  "discount" numeric(10,2) NOT NULL DEFAULT 0,
  "booking_type_id" int4 NOT NULL,
  "booking_type_name" varchar NOT NULL,
  "driver_id" int4,
  "driver_name" varchar,
  "driver_daily_cost" numeric(10,2),
  "total_driver_cost" numeric(10,2),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

CREATE TABLE "driver_incentives" (
  "id" SERIAL NOT NULL,
  "booking_id" int4 NOT NULL,
  "driver_id" int4 NOT NULL,
  "driver_name" varchar NOT NULL,
  "incentive" numeric(10,2) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" varchar,
  "updated_at" timestamp,
  "updated_by" varchar,
  "deleted_at" timestamp,
  "deleted_by" varchar,
  "is_active" int4 DEFAULT 1,
  PRIMARY KEY ("id")
);

ALTER TABLE "customers" ADD CONSTRAINT "customers_membership_id_fkey" FOREIGN KEY ("membership_id") REFERENCES "memberships" ("id");

ALTER TABLE "bookings" ADD CONSTRAINT "bookings_booking_type_id_fkey" FOREIGN KEY ("booking_type_id") REFERENCES "booking_types" ("id");

ALTER TABLE "bookings" ADD CONSTRAINT "bookings_car_id_fkey" FOREIGN KEY ("car_id") REFERENCES "cars" ("id");

ALTER TABLE "bookings" ADD CONSTRAINT "bookings_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "bookings" ADD CONSTRAINT "bookings_driver_id_fkey" FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");

ALTER TABLE "driver_incentives" ADD CONSTRAINT "driver_incentives_booking_id_fkey" FOREIGN KEY ("booking_id") REFERENCES "bookings" ("id");

ALTER TABLE "driver_incentives" ADD CONSTRAINT "driver_incentives_driver_id_fkey" FOREIGN KEY ("driver_id") REFERENCES "drivers" ("id");
