package usecases

func (uc *Usecases) HandleHardAccel() {
	// if accel < -x
	uc.SendDeccelAlert()

	// if accel > x
	uc.MarkHardAccel()
}

func (uc *Usecases) SendDeccelAlert() {
	// Send alert for hard braking event
	// record in postgres
}

func (uc *Usecases) MarkHardAccel() {

}
