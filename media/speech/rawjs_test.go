package speech_test

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/media/speech"
)

func ExampleSpeechRecognitionEvent_Emma() {
	var event *speech.SpeechRecognitionEvent

	// cast to correct type
	value := webapi.DocumentFromJS(event.Emma())

	// do something with value
	_ = value
}
