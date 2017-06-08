package metric

import (
	"context"
	"github.com/spf13/viper"
)

func HTTP(ctx context.Context) {

	b :=  viper.GetBool("core.metric.http")

	if !b {
		return
	}

	go func() {
		log.Debug("start pprof server")
		mux := http.NewServeMux()

		// register pprof handler
		mux.HandleFunc("/debug/pprof/", func(w http.ResponseWriter, r *http.Request) {
			http.DefaultServeMux.ServeHTTP(w, r)
		})

		// register metrics handler
		mux.HandleFunc("/debug/vars", metricsHandler)

		endpoint := http.ListenAndServe("localhost:6060", mux)
		log.Debug("stop pprof server: %v", endpoint)
	}


}

