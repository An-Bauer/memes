package imgflip

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func load() error {
	url := "https://i.imgflip.com/9n6n6z.jpg"
	filename := "image.jpg"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch file: %s", resp.Status)
	}

	// Create the local file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	return err
}
