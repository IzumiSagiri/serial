package serialportmethods

/*
#cgo LDFLAGS: -lwinmm
#include <stdlib.h>
#include <stdio.h>
#include <Windows.h>
#include <mmsystem.h>
#pragma comment(lib, "winmm.lib")
#pragma comment(lib, "kernel32.lib")

void stopWhateverSound()
{
	PlaySound(NULL, NULL, SND_FILENAME);
}
void playStart()
{
	PlaySound(TEXT("start.wav"), NULL, SND_FILENAME | SND_ASYNC);
}
void playEnd()
{
	PlaySound(TEXT("end.wav"), NULL, SND_FILENAME | SND_ASYNC);
}
void play20WhiteNoise()
{
	PlaySound(TEXT("white_noise.wav"), NULL, SND_FILENAME | SND_ASYNC | SND_LOOP);
}
void play60WhiteNoise()
{
	PlaySound(TEXT("white_60.wav"), NULL, SND_FILENAME | SND_ASYNC | SND_LOOP);
}
void play100WhiteNoise()
{
	PlaySound(TEXT("white_100.wav"), NULL, SND_FILENAME | SND_ASYNC | SND_LOOP);
}
void play60DiNoise()
{
	PlaySound(TEXT("di_60.wav"), NULL, SND_FILENAME | SND_ASYNC | SND_LOOP);
}
void play100DiNoise()
{
	PlaySound(TEXT("di_100.wav"), NULL, SND_FILENAME | SND_ASYNC | SND_LOOP);
}
*/
import "C"

func stopSound() {
	C.stopWhateverSound()
}

func playStart() {
	C.playStart()
}

func playEnd() {
	C.playEnd()
}

func play20WhiteNoise() {
	C.play20WhiteNoise()
}

func play60WhiteNoise() {
	C.play60WhiteNoise()
}

func play100WhiteNoise() {
	C.play100WhiteNoise()
}

func play60DiNoise() {
	C.play60DiNoise()
}

func play100DiNoise() {
	C.play100DiNoise()
}
