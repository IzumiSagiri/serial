package serialportmethods

import "io/ioutil"
import "fmt"
import "time"
import "log"
import "testing"

// Reg Testing.

var allPortsName []string
var allPortsArray []Port

func TestCalculateAverage(t *testing.T) {
    var status Status
    var scrl, scrr SCR
    var asahikaseil, asahikaseir Asahikasei
    status.TagCounter = 1
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 2
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 3
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 4
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 5
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 6
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 7
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 8
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 9
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 10
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 11
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 12
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 13
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 14
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 15
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 16
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 17
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 18
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 19
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 20
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 21
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 22
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 23
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 24
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 25
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 26
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 27
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 28
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    status.TagCounter = 29
    scrl.SCR = 1
    scrr.SCR = 2
    asahikaseil.SCL = 3
    asahikaseir.SCL = 4
    CalculateAverage(status, scrl, scrr, asahikaseil, asahikaseir)
    fmt.Println("SCRL Calm: " + scrl.SCRCalm + " SCRL Sound: " + scrl.SCRSound + "SCRR Calm: " + scrr.SCRCalm + " SCRR Sound: " + scrr.SCRSound + " SCLSum07: " + asahikaseil.SCLSum07 + " SCLSum27: " + asahikaseil.SCLSum27 +" SCLSum01: " + asahikaseil.SCLSum01 + " SCLSum12: " + asahikaseil.SCLSum12 +" SCLSum23: " + asahikaseil.SCLSum23 + " SCLSum34: " + asahikaseil.SCLSum34 +  " SCLSum45: " + asahikaseil.SCLSum45 + " SCLSum56: " + asahikaseil.SCLSum56 +" SCLSum67: " + asahikaseil.SCLSum67 + " SCLSumSum: " + asahikaseil.SCLSumSum +
func TestMain(t *testing.T) {
	allPortsName = GetAvailableComms()
	for _, v := range allPortsName {
		var port Port
		port.Name = v
		port.Baudrate = 115200
		if port.Open() {
			allPortsArray = append(allPortsArray, port)
		}
	}
}

//Port Testing.

func TestPortReading(t *testing.T) {
	for _, v := range allPortsArray {
		fmt.Println(v.Name, ":")
		for i := 0; i < 10; i++ {
			fmt.Print(v.Read())
		}
		fmt.Print("\n")
	}
}

func TestAsahikasei(t *testing.T) {
	for _, v := range allPortsArray {
		temp := checkIfPortIsAsahikasei(v)
		fmt.Println(temp)
	}
}

func TestRecognization(t *testing.T) {
	var asahikaseiPorts, SCRPorts []Port
	asahikaseiPorts, SCRPorts = RecognizeAsahikaseiAndSCRPorts(allPortsArray, asahikaseiPorts, SCRPorts)
	for _, v := range asahikaseiPorts {
		fmt.Println(v.Name)
	}
	fmt.Println("//////")
	for _, v := range SCRPorts {
		fmt.Println(v.Name)
	}
}

// Data Testing.

func TestSCRSet(t *testing.T) {
	var scr SCR
	scr.Set(0x3F)
	got := scr.SCR
	if got != 0x3F {
		t.Errorf("SCR Set value error.")
	}
}

func TestAsahikaseiSet(t *testing.T) {
	cases := []struct {
		ead1_h, ead1_l, ead2_h, ead2_l    byte
		alg11, alg12, alg13, alg14, want2 int
		want1                             float64
	}{
		{0x81, 0x09, 0x7F, 0x9A, -1, -1, -1, 0, -3, 0.03763870941275409},
	}
	for _, c := range cases {
		var asahikasei Asahikasei
		asahikasei.Set(c.ead1_h, c.ead1_l, c.ead2_h, c.ead2_l, c.alg11, c.alg12, c.alg13, c.alg14)
		got1 := asahikasei.SCL
		got2 := asahikasei.Level
		if got1 != c.want1 {
			t.Errorf("Set Error, %Q, want %q", got1, c.want1)
		}
		if got2 != c.want2 {
			t.Errorf("Set Error, %Q, want %q", got2, c.want2)
		}
	}
}

func TestCalculateFromEADs(t *testing.T) {
	cases := []struct {
		ead1_h, ead1_l, ead2_h, ead2_l byte
		want                           float64
	}{
		{0x81, 0x09, 0x7F, 0x9A, 0.03763870941275409},
	}
	for _, c := range cases {
		got := calculateFromEADs(c.ead1_h, c.ead1_l, c.ead2_h, c.ead2_l)
		if got != c.want {
			t.Errorf("calculateFromEADs(%q) == %Q, want %q", c.ead1_h, got, c.want)
		}
	}
}

func TestCalculateTemprature(t *testing.T) {
	cases := []struct {
		ead4_h byte
		want   float64
	}{
		{128, 29.537599999999998},
	}
	for _, c := range cases {
		got := calculateTemperature(c.ead4_h)
		if got != c.want {
			t.Errorf("calculateTemperature(%q) == %Q, want %q", c.ead4_h, got, c.want)
		}
	}
}

func TestCalculateLevel(t *testing.T) {
	cases := []struct {
		alg11, alg12, alg13, alg14, want int
	}{
		{1, 1, 1, 1, 4},
		{1, 1, 1, 0, 3},
		{1, 1, 0, 0, 2},
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{-1, 0, 0, 0, -1},
		{-1, -1, -1, -1, -4},
	}
	for _, c := range cases {
		got := calculateLevel(c.alg11, c.alg12, c.alg13, c.alg14)
		if got != c.want {
			t.Errorf("calculateLevel(%q) == %Q, want %q", c.alg11, got, c.want)
		}
	}
}

// File Testing.

func TestFileWriting(t *testing.T) {
	var f FileController
	f.Name = "testFile.csv"
	f.Open()
	f.Write("This is a test.")
	f.Close()
	tempFile, err := ioutil.ReadFile("testFile.csv")
	if err != nil {
		log.Fatal(err)
	}
	got := string(tempFile)
	if got != "This is a test." {
		t.Errorf("File reading got %q, want This is a test..", got)
	}
}

func TestOutputData(t *testing.T) {
	var f FileController
	var status Status
	var scr1, scr2 SCR
	var asahikasei1, asahikasei2 Asahikasei
	status.Start()
	status.DurationPassed = time.Minute
	status.TagCounter = 15
	scr1.Set(0x38)
	scr2.Set(0x4F)
	asahikasei1.Set(0x81, 0x09, 0x7F, 0x9A, -1, -1, -1, 0)
	asahikasei2.Set(0x81, 0x09, 0x7F, 0x9A, -1, -1, 0, 0)
	f.Name = "testFileData.csv"
	f.Open()
	f.OutputData(status, scr1, scr2, asahikasei1, asahikasei2)
	f.Close()
	tempFile, err := ioutil.ReadFile("testFileData.csv")
	if err != nil {
		log.Fatal(err)
	}
	want := "1m0s, 60stimu3, 56, 79, 0.03763870941275409, 0.03763870941275409, -3, -2\n"
	got := string(tempFile)
	if got != want {
		t.Errorf("File reading got %q, want %q.", got, want)
	}
}

func TestOutputSummary(t *testing.T) {
	var f FileController
	f.Name = "testFileSummary.csv"
	f.Open()
	f.OutputSummary(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	f.Close()
	tempFile, err := ioutil.ReadFile("testFileSummary.csv")
	if err != nil {
		log.Fatal(err)
	}
	want := "SCL calm, Number\namong 10m, 0\n2-10, 1\n0-1, 2\n1-2, 3\n2-3, 4\n3-4, 5\n4-5, 6\n5-6, 7\n6-7, 8\n7-8, 9\n8-9, 10\n9-10, 11\nscl sound, Number\namong 11m, 12\nscr palm, Number\namong 10m, 13\nscr sound, Number\namong 11m, 14"
	got := string(tempFile)
	if got != want {
		t.Errorf("File reading got %q, want %q.", got, want)
	}
}

// Status Testing.

func TestStatusController(t *testing.T) {
	var status Status
	SoundController(status)
	status.Start()
	fmt.Println(status.TimeStarted)
	fmt.Println(status.DurationSession)
	fmt.Println(status.Data[status.TagCounter].Tag)
	status.Set()
	fmt.Println(status.TagCounter)
	fmt.Println(status.DurationNow)
	fmt.Println(status.DurationPassed)
	fmt.Println(status.DurationLeft)
	status.Tick()
	SoundController(status)
	fmt.Println(status.DurationSession)
	fmt.Println(status.Data[status.TagCounter].Tag)
	fmt.Println(status.TagCounter)
	fmt.Println(status.DurationNow)
	fmt.Println(status.DurationPassed)
	fmt.Println(status.DurationLeft)
}
