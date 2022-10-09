.read fa17data.sql
.read sp18data.sql

-- Q2
CREATE TABLE obedience AS
  SELECT seven, denero from students;

-- Q3
CREATE TABLE smallest_int AS
  SELECT time, smallest from students
  WHERE smallest > 15
  ORDER BY smallest
  LIMIT 20;

-- Q4
CREATE TABLE matchmaker AS
  SELECT a.pet, a.song, a.color, b.color
  FROM students as a, students as b
  WHERE a.pet = b.pet AND a.song = b.song
        AND a.time < b.time;
