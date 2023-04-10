package evant_consumer

import (
	"log"
	"time"

	"telegram_bot_link/events"
)


type Consumer struct {
	fetcher    events.Fatcher
	processor  events.Processor
	batchSize  int
}

func New(fetcher events.Fatcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error{
	for {
		gotEvents, err := c.fetcher.Fatch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s",err.Error())
			continue
		}

		if len(gotEvents)== 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handEvents(gotEvents); err != nil {
			log.Print(err)
		}
	}
}


//1.потеря событий :ретраи, возращениее в хранилище,
                //   фоллбэк(можно в опер.память но при перезапуске утрата),
				//   потверждение для Fatch 
//2. обработка всей пачки:оставливаться после первой ошибке, счетчик ошибок
//3 Параллейная обрабока(sunc.WaitGroup)
func(c *Consumer) handEvents(events []events.Event) error {
	for _, event := range events{
		log.Printf("got new event: %s", event.Text)
		if err := c.processor.Process(event); err != nil{
			log.Printf("can't handle event: %s", err.Error())
			continue
		}
	}
	return nil
}