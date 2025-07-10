package employees

import (
	"github.com/Kiveri/wh-be/internal/config"
	"time"
)

type Repo struct {
	cluster *config.Cluster
	timer   timer
}

type timer interface {
	NowMoscow() time.Time
}

func NewRepo(cluster *config.Cluster, timer timer) *Repo {
	return &Repo{
		cluster: cluster,
		timer:   timer,
	}
}
