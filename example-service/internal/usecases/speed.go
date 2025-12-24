package usecases

func (uc *Usecases) SendSpeedAlert() {
	//Send Alert
	//Mark the alert in postgres
}

func (uc *Usecases) CheckForSpeeding() {
	// if not speeding
	// 		if on a speed run
	// 			record in postgres
	// 			reset counter in redis

	// else
	// 		Add to high speed counter in redis
}

func (uc *Usecases) GetSpeedReport() {
	// get all speed alerts for a certain driver in the last month from postgres
	// speed alerts and long speed events
}
