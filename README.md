# telegram_bot_link_save

  <p>Это начало проекта бота по работе с закладками. В настоящее время это очень скромный функционал - сохранение ссылок и выдача их в случайном порядке. Идея не моя (ютуб-канал: "Николай Тузов — Golang"), это первый проект на Go, целью было начать писать с базовыми знаниями (книга: Язык GO Для Начинающих М. Жашкевич). Также ChatCPT составлял документацию, так что в некоторых местах она очень подробная.</p>

```
# сборка
go build

# запуск
./telegram_bot_link -tg-bot-token 'ваш токен тг '
```
  <p>Задачи:</p>
  <ul>
    <li>Перейти на использование баз данных</li>
    <li>Реализовать напоминания (ежедневно выдавать случайную ссылку)</li>
    <li>Попробовать сохранять данные в Notion через API</li>
  </ul>