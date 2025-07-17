# Микросервис.

## Реализовано

- Сервис регистрации
  * Регистрация
  * Авторизация
  * Проверка пользователя по JWT токену
- Сервис пользователя
  * Создания пользователя
  * Изменения пользователя

## Endpoints

- POST /auth/sign_up - регистрация
  - JSON login, email, phone, password - Пока полностью заполненые поля
- POST /auth/sign_in - авторизация
  - JSON identifier, password - identifier либо логин, либо емайл, либо телефон.
- GET /user/ - информация о пользователе
  - JSON none 
- PATCH /user/ - изменяет информацию о пользователе
  - JSON (city, description, name, born_data, links)
- GET /admin/user_list - только админ, список с полной информацией пользователя