package fdfs_client

import (
	"testing"
	"sync"
	"time"
	"fmt"
)

func TestUpload(t *testing.T) {
	client, err := NewClientWithConfig("fdfs.conf")
	if err != nil {
		fmt.Println(err.Error())
		if client != nil {
			client.Destory()
        }
		return
    }
	defer client.Destory()
	fileId, err := client.UploadByFilename("client_test.go")
	if err != nil {
		fmt.Println(err.Error())
		client.Destory()
		return
    }
	fmt.Println(fileId.GroupName + "/" + fileId.RemoteFileName)
	if err := client.DownloadByFileId(fileId.GroupName + "/" + fileId.RemoteFileName,"tempFile");err != nil {
		fmt.Println(err.Error())
		client.Destory()
		return
    }
}

func TestUpload100(t *testing.T) {
	return
	client, err := NewClientWithConfig("fdfs.conf")
	if err != nil {
		fmt.Println(err.Error())
		if client != nil {
			client.Destory()
        }
		return
    }
	defer client.Destory()
	var wg sync.WaitGroup
	for i := 0;i != 100;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0;j != 10;j++ {
				fileId, err := client.UploadByFilename("client_test.go")
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(fileId.GroupName + "/" + fileId.RemoteFileName)
            }
		}()
    }
	time.Sleep(time.Second * 200)
	wg.Wait()	
}
