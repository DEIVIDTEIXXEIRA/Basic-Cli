package killanddelete

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/urfave/cli"
)

// Volume representa as informações de um processo
type Volume struct {
	Nome          string
	Total         uint64
	Usado         uint64
	Disponivel    uint64
	PercentualUso float64
	Montagem      string
}

// ActionVolumes é responsável por obter informações sobre os volumes de arquivos montados
func ActionVolumes(c *cli.Context) error {
	stats, err := disk.Partitions(true)
	if err != nil {
		return err
	}

	var volumes []*Volume

	for _, stat := range stats {
		uso, err := disk.Usage(stat.Mountpoint)
		if err != nil {
			continue
		}

		volume := &Volume{
			Nome:          stat.Device,
			Total:         uso.Total,
			Usado:         uso.Used,
			Disponivel:    uso.Free,
			PercentualUso: uso.UsedPercent,
			Montagem:      stat.Mountpoint,
		}

		volumes = append(volumes, volume)
	}

	volumesByteArr, err := json.MarshalIndent(volumes, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(volumesByteArr))
	return nil
}
