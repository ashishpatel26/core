package v1

import (
	"strconv"

	argoprojv1alpha1 "github.com/argoproj/argo/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	"github.com/jmoiron/sqlx"
	"github.com/onepanelio/core/s3"
	"github.com/onepanelio/core/util/logging"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	artifactRepositoryEndpointKey       = "artifactRepositoryEndpoint"
	artifactRepositoryBucketKey         = "artifactRepositoryBucket"
	artifactRepositoryRegionKey         = "artifactRepositoryRegion"
	artifactRepositoryInSecureKey       = "artifactRepositoryInsecure"
	artifactRepositoryAccessKeyValueKey = "artifactRepositoryAccessKey"
	artifactRepositorySecretKeyValueKey = "artifactRepositorySecretKey"
)

type Config = rest.Config

type DB = sqlx.DB

type Client struct {
	kubernetes.Interface
	argoprojV1alpha1 argoprojv1alpha1.ArgoprojV1alpha1Interface
	*DB
}

func (c *Client) ArgoprojV1alpha1() argoprojv1alpha1.ArgoprojV1alpha1Interface {
	return c.argoprojV1alpha1
}

func NewConfig() (config *Config) {
	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(), &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		panic(err)
	}

	return
}

func NewClient(config *Config, db *sqlx.DB) (client *Client, err error) {
	config.BearerTokenFile = ""
	config.Username = ""
	config.Password = ""
	config.CertData = nil
	config.CertFile = ""

	kubeConfig, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	argoConfig, err := argoprojv1alpha1.NewForConfig(config)
	if err != nil {
		return
	}

	return &Client{Interface: kubeConfig, argoprojV1alpha1: argoConfig, DB: db}, nil
}

func (c *Client) getNamespaceConfig(namespace string) (config map[string]string, err error) {
	configMap, err := c.GetConfigMap(namespace, "onepanel")
	if err != nil {
		logging.Logger.Log.WithFields(log.Fields{
			"Namespace": namespace,
			"Error":     err.Error(),
		}).Error("getNamespaceConfig failed getting config map.")
		return
	}
	config = configMap.Data

	secret, err := c.GetSecret(namespace, "onepanel")
	if err != nil {
		logging.Logger.Log.WithFields(log.Fields{
			"Namespace": namespace,
			"Error":     err.Error(),
		}).Error("getNamespaceConfig failed getting secret.")
		return
	}

	config[artifactRepositoryAccessKeyValueKey] = string(secret.Data[artifactRepositoryAccessKeyValueKey])
	config[artifactRepositorySecretKeyValueKey] = string(secret.Data[artifactRepositorySecretKeyValueKey])

	return
}

func (c *Client) getS3Client(namespace string, config map[string]string) (s3Client *s3.Client, err error) {
	insecure, err := strconv.ParseBool(config[artifactRepositoryInSecureKey])
	if err != nil {
		logging.Logger.Log.WithFields(log.Fields{
			"Namespace": namespace,
			"ConfigMap": config,
			"Error":     err.Error(),
		}).Error("getS3Client failed when parsing bool.")
		return
	}
	s3Client, err = s3.NewClient(s3.Config{
		Endpoint:  config[artifactRepositoryEndpointKey],
		Region:    config[artifactRepositoryRegionKey],
		AccessKey: config[artifactRepositoryAccessKeyValueKey],
		SecretKey: config[artifactRepositorySecretKeyValueKey],
		InSecure:  insecure,
	})
	if err != nil {
		logging.Logger.Log.WithFields(log.Fields{
			"Namespace": namespace,
			"ConfigMap": config,
			"Error":     err.Error(),
		}).Error("getS3Client failed when initializing a new S3 client.")
		return
	}

	return
}
