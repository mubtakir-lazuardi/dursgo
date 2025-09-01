package scanner

import (
	"github.com/mubtakir-lazuardi/dursgo/internal/crawler"
	"github.com/mubtakir-lazuardi/dursgo/internal/httpclient"
	"github.com/mubtakir-lazuardi/dursgo/internal/logger"
)

type Scanner interface {
	Name() string
	Scan(req crawler.ParameterizedRequest, client *httpclient.Client, log *logger.Logger, opts ScannerOptions) ([]VulnerabilityResult, error)
}