package process

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/Lazy-Parser/Collector/api"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/exchange"
	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/Server/publisher"
)

type ExchangeMexc struct {
	Process

	pub    *publisher.Publisher
	mexc   exchange.Exchange
	spotCh chan *market.MexcSpotTick
	doneCh chan struct{}
}

func NewMexcProcess(title string, pub *publisher.Publisher) *ExchangeMexc {
	wd, _ := os.Getwd()
	cfg, err := config.NewConfig(filepath.Join(wd, ".env"))
	if err != nil {
		//return fmt.Errorf("failed to load config: %v", err)
		log.Printf("failed to load config: %v", err)
		return nil
	}

	mexcApi := api.NewMexcApi(cfg)
	m, err := exchange.NewMexc(mexcApi)
	if err != nil {
		//return fmt.Errorf("failed to create mexc: %v", err)
		log.Printf("failed to create mexc: %v", err)
		return nil
	}

	mexc := ExchangeMexc{
		pub:    pub,
		mexc:   m,
		spotCh: make(chan *market.MexcSpotTick, 4024),
		doneCh: make(chan struct{}),
	}
	mexc.initProcess(title)

	return &mexc
}

func (mexc *ExchangeMexc) SetupProcess(ctx context.Context) error {
	if err := mexc.mexc.BufferLoop(ctx); err != nil {
		return err
	}

	log.Println("mexc setup process done")

	return nil
}

// can be blocking
func (mexc *ExchangeMexc) Do(ctx context.Context) {
	defer close(mexc.doneCh)

	go func() {
		if err := mexc.mexc.ListenSpot(ctx, mexc.spotCh); err != nil {
			log.Printf("listen spot error: %v", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case tick := <-mexc.spotCh:
			if err := mexc.pub.Publish("exchange.mexc.spot", *tick); err != nil {
				log.Printf("publish spot error: %v", err)
				ctx.Done() // stop
			}
		}
	}
}

func (mexc *ExchangeMexc) StopProcess(cancel context.CancelFunc) {
	cancel()
	<-mexc.doneCh
	close(mexc.spotCh)
}

//func (mexc *Mexc) Run() error {
//}
