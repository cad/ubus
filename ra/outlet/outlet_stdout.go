package outlet

import (
	"fmt"
	"robotics.neu.edu.tr/ra27-telemetry/ra/middleware"
)

const (
	OUTLET_TYPE_STDOUT = "STDOUT"
)

type STDOUTOutlet Outlet

func MakeSTDOUTOutlet(config OutletConfig, preproccessors_c []middleware.Config) *Outlet {
	outlet := Outlet{
		Config: config,
		Type: OUTLET_TYPE_STDOUT,
		Preproccessors: middleware.GetProccessors(preproccessors_c),
		handlerStart: OutletSTDOUT_Start,
		handlerStop: OutletSTDOUT_Stop,
		handlerDrain: OutletSTDOUT_Drain,
		handlerIsRunning: OutletSTDOUT_IsRunning,
	}
	return &outlet
}

func OutletSTDOUT_Start(o *Outlet) {

}

func OutletSTDOUT_Stop(o *Outlet) {

}

func OutletSTDOUT_IsRunning(o *Outlet) bool {
	return true
}

func OutletSTDOUT_Drain(o *Outlet, message OutletMessage) {
	fmt.Printf("%+v\n", message)
}
