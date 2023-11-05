package seeder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type DataSeeder struct {
	Url    string
	Client *http.Client
}

func NewDataSeeder(url string) *DataSeeder {
	return &DataSeeder{
		Url:    url,
		Client: http.DefaultClient,
	}
}

func (d *DataSeeder) SeedProducts() {
	n := 4
	for i := 1; i <= n; i++ {
		imageName := fmt.Sprintf("shoe%d.png", i)
		uuid := d.UploadImage(imageName)
		d.addProduct(fmt.Sprintf("Shoe %d", i), 1000000+(100*i), uuid)
	}

	fmt.Printf("seed %d products successfully\n", n)
}

func (d *DataSeeder) addProduct(name string, price int, uuid string) {
	body := map[string]interface{}{
		"name":        name,
		"description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book",
		"price":       price,
		"image_url":   uuid,
	}

	bz, _ := json.Marshal(body)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/product/add", d.Url), bytes.NewReader(bz))
	if err != nil {
		panic(err)
	}

	resp, err := d.Client.Do(request)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("something wrong with add product?")
	}
}

func (d *DataSeeder) UploadImage(imageName string) string {
	imagePath := filepath.Join("seed", "products", imageName)
	buf, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", imageName)
	_, err = io.Copy(fw, buf)
	if err != nil {
		panic(err)
	}
	writer.Close()

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/image/upload", d.Url), bytes.NewReader(body.Bytes()))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response, err := d.Client.Do(request)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("something wrong??")
	}

	bz, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	resp := struct {
		UUID string `json:"uuid"`
	}{}
	if err := json.Unmarshal(bz, &resp); err != nil {
		panic(err)
	}

	return resp.UUID
}
