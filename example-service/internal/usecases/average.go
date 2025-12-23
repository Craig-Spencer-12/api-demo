package usecases

import (
	"common/dto"
	"fmt"
)

func (uc *Usecases) CalculateAverageTruckSpeeds() {
	truckIDs, err := uc.RedisDB.GetAllTruckIDs()
	if err != nil {
		fmt.Printf("failed to get truck IDs: %v\n", err)
		return
	}

	for _, truckID := range truckIDs {
		speeds, err := uc.RedisDB.GetTruckSpeeds(truckID)
		if err != nil {
			fmt.Printf("failed to get speeds for %s: %v\n", truckID, err)
			continue
		}

		if len(speeds) == 0 {
			continue
		}

		avg := averageSpeed(speeds)
		uc.PostgresDB.AddAverageSpeed(dto.Telemetry{
			TruckID: truckID,
			Speed:   avg,
		})
		uc.RedisDB.ClearTruckSpeeds(truckID)
	}
}

func averageSpeed(speeds []float64) float64 {
	if len(speeds) == 0 {
		return 0
	}
	sum := 0.0
	for _, s := range speeds {
		sum += s
	}
	return sum / float64(len(speeds))
}
