# Files Exchanger
gRPC сервис + клиент для обмена файлами. Поддерживает 3 основные операции:
- Загрузка файла на сервер
- Получение списка всех файлов
- Получение файла с сервера по имени

# Инструкция по запуску
Для запуска сервиса нужно собрать его с тегом app и запустить исполняемый файл
```
go build -tags app -o cmd/app/main.go && ./app
```

Перед запуском может понадобиться инициализировать бд
```
go build -tags db_init -o db_init cmd/db_init/main.go && ./db_init
```

## Работа с именем файла
Имя файла требует тщательной валидации. Оно может стать опасным местом приложения, которое может дать доступ ко всей файловой системе сервиса. Таким образом, может произойти перезапись системных или важных файлов, а также занесение на сервер зловредных файлов
### Решение
Вместо того, чтобы тщательно и скурпулезно настраивать валидацию имени файла на предмет расширений, запрещенных символов и наличия относительного пути, мы воспользуемся хэшированием. Этот способ поможет нам вообще не париться о там, какое имя файла нам было передано. Единственное, что нужно провалидировать - это расширение и длину файла

Это было бы идеальное и простое решение, если бы не было запроса на получение списка всех файлов. Получить имя файла назад мы не сможем. А кодирование(например, при помощи base64) не решает нашу проблему - все еще можно подобрать название файла, которое при кодировании станет вредоносным

Остается только один вариант - использовать базу данных, в которой хэшу будет сопоставляться имя файла. В качестве базы данных будем использовать SQLite

## Версионирование
- v1: Простая унарная доставка бинарных файлов при помощи bytes
- v2: Остаются прежние возможности + возможность стриминговой отправки и получения файлов
- v3: Вводятся дополнительные улучшения - добавляем коды ошибок и работу с offset для файлов 
