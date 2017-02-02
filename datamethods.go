package serial

import "fmt"
import "sync"
import "log"

type SCR struct {
	SCR byte
	SCRCalm,
	SCRSound int
	Port Port
	Mux  sync.Mutex
	Cond sync.Cond
	Changed,
	Accessable bool
}

type Asahikasei struct {
	SCL,
	SCLSum07,
	SCLSum27,
	SCLSum01,
	SCLSum12,
	SCLSum23,
	SCLSum34,
	SCLSum45,
	SCLSum56,
	SCLSum67,
	SCLSumSound float64
	SCLCounterSum07,
	SCLCounterSum27,
	SCLCounterSum01,
	SCLCounterSum12,
	SCLCounterSum23,
	SCLCounterSum34,
	SCLCounterSum45,
	SCLCounterSum56,
	SCLCounterSum67,
	SCLCounterSumSound int

	Level int
	Ead1_h,
	Ead1_l,
	Ead2_h,
	Ead2_l byte
	Alg11,
	Alg12,
	Alg13,
	Alg14 int
	Port Port
	Mux  sync.Mutex
	Cond sync.Cond
	Changed,
	Accessable bool
}

func (scr *SCR) StartFlushing() {
	if scr.Accessable {
		go func() {
			for {
				SCR := scr.Port.GetSCR()
				scr.Mux.Lock()
				scr.SCR = SCR
				scr.Mux.Unlock()
				scr.Cond.L.Lock()
				scr.Changed = true
				scr.Cond.Signal()
				scr.Cond.L.Unlock()
			}
		}()
	}
}

func (asahikasei *Asahikasei) StartFlushing() {
	if asahikasei.Accessable {
		go func() {
			for {
				ead1_h, ead1_l, ead2_h, ead2_l, alg11, alg12, alg13, alg14, fSuccess := asahikasei.Port.GetAsahikasei()
				if fSuccess != 0 {
					continue
				} else {
					asahikasei.Mux.Lock()
					asahikasei.Ead1_h = ead1_h
					asahikasei.Ead1_l = ead1_l
					asahikasei.Ead2_h = ead2_h
					asahikasei.Ead2_l = ead2_l
					asahikasei.Alg11 = alg11
					asahikasei.Alg12 = alg12
					asahikasei.Alg13 = alg13
					asahikasei.Alg14 = alg14
					asahikasei.Mux.Unlock()
					asahikasei.Cond.L.Lock()
					asahikasei.Changed = true
					asahikasei.Cond.Signal()
					asahikasei.Cond.L.Unlock()
				}
			}
		}()
	}
}

func CalculateFromEADs(ead1_h, ead1_l, ead2_h, ead2_l byte) float64 {
	return (1.0 / 300000.0) * (((float64(ead1_h)*256.0 + float64(ead1_l)) - (float64(ead2_h)*256.0 + float64(ead2_l))) / (float64(0xFFFF) - (float64(ead1_h)*256.0 + float64(ead1_l)))) * 1000000.0
}

func calculateTemperature(ead4_h byte) float64 {
	return 21 + float64(ead4_h)*0.0667
}

func CalculateLevel(alg11, alg12, alg13, alg14 int) int {
	var result int
	switch {
	case alg14 != 0:
		switch {
		case alg14 == alg13 && alg13 == alg12 && alg12 == alg11 && alg11 != 0:
			result = 4 * alg14
		default:
			log.Fatal(fmt.Sprint("Invalid level data.", alg11, alg12, alg13, alg14))
		}
	case alg13 != 0:
		switch {
		case alg13 == alg12 && alg12 == alg11 && alg11 != 0:
			result = 3 * alg13
		default:
			log.Fatal(fmt.Sprint("Invalid level data.", alg11, alg12, alg13, alg14))
		}
	case alg12 != 0:
		switch {
		case alg12 == alg11 && alg11 != 0:
			result = 2 * alg12
		default:
			log.Fatal(fmt.Sprint("Invalid level data.", alg11, alg12, alg13, alg14))
		}
	case alg11 != 0:
		result = alg11
	default:
		result = 0
	}
	return result
}
