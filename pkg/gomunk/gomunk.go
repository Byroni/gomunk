package gomunk

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type goMunk struct {
	fileStoreHandler FileStoreHandler
}

func init() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal("Could not get executable directory: ", err)
	}

	exePath := filepath.Dir(exe)

	viper.AddConfigPath(exePath)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err, "Defaulting to project root directly")

		// Look for config file in project root
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s", err)
			os.Exit(1)
		}
	}

	// Set defaults
	viper.SetDefault("AWS_REGION", "us-east-1")
	viper.SetDefault("AWS_BUCKET", "gomunk-file-store")

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to parse configuration, %v", err)
		os.Exit(1)
	}
}

func GoMunk() (*goMunk, error) {
	switch Config.PROVIDER {
	case "aws":
		return &goMunk{
			fileStoreHandler: AWSFileStore(),
		}, nil
	default:
		return &goMunk{}, errors.New("must specify a valid cloud provider. Check if you have PROVIDER set in your configuration file")
	}
}

func (gm *goMunk) ListFiles() {
	gm.fileStoreHandler.ListFiles()
}

func (gm *goMunk) UploadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	key := filepath.Base(path)

	fmt.Printf("Uploading file: %s...\n", key)

	err = gm.fileStoreHandler.UploadFile(key, file)
	if err != nil {
		return err
	}

	fmt.Println("Uploaded file to file store")
	return nil
}
