create table todos (
id serial not null ,
todo varchar ,
completed boolean default false,
--createdAt date ,
primary key (id)

) ;
insert into todos(todo , completed)  VALUES ('work', false) ;