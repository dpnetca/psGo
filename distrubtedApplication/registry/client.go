package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Resgistration) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service resonded with code %v", res.StatusCode)
	}
	return nil
}

func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL, bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service. Registry service responded with code %v", res.StatusCode)
	}
	return err
}
