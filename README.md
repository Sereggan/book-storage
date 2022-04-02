# book-storage

test task Задание: Спроектировать базу данных, в которой содержится авторы  
книг и сами книги. Необходимо написать сервис который будет по  
автору искать книги, а по книге искать её авторов.  
Требования к сервису:

* Сервис должен принимать запрос по GRPC.
* Код сервиса должен быть хорошо откомментирован
* Код должен быть покрыт unit тестами
* В сервисе должен лежать Dockerfile, для запуска базы данных с  
  тестовыми данными
* Должна быть написана документация, как запустить сервис.  
  Плюсом будет если в документации будут указания на команды,  
  для запуска сервиса и его окружения, через Makefile
* код должен быть выложен на github.

# Result:

* Для запуска:
  * go mod tidy
  * make run-database - запуск postgresql в docker
  * make generate - генерация из .proto и моков
  * make run - запуск сервера
  * make run-test-client - запуск тестового клиента
* Комментариев в коде нет, все имена пакетов методов вполне говорят сами за себя.
* Тесты я тоже не понял на что писать, хотя сделал генерацию моков. Просто приложение, читающее из базы
