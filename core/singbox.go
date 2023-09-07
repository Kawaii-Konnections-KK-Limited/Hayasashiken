package core

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	box "github.com/sagernet/sing-box"
	"github.com/sagernet/sing-box/log"
	"github.com/sagernet/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

var disableColor bool

// func RunByLinkDep(wg *sync.WaitGroup, config *[]byte) error {
// 	osSignals := make(chan os.Signal, 1)
// 	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
// 	defer signal.Stop(osSignals)
// 	for {
// 		// wg,
// 		instance, err := createByLink(config)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		for {
// 			osSignal := <-osSignals
// 			if osSignal == syscall.SIGHUP {

// 				if err != nil {
// 					log.Error(E.Cause(err, "reload service"))
// 					continue
// 				}
// 			}

//				closeCtx, closed := context.WithCancel(context.Background())
//				go closeMonitor(closeCtx)
//				instance.Close()
//				closed()
//				if osSignal != syscall.SIGHUP {
//					return nil
//				}
//				break
//			}
//		}
//	}
func RunByLink(wg *sync.WaitGroup, config *[]byte, ctx context.Context, kills *chan bool) error {
	// osSignals := make(chan os.Signal, 1)
	// signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	// defer signal.Stop(osSignals)
	for {
		instance, err := createByLink(config, ctx)
		if err != nil {

			wg.Done()
			return err
		}

		if instance == nil {

			//return new error
			wg.Done()
			return errors.New("instance is nil")
		}

		for {
			wg.Done()

			select {
			case <-ctx.Done():
				// exit gracefully
				fmt.Println("Context is done3")

				// closeCtx, closed := context.WithCancel(ctx)
				// go closeMonitor(closeCtx)

				instance.Close()

				// closed()
				return nil

			case k := <-*kills:

				if k {
					fmt.Println("kill")
					// closeCtx, closed := context.WithCancel(ctx)
					// go closeMonitor(closeCtx)

					instance.Close()

					// closed()

					return nil
				}

			}
		}
	}
}
func RunByLinkProxy(r *chan bool, config *[]byte, ctx context.Context, kills *chan bool, failed *chan bool) error {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
	defer signal.Stop(osSignals)
	for {
		instance, cancel, err := createByLinkProxy(config)
		if err != nil {

			*r <- true
			*failed <- true
			return err
		}
		if cancel == nil {
			//return new error

			*r <- true
			*failed <- true
			return errors.New("cancel is nil")
		}
		if instance == nil {

			//return new error
			*r <- true
			*failed <- true
			return errors.New("instance is nil")
		}

		for {
			*r <- true

			select {
			case <-ctx.Done():
				// exit gracefully
				fmt.Println("Context is done3")

				cancel()

				closeCtx, closed := context.WithCancel(ctx)
				go closeMonitor(closeCtx)

				instance.Close()

				closed()
				return nil

			case k := <-*kills:

				if k {

					closeCtx, closed := context.WithCancel(ctx)
					go closeMonitor(closeCtx)

					cancel()

					instance.Close()

					closed()

					return nil
				}

			}

		}

	}

}
func createByLinkProxy(config *[]byte) (*box.Box, context.CancelFunc, error) {

	var options option.Options
	err := options.UnmarshalJSON(*config)

	if err != nil {
		return nil, nil, err
	}
	if disableColor {
		if options.Log == nil {
			options.Log = &option.LogOptions{}
		}
		options.Log.DisableColor = true
	}
	ctx, cancel := context.WithCancel(context.Background())
	instance, err := box.New(box.Options{
		Context: ctx,
		Options: options,
	})
	if err != nil {
		cancel()
		return nil, nil, E.Cause(err, "create service")
	}
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	defer func() {
		signal.Stop(osSignals)
		close(osSignals)
	}()

	go func() {
		_, loaded := <-osSignals
		if loaded {
			cancel()
		}
	}()
	err = instance.Start()
	if err != nil {
		cancel()
		return nil, nil, E.Cause(err, "start service")
	}
	return instance, cancel, nil
}

// , context.CancelFunc
func createByLink(config *[]byte, ctx context.Context) (*box.Box, error) {

	var options option.Options
	err := options.UnmarshalJSON(*config)

	if err != nil {
		return nil, err
	}
	if disableColor {
		if options.Log == nil {
			options.Log = &option.LogOptions{}
		}
		options.Log.DisableColor = true
	}
	// ctx, cancel := context.WithCancel(context.Background())
	instance, err := box.New(box.Options{
		Context: ctx,
		Options: options,
	})
	if err != nil {
		// cancel()
		return nil, E.Cause(err, "create service")
	}
	// osSignals := make(chan os.Signal, 1)
	// signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	// defer func() {
	// 	signal.Stop(osSignals)
	// 	close(osSignals)
	// }()

	// go func() {
	// 	_, loaded := <-osSignals
	// 	if loaded {
	// 		cancel()
	// 	}
	// }()
	err = instance.Start()
	if err != nil {
		// cancel()
		return nil, E.Cause(err, "start service")
	}
	return instance, nil
}
func create(wg *sync.WaitGroup) (*box.Box, context.CancelFunc, error) {
	var (
		configContent []byte
		err           error
	)
	if err != nil {
		fmt.Printf("error")
	}
	configContent, err = os.ReadFile("config.json")
	if err != nil {
		fmt.Printf("file doesnt exist")
	}
	var options option.Options
	err = options.UnmarshalJSON(configContent)

	if err != nil {
		return nil, nil, err
	}
	if disableColor {
		if options.Log == nil {
			options.Log = &option.LogOptions{}
		}
		options.Log.DisableColor = true
	}
	ctx, cancel := context.WithCancel(context.Background())
	instance, err := box.New(box.Options{
		Context: ctx,
		Options: options,
	})
	if err != nil {
		cancel()
		return nil, nil, E.Cause(err, "create service")
	}
	wg.Done()
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer func() {
		signal.Stop(osSignals)
		close(osSignals)
	}()

	go func() {
		_, loaded := <-osSignals
		if loaded {
			cancel()
		}
	}()
	err = instance.Start()
	if err != nil {
		cancel()
		return nil, nil, E.Cause(err, "start service")
	}
	return instance, cancel, nil
}
func closeMonitor(ctx context.Context) {
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		return
	default:
	}
	log.Fatal("sing-box did not close!")
}
func Run(wg *sync.WaitGroup) error {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(osSignals)
	for {
		instance, cancel, err := create(wg)
		if err != nil {
			fmt.Println(err)
		}

		for {
			osSignal := <-osSignals
			if osSignal == syscall.SIGHUP {

				if err != nil {
					log.Error(E.Cause(err, "reload service"))
					continue
				}
			}
			cancel()
			closeCtx, closed := context.WithCancel(context.Background())
			go closeMonitor(closeCtx)
			instance.Close()
			closed()
			if osSignal != syscall.SIGHUP {
				return nil
			}
			break
		}
	}
}
