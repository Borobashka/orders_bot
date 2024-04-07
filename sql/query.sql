CREATE TABLE document
(
    document_id INT NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    year integer,
    path character varying(126),
    name character varying(50), 
    author character varying(50),
    creationdate date,
    employee_id integer 
);

CREATE TABLE employee
(
    employee_id integer UNIQUE,
    name character varying,
    creationdate date,
    exhausted boolean,
    role character varying,
    phone character varying
);

ALTER TABLE document 
ADD CONSTRAINT FK_documents_employee FOREIGN KEY(employee_id) REFERENCES employee(employee_id)

DROP TABLE document

SELECT *
FROM employee;

INSERT INTO document (year, path, name, author, creationdate, employee_id)
VALUES 
		(2024, 'src\github.com\borobashka', 'о проведении экзамена', 'охотник №1', '2024-04-01', 123123),
		(2024, 'src\github.com\borobashka', 'о проведении экзамена', 'охотник №2', '2024-04-01', 123124),
		(2024, 'src\github.com\borobashka', 'о проведении экзамена', 'охотник №3', '2024-04-01', 123126);
		
INSERT INTO employee (employee_id, name, creationdate, exhausted, role, phone)
VALUES 
		(123123, 'Айзек Нетеро', '2024-04-01', false, 'Глава организации', '+791232323'),
		(123124, 'Киллуа Золдик', '2024-04-01', false, 'Охотник', '+791232323'),
		(123125, 'Куроро Люцифер', '2024-04-01', false, 'Глава Пауков', '+791232323'),
		(123126, 'Хисока Морроу', '2024-04-01', false, 'Охотник', '+791232323'),
		(123127, 'Гон Фрикс', '2024-04-01', false, 'Охотник', '+791232323');
		
TRUNCATE TABLE document  RESTART IDENTITY

UPDATE employee SET role='Изгой' WHERE employee_id = 123126

