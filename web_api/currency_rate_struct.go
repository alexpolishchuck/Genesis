package main

type currency_rate_struct struct {
	Success   bool
	Terms     string
	Privacy   string
	Timestamp int
	Target    string
	Rates     struct {
		BTC float64
	}
}
