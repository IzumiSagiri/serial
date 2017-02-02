package serial

import "log"

func SoundController(status Status) {
	switch status.TagCounter {
	default:
		log.Fatal("Something went wrong.")
	case 0:
		{
			playStart()
		}
	case 1, 2, 3, 4, 5, 6, 7:
	case 9, 12, 15:
		if status.PlayerCounter != status.TagCounter {
			play60WhiteNoise()
			status.PlayerCounter = status.TagCounter
		}
	case 18:
		if status.PlayerCounter != status.TagCounter {
			play60DiNoise()
			status.PlayerCounter = status.TagCounter
		}
	case 21, 24, 27:
		if status.PlayerCounter != status.TagCounter {
			play100WhiteNoise()
			status.PlayerCounter = status.TagCounter
		}
	case 30:
		playEnd()
	case 8, 10, 11, 13, 14, 16, 17, 19, 20, 22, 23, 25, 26, 28, 29:
		play20WhiteNoise()
	}
}
