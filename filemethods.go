package serial

import "os"
import "log"
import "fmt"

type FileController struct {
	FilePTR *os.File
	Name    string
}

func (f *FileController) Open() {
	var err error
	f.FilePTR, err = os.Create(f.Name)
	if err != nil {
		log.Fatal(err)
	}
}

func (f FileController) Close() {
	err := f.FilePTR.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (f FileController) Write(contents string) {
	f.FilePTR.WriteString(contents)
}

func (f FileController) OutputData(status Status, scr1, scr2 SCR, asahikasei1, asahikasei2 Asahikasei) {
	f.Write(fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v, %v\n", status.DurationPassed, status.Data[status.TagCounter].Tag, scr1.SCR, scr2.SCR, asahikasei1.SCL, asahikasei2.SCL, asahikasei1.Level, asahikasei2.Level))
}

func (f FileController) OutputSummary(sclAverageCalmRight07, sclAverageCalmLeft07, sclAverageCalmBoth07, sclAverageCalmRight27, sclAverageCalmLeft27, sclAverageCalmBoth27, sclAverageCalmRight01, sclAverageCalmLeft01, sclAverageCalmBoth01, sclAverageCalmRight12, sclAverageCalmLeft12, sclAverageCalmBoth12, sclAverageCalmRight23, sclAverageCalmLeft23, sclAverageCalmBoth23, sclAverageCalmRight34, sclAverageCalmLeft34, sclAverageCalmBoth34, sclAverageCalmRight45, sclAverageCalmLeft45, sclAverageCalmBoth45, sclAverageCalmRight56, sclAverageCalmLeft56, sclAverageCalmBoth56, sclAverageCalmRight67, sclAverageCalmLeft67, sclAverageCalmBoth67, sclAverageSoundRight, sclAverageSoundLeft, sclAverageSoundBoth float64, scrSumCalmRight, scrSumCalmLeft, scrSumCalmBoth, scrSumSoundRight, scrSumSoundLeft, scrSumSoundBoth int) {
	f.Write(fmt.Sprintf("SCL calm, Right, Left, Both\namong 7m, %v, %v,%v\n2-7, %v, %v,%v\n0-1, %v, %v, %v\n1-2, %v, %v,%v\n2-3, %v, %v,%v\n3-4, %v, %v, %v\n4-5, %v, %v, %v\n5-6, %v, %v, %v\n6-7, %v, %v, %v\nscl sound, Right, Left, Both\namong 11m, %v, %v, %v\nscr palm, Right, Left, Both\namong 7m, %v, %v, %v\nscr sound, Right, Left, Both\namong 11m, %v, %v, %v", sclAverageCalmRight07, sclAverageCalmLeft07, sclAverageCalmBoth07, sclAverageCalmRight27, sclAverageCalmLeft27, sclAverageCalmBoth27, sclAverageCalmRight01, sclAverageCalmLeft01, sclAverageCalmBoth01, sclAverageCalmRight12, sclAverageCalmLeft12, sclAverageCalmBoth12, sclAverageCalmRight23, sclAverageCalmLeft23, sclAverageCalmBoth23, sclAverageCalmRight34, sclAverageCalmLeft34, sclAverageCalmBoth34, sclAverageCalmRight45, sclAverageCalmLeft45, sclAverageCalmBoth45, sclAverageCalmRight56, sclAverageCalmLeft56, sclAverageCalmBoth56, sclAverageCalmRight67, sclAverageCalmLeft67, sclAverageCalmBoth67, sclAverageSoundRight, sclAverageSoundLeft, sclAverageSoundBoth, scrSumCalmRight, scrSumCalmLeft, scrSumCalmBoth, scrSumSoundRight, scrSumSoundLeft, scrSumSoundBoth))
}

func (f FileController) OutputAverage(scrl, scrr SCR, asahikaseil, asahikaseir Asahikasei) {
	f.OutputSummary(asahikaseir.SCLSum07/float64(asahikaseir.SCLCounterSum07), asahikaseil.SCLSum07/float64(asahikaseil.SCLCounterSum07), (asahikaseir.SCLSum07/float64(asahikaseir.SCLCounterSum07)+asahikaseil.SCLSum07/float64(asahikaseil.SCLCounterSum07))/2, asahikaseir.SCLSum27/float64(asahikaseir.SCLCounterSum27), asahikaseil.SCLSum27/float64(asahikaseil.SCLCounterSum27), (asahikaseir.SCLSum27/float64(asahikaseir.SCLCounterSum27)+asahikaseil.SCLSum27/float64(asahikaseil.SCLCounterSum27))/2, asahikaseir.SCLSum01/float64(asahikaseir.SCLCounterSum01), asahikaseil.SCLSum01/float64(asahikaseil.SCLCounterSum01), (asahikaseir.SCLSum01/float64(asahikaseir.SCLCounterSum01)+asahikaseil.SCLSum01/float64(asahikaseil.SCLCounterSum01))/2, asahikaseir.SCLSum12/float64(asahikaseir.SCLCounterSum12), asahikaseil.SCLSum12/float64(asahikaseil.SCLCounterSum12), (asahikaseir.SCLSum12/float64(asahikaseir.SCLCounterSum12)+asahikaseil.SCLSum12/float64(asahikaseil.SCLCounterSum12))/2, asahikaseir.SCLSum23/float64(asahikaseir.SCLCounterSum23), asahikaseil.SCLSum23/float64(asahikaseil.SCLCounterSum23), (asahikaseir.SCLSum23/float64(asahikaseir.SCLCounterSum23)+asahikaseil.SCLSum23/float64(asahikaseil.SCLCounterSum23))/2, asahikaseir.SCLSum34/float64(asahikaseir.SCLCounterSum34), asahikaseil.SCLSum34/float64(asahikaseil.SCLCounterSum34), (asahikaseir.SCLSum34/float64(asahikaseir.SCLCounterSum34)+asahikaseil.SCLSum34/float64(asahikaseil.SCLCounterSum34))/2, asahikaseir.SCLSum45/float64(asahikaseir.SCLCounterSum45), asahikaseil.SCLSum45/float64(asahikaseil.SCLCounterSum45), (asahikaseir.SCLSum45/float64(asahikaseir.SCLCounterSum45)+asahikaseil.SCLSum45/float64(asahikaseil.SCLCounterSum56))/2, asahikaseir.SCLSum56/float64(asahikaseir.SCLCounterSum56), asahikaseil.SCLSum56/float64(asahikaseil.SCLCounterSum56), (asahikaseir.SCLSum56/float64(asahikaseir.SCLCounterSum56)+asahikaseil.SCLSum56/float64(asahikaseil.SCLCounterSum56))/2, asahikaseir.SCLSum67/float64(asahikaseir.SCLCounterSum67), asahikaseil.SCLSum67/float64(asahikaseil.SCLCounterSum67), (asahikaseir.SCLSum67/float64(asahikaseir.SCLCounterSum67)+asahikaseil.SCLSum67/float64(asahikaseil.SCLCounterSum67))/2, asahikaseir.SCLSumSound/float64(asahikaseir.SCLCounterSumSound), asahikaseil.SCLSumSound/float64(asahikaseil.SCLCounterSumSound), (asahikaseir.SCLSumSound/float64(asahikaseir.SCLCounterSumSound)+asahikaseil.SCLSumSound/float64(asahikaseil.SCLCounterSumSound))/2, scrr.SCRCalm, scrl.SCRCalm, scrr.SCRCalm+scrl.SCRCalm, scrr.SCRSound, scrl.SCRSound, scrr.SCRSound+scrl.SCRSound)
}
