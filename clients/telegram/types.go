package telegram
// UpdatesResponse имеет два поля: Ok, которое является булевой переменной, 
// указывающей на успешность выполнения запроса, и Result, 
// который представляет собой массив объектов типа Update.

// Update, в свою очередь,
// представляет собой единичное обновление в чате Telegram и имеет два поля:
// ID - уникальный идентификатор обновления,
// и Message - текст сообщения.

type UpdatesResponse struct {
    Ok      bool   `json:"ok"`
    Result []Update `json:"result"`
}
type Update struct {
    ID      int    `json:"update_id"`
    Message string `json:"message"`
}
