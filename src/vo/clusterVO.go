package vo

import "time"

type ClusterVO struct {
	ClusterName     string `gorm:"uniqueIndex" json:"clusterName"`
	Description     string `json:"description"`
	SteamCmd        string `json:"steamcmd"`
	ForceInstallDir string `json:"force_install_dir"`
	Backup          string `json:"backup"`
	ModDownloadPath string `json:"mod_download_path"`
	Uuid            string `json:"uuid"`
	Beta            bool   `json:"beta"`
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Master          bool `json:"master"`
	Caves           bool `json:"caves"`

	Connected      int    `json:"connected"`
	MaxConnections int    `json:"maxConnections"`
	Mode           string `json:"mode"`
	Mods           int    `json:"mods"`
	Season         string `json:"season"`
	Password       string `json:"password"`
	RowId          string `json:"rowId"`
	Region         string `json:"region"`
}
