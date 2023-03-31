//  Copyright 2022 Daniel Stamer, Wietse Venema

//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at

//      http://www.apache.org/licenses/LICENSE-2.0

//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package runinspect

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Port() (string, error) {
	port, ex := os.LookupEnv("PORT")
	if !ex {
		err := errors.New("PORT not declared")
		log.Printf("error: %v", err)
		return "", err
	}
	return port, nil
}

func Service() (string, error) {
	service, ex := os.LookupEnv("K_SERVICE")
	if !ex {
		err := errors.New("K_SERVICE not declared, are we running in a KNative-like environment?")
		log.Printf("error: %v", err)
		return "", err
	}
	return service, nil
}

func Revision() (string, error) {
	revision, ex := os.LookupEnv("K_REVISION")
	if !ex {
		err := errors.New("K_REVISION not declared, are we running in a KNative-like environment?")
		log.Printf("error: %v", err)
		return "", err
	}
	return revision, nil
}

func Region() (string, error) {
	bytes, err := getMetadata("computeMetadata/v1/instance/region")
	if err != nil {
		return "", err
	}

	region := strings.Split(string(bytes), "/")[3]

	return region, nil
}

func getMetadata(resource string) ([]byte, error) {
	headers := http.Header{
		"Metadata-Flavor": {"Google"},
	}

	request, err := http.NewRequest("GET", fmt.Sprintf("http://metadata.google.internal/%s", resource), nil)
	if err != nil {
		log.Printf("error: %v", err)
		return []byte{}, err
	}
	request.Header = headers

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("error: %v", err)
		return []byte{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error: %v", err)
		return []byte{}, err
	}

	return body, nil
}
