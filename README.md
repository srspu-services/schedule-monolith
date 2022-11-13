# Schedule MonolithFirst Bot

Бот с расписанием для бесед и личных сообщений в ВК и Телеграм

### SRSPU API
Работа бота осуществляется на следующих эндпонтах:
- Получение списка всех факультетов: https://schedule.npi-tu.ru/api/v1/faculties
- Получение списка групп факультетов: https://schedule.npi-tu.ru/api/v1/faculties/<FACULTY_ID>/years/<YEAR>/groups
- Получение расписания для группы:  https://schedule.npi-tu.ru/api/v1/faculties/<FACULTY_ID>/years/<YEAR>/groups/<GROUP_ID>/schedule

### Technologies
- Message Broker (для предотвращения бана со стороны соц сетей): `RabbitMQ`
- База данных: `PostgreSQL`
- Контейнеризация: `Docker`
- ORM: `GORM`