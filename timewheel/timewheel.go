package timewheel

import (
	gotimewheel "github.com/rfyiamcool/go-timewheel"
	"time"
)

const (
	DefaultTick         = 100 * time.Millisecond
	DefaultBucketNumber = 200
)

type Service struct {
	tick         time.Duration
	bucketNumber int
	tw           *gotimewheel.TimeWheel
}

type Option func(*Service)

func WithTick(tick time.Duration) Option {
	return func(service *Service) {
		service.tick = tick
	}
}

func WithBucketNumber(num int) Option {
	return func(service *Service) {
		service.bucketNumber = num
	}
}

func New(opts ...Option) *Service {
	svc := &Service{
		tick:         DefaultTick,
		bucketNumber: DefaultBucketNumber,
	}
	for _, opt := range opts {
		opt(svc)
	}

	tw, err := gotimewheel.NewTimeWheel(svc.tick, svc.bucketNumber)
	if err != nil {
		panic(err)
	}
	svc.tw = tw
	return svc
}

func (s *Service) Start() error {
	s.tw.Start()
	return nil
}

func (s *Service) Stop() {
	s.tw.Stop()
}

func (s *Service) Wait(timeout time.Duration) {
	return
}

type Task interface {
	Run()
}

type TaskHandler func()

func (h TaskHandler) Run() {
	h()
}

var _ Task = TaskHandler(nil)

func (s *Service) Add(delay time.Duration, task Task) {
	s.tw.Add(delay, func() {
		task.Run()
	})
}

func (s *Service) AddCron(interval time.Duration, task Task) {
	s.tw.AddCron(interval, func() {
		task.Run()
	})
}
