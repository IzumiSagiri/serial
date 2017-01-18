package serialportmethods

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
