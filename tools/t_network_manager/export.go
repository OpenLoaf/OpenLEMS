package t_network_manager

import (
	"sync"
	"t_network_manager/internal"
	"t_network_manager/public"
)

var (
	networkManagerInstance public.NetworkManager
	once                   sync.Once
)

func GetInstance() public.NetworkManager {
	once.Do(func() {
		networkManagerInstance = internal.NewSManager()
	})
	return networkManagerInstance
}
