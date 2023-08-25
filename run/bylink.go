package run

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/configs"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/core"
	"github.com/Kawaii-Konnections-KK-Limited/Hayasashiken/raytest"
)

func SingByLink(Rawurl *string, Testurl *string, InputPort *int, TimeOut *int32, InIp *string, ctx context.Context, kills *chan bool) (int32, error) {

	c, err := configs.Configbuilder(Rawurl, InputPort, InIp)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var wg sync.WaitGroup
	wg.Add(1)

	go core.RunByLink(&wg, &c, ctx, kills)
	wg.Wait()
	for {
		select {
		case <-ctx.Done():

			return 0, nil

		default:
			res, err := raytest.GetTest(InputPort, Testurl, TimeOut)
			if err != nil {
				fmt.Println(err)
				return 0, err
			}
			return res, nil

		}
	}

}
func SingByLinkProxy(Rawurl *string, Testurl *string, InputPort *int, TimeOut *int32, InIp *string, ctx context.Context, kills *chan bool) (int32, error) {

	c, err := configs.Configbuilder(Rawurl, InputPort, InIp)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	instanceReady := make(chan bool, 1)
	instanceFailed := make(chan bool, 1)

	if ctx == nil || instanceReady == nil || c == nil {
		fmt.Println("Context is nil")
	}

	go core.RunByLinkProxy(&instanceReady, &c, ctx, kills, &instanceFailed)
	// wg.Wait()
	for {

		select {
		case <-ctx.Done():
			fmt.Println("Context is done2")
			return 0, nil

		default:
			select {
			case <-instanceReady:

				select {
				case <-instanceFailed:
					return 0, nil

				default:
					res, err := raytest.GetTest(InputPort, Testurl, TimeOut)
					if err != nil {

						return 0, err
					}

					return res, nil
				}

			default:

				fmt.Println("Instance is not ready")
				time.Sleep(1 * time.Second)

			}
		}
	}

}
