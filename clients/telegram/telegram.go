package telegram

import (
    "encoding/json"
    "io"
    "net/http"
    "net/url"
    "path"
    "strconv"
    "telegram_bot_link/lib/e"
)

const (
    getUpdatesMethod   = "getUpdates"
    sendMessageMethod  = "sendMessage"
)

type Client struct {
    host     string
    basePath string
    client   http.Client
}

// Функция создает новый экземпляр клиента Telegram API.
// Она принимает два параметра - адрес хоста и токен API.
func New(host string, token string) *Client {
    return &Client{
        host:     host,
        basePath: newBasePath(token),
        client:   http.Client{},
    }
}

// Функция возвращает базовый путь API для клиента на основе токена.
func newBasePath(token string) string {
    return "bot" + token
}

// Метод, который получает обновления чата из Telegram API с заданным сдвигом (offset) и лимитом (limit).
// Он возвращает список обновлений и ошибку.
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
    q := url.Values{}
    q.Add("offset", strconv.Itoa(offset))
    q.Add("limit", strconv.Itoa(limit))

    data, err := c.doRequest(getUpdatesMethod, q)
    if err != nil {
        return nil, err
    }

    var res UpdatesResponse
    if err := json.Unmarshal(data, &res); err != nil {
        return nil, err
    }

    return res.Result, nil
}

// Метод отправки сообщения в Телеграмм.
// Он принимает идентификатор чата и текст сообщения в качестве параметров и возвращает ошибку.
func (c *Client) SendMessage(chatID int, text string) error {
    q := url.Values{}
    q.Add("chat_id", strconv.Itoa(chatID))
    q.Add("text", text)
    _, err := c.doRequest(sendMessageMethod, q)

    if err != nil {
        return e.Wrap("can't sent message", err)
    }

    return nil
}

// Метод для выполнения запроса к Telegram API.
// Он принимает название метода и параметры запроса, и возвращает ответ в виде байтов и ошибку.
func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
    defer func() {
        err = e.WrapIfErr("can't do request", err)
    }()

    u := url.URL{
        Scheme: "https",
        Host:   c.host,
        Path:   path.Join(c.basePath, method),
    }

    req, err := http.NewRequest(http.MethodGet, u.String(), nil)

    if err != nil {
        return nil, e.Wrap("error creating request", err)
    }
    req.URL.RawQuery = query.Encode()

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, e.Wrap("error sending request", err)
    }

    defer func() { _ = resp.Body.Close() }()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, e.Wrap("error reading response body", err)
    }
    return body, nil
}

