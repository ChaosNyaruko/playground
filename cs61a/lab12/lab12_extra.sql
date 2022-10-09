.read lab12.sql

-- Q5
CREATE TABLE greatstudents AS
  SELECT a.date, a.color, a.pet, a.number, b.number
  FROM students as a, fa17students as b
  WHERE a.date = b.date AND a.color = b.color AND a.pet = b.pet;

-- Q6
CREATE TABLE sevens AS
  SELECT seven
  FROM students as a, checkboxes as b
  WHERE a.time = b.time AND a.number = 7 AND b.'7' = 'True';

-- Q7
CREATE TABLE fa17favnum AS
  SELECT s.number, count(*) as c
  FROM fa17students as s
  GROUP BY s.number
  ORDER BY c desc
  LIMIT 1;
 

CREATE TABLE fa17favpets AS
  SELECT s.pet, count(*) as c
  FROM fa17students as s
  GROUP BY s.pet
  ORDER BY c desc, s.pet
  LIMIT 10;


CREATE TABLE sp18favpets AS
  SELECT s.pet, count(*) as c
  FROM students as s
  GROUP BY s.pet
  ORDER BY c desc, s.pet
  LIMIT 10;


CREATE TABLE sp18dog AS
  SELECT s.pet, count(*) as c
  FROM students as s
  WHERE s.pet = 'dog';


CREATE TABLE sp18alldogs AS
  SELECT s.pet, count(*) as c
  FROM students as s
  WHERE s.pet LIKE '%dog%';

CREATE TABLE obedienceimages AS
  SELECT s.seven, s.denero, count(*) as c
  FROM students as s
  WHERE s.seven = '7'
  GROUP BY s.denero;

-- Q8
CREATE TABLE smallest_int_count AS
  SELECT s.smallest, count(*) as c
  FROM students as s
  GROUP BY s.smallest;

