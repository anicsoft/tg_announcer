-- +goose Up
-- Step 1: Add new columns with TIMESTAMPTZ data type
ALTER TABLE Announcements ADD COLUMN start_date_time TIMESTAMPTZ;
ALTER TABLE Announcements ADD COLUMN end_date_time TIMESTAMPTZ;

-- Step 2: Migrate data from old columns to new columns
UPDATE Announcements
SET start_date_time = start_date + start_time,
    end_date_time = CASE
                        WHEN end_time < start_time THEN end_date + end_time + INTERVAL '1 day'
    ELSE end_date + end_time
END;

-- Step 3: Drop old columns
ALTER TABLE Announcements DROP COLUMN start_date;
ALTER TABLE Announcements DROP COLUMN end_date;
ALTER TABLE Announcements DROP COLUMN start_time;
ALTER TABLE Announcements DROP COLUMN end_time;

-- +goose Down
-- Step 1: Add old columns back with their original data types
ALTER TABLE Announcements ADD COLUMN start_date DATE;
ALTER TABLE Announcements ADD COLUMN end_date DATE;
ALTER TABLE Announcements ADD COLUMN start_time TIME;
ALTER TABLE Announcements ADD COLUMN end_time TIME;

-- Step 2: Migrate data from new columns to old columns
UPDATE Announcements
SET start_date = start_date_time::DATE,
    end_date = end_date_time::DATE,
    start_time = start_date_time::TIME,
    end_time = end_date_time::TIME;

-- Step 3: Drop new columns
ALTER TABLE Announcements DROP COLUMN start_date_time;
ALTER TABLE Announcements DROP COLUMN end_date_time;
