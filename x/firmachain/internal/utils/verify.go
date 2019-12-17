package utils

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

var VerifiedUrl = []string{"ipfs.infura.io", "ipfs.firmachain.org"}

// TODO:
// KVStore에서 data 불러와 verify
func VerifyFileFromKVStore() {

}

// Todo
// Tx로 query날려 data 불러와 verify
func VerifyFileFromTx() {

}

func VerifyFile(path string, hash string) error {

	if err := VerifyUrl(path); err != nil {
		return err
	}

	if err := CheckFileHash(path, hash); err != nil {
		return err
	}

	/*
		fileName := temp[len(temp)-1]
		temp := strings.Split(path, "/")
		if err := DownloadFile(fileName, path); err != nil {
			return err
		}

		if err := GetCheckSum(fileName, hash); err != nil {
			RemoveFile(fileName)
			return err
		}
		if err := RemoveFile(fileName); err != nil {
			return err
		}
	*/

	return nil
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func VerifyUrl(path string) error {
	u, err := url.Parse(path)
	host, _, _ := net.SplitHostPort(u.Host)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, found := Find(VerifiedUrl, host)

	if !found {
		return errors.New("url is not ipfs")
	}

	return nil
}

func CheckFileHash(fileUrl string, hash string) error {
	resp, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	h := sha256.New()
	if _, err := io.Copy(h, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
		return err
	}
	CheckSum := fmt.Sprintf("%x", h.Sum(nil))

	if CheckSum != hash {
		fmt.Println("CheckSum / " + CheckSum)
		fmt.Println("Hash / " + hash)
		return errors.New("Hash does not match.")
	}
	return err

}

func GetCheckSum(filepath string, hash string) error {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		fmt.Println(err)
		return err
	}

	CheckSum := fmt.Sprintf("%x", h.Sum(nil))

	if CheckSum != hash {
		fmt.Println("CheckSum / " + CheckSum)
		fmt.Println("Hash / " + hash)
		return errors.New("Hash does not match.")
	}
	return nil
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func RemoveFile(filepath string) error {
	err := os.Remove(filepath)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
