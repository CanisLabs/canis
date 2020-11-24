package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/makiuchi-d/gozxing"
	qr "github.com/makiuchi-d/gozxing/qrcode"
	"github.com/skip2/go-qrcode"
)

func main() {
	encode()
	//decode()
}

func encode() {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:7779/agents/test123/invitation/xyz987", nil)
	if err != nil {
		log.Println("error creating request", err)
		return
	}
	req.Header.Add("X-API-Key", "D3YYdahdgC7VZeJwP4rhZcozCRHsqQT3VKxK9hTc2Yoh")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("failed to execute request", err)
		return
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("failed to close body", err)
		}
	}()

	m := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		log.Println("error decoding map", err)
		return
	}

	log.Println(m)

	b := m["Invitation"].(string)

	ci := base64.URLEncoding.EncodeToString([]byte(b))
	str := fmt.Sprintf("http://192.168.86.30/?c_i=%s", ci)

	fmt.Println("poop", ci)

	fname := "./invite.png"
	err = qrcode.WriteFile(str, qrcode.Medium, 256, fname)
	if err != nil {
		log.Fatal(err)
	}

}

func decode() {
	imgdata, err := ioutil.ReadFile("./verity-invite.png")
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		log.Fatalln(err)
	}
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qr.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)
}
