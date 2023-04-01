package telegram
// Этот код определяет структуры данных для десериализации 
// JSON-ответа API Telegram на запрос getUpdates.
// Структура UpdatesResponse содержит поле Ok,
// которое указывает, успешно ли был выполнен запрос,
// и поле Result, которое содержит массив объектов типа Update.
// Каждый объект типа Update содержит поле ID,
// которое представляет уникальный идентификатор обновления,
// и поле Message, которое ссылается на объект типа IncomingMessage.
// Объект типа IncomingMessage содержит поле Text,
// которое содержит текст сообщения, поле From,
// которое ссылается на объект типа From с информацией об отправителе сообщения (например, его имя пользователя), и поле Chat,
// которое ссылается на объект типа Chat с информацией о чате, в котором было получено сообщение.
// Объекты типа From и Chat содержат дополнительную информацию о пользователе и чате соответственно.

type UpdatesResponse struct {
    Ok      bool   `json:"ok"`
    Result []Update `json:"result"`
}
type Update struct {
    ID      int    `json:"update_id"`
    Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct{
    Text string  `json:"text"`
    From From    `json:"from"`
    Chat Chat    `json:"chat"` 
}

type From struct {
    Username string `json:"username"`
}

type Chat struct {
    ID int `json:"id"`
}