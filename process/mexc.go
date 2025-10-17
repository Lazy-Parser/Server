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
)

type ExchangeMexc struct {
	Process

	mexc   exchange.Exchange
	spotCh chan *market.MexcSpotTick
	doneCh chan struct{}
}

func NewMexcProcess(title string) *ExchangeMexc {
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
	if err := mexc.mexc.ListenSpot(ctx, mexc.spotCh); err != nil {
		log.Printf("listen spot error: %v", err)
	}
}

func (mexc *ExchangeMexc) StopProcess(cancel context.CancelFunc) {
	cancel()
	<-mexc.doneCh
	close(mexc.spotCh)
}

//func (mexc *Mexc) Run() error {
//}
