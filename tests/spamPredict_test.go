package tests

import (
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse/spamPrediction"
	"log"
	"testing"
)

func TestSpamPrediction(t *testing.T) {
	coffeehouse.SetKey("<API Key here>")
	res, err := spamPrediction.DoRequest("Hey there, I'm Sayan. How are you doing?", "en", "", "", "", "")
	if err != nil {
		log.Fatal(err.Error())
		t.Log(err)
	}
	if (*res).Success != true {
		t.Log(res)
		t.Errorf("[Intellivoid.Coffeehouse-go (emotionAnalysis)] Failed request, response code: %d", (*res).ResponseCode)

	}
}