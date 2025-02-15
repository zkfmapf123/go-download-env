package aws

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zkfmapf123/go-download-env/internal/filesystem"
)

func (a *AWSEnvParmas) IsExistBucket() (*s3.ListObjectsV2Output, bool) {
	obj, err :=a.s3Client.ListObjectsV2(
		context.TODO(),
		&s3.ListObjectsV2Input{
			Bucket: aws.String(a.GetS3Bucket()),
		},
	)

	return obj, err == nil
}


func (a *AWSEnvParmas) UpdateS3Architecture(value filesystem.ProjectSettingParams) error {

	_, isExist := a.IsExistBucket()
	if !isExist {
		return fmt.Errorf("%s bucket is not exist", a.GetS3Bucket())
	}

	// concurrency
	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	if value.IsUseCommonEnvironments {
		wg.Add(1)
		go func(pjt string) {
			defer wg.Done()
			if err := a.createFolders("common", "", ""); err != nil {
				select {
				case errChan <- err:
				default:
				}
			}
		}("common")
	}

	for pjtFolder, pjts := range value.Projects {
		
		for _, task := range pjts{
			
			for _, env := range value.Envs{
				wg.Add(1)

				go func(pjt, t, e string) {
					defer wg.Done()
					err := a.createFolders(pjt, t, e)
					if err != nil {
						select {
						case errChan <- err:
						default:
						}
					}

				}(pjtFolder,task,env)
			}
		}
	}

	done := make(chan bool)
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case err := <- errChan:
		return err
	case <- done:
		return nil
	}
}

func (a *AWSEnvParmas) createFolders(pjtFolder, task, env string) error {

	if task =="" && env == "" {
		_, err := a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(a.GetS3Bucket()),
			Key: aws.String(fmt.Sprintf("%s/", pjtFolder)),
		})	

		return err
	}
	
	_, err := a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(a.GetS3Bucket()),
		Key: aws.String(fmt.Sprintf("%s/%s/%s/", pjtFolder, task, env)),
	})

	return err
}

func (a *AWSEnvParmas) PutObject(path, key, value string ) error {
	_, err := a.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(a.GetS3Bucket()),
		Key: aws.String(fmt.Sprintf("%s/%s", path, key)),
		Body: bytes.NewReader([]byte(value)),
	})

	return err
}
