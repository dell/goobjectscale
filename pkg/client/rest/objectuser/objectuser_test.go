package objectuser_test

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest"
	"github.com/stretchr/testify/require"
)

var (
	r          *recorder.Recorder
	httpClient *http.Client
	clientset  *rest.ClientSet
)

func TestMain(m *testing.M) {
	var err error
	r, err = recorder.New("fixtures")
	if err != nil {
		log.Fatal(err)
	}
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		delete(i.Request.Headers, "X-SDS-AUTH-TOKEN")
		return nil
	})
	r.SetTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	httpClient = &http.Client{Transport: r}
	clientset = rest.NewClientSet(
		"https://testserver:443",
		"OSTOKEN-eyJ4NXUiOi",
		httpClient,
		false,
	)
	defer func() {
		if err := r.Stop(); err != nil {
			log.Fatal(err)
		}
	}()
	m.Run()
	os.Exit(0)
}

func TestObjectUser_List(t *testing.T) {
	data, err := clientset.ObjectUser().List(nil)
	require.NoError(t, err)
	require.Len(t, data.BlobUser, 1)
	require.Equal(t, data.BlobUser[0].UserID, "zmvodjnrbmjxagvwcxf5cg==")
	require.Equal(t, data.BlobUser[0].Namespace, "small-operator-acceptance")
}

var objectUserGetInfoTest = []struct {
	in      string
	out     *model.ObjectUserInfo
	withErr bool
}{
	{
		in: "zmvodjnrbmjxagvwcxf5cg==",
		out: &model.ObjectUserInfo{
			Namespace: "small-operator-acceptance",
			Name:      "zmvodjnrbmjxagvwcxf5cg==",
			Created:   "Mon Nov 25 09:55:38 GMT 2019",
		},
	},
	{
		in:      "non existed UID",
		out:     nil,
		withErr: true,
	},
}

func TestObjectUser_GetInfo(t *testing.T) {
	for _, tt := range objectUserGetInfoTest {
		t.Run(tt.in, func(t *testing.T) {
			data, err := clientset.ObjectUser().GetInfo(tt.in, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, data, tt.out)
		})
	}
}

var objectUserGetSecretTest = []struct {
	in      string
	out     *model.ObjectUserSecret
	withErr bool
}{
	{
		in: "zmvodjnrbmjxagvwcxf5cg==",
		out: &model.ObjectUserSecret{
			SecretKey1:    "cmowa2dxeXZrM2U1NWptdHdrdTB6a3B4YnRub3RwZHY=",
			KeyTimestamp1: "2019-11-25 09:56:32.364",
			Link: model.Link{
				HREF: "/object/secret-keys",
				Rel:  "self",
			},
		},
	},
	{
		in:      "non existed UID",
		out:     nil,
		withErr: true,
	},
}

func TestObjectUser_GetSecret(t *testing.T) {
	for _, tt := range objectUserGetSecretTest {
		t.Run(tt.in, func(t *testing.T) {
			data, err := clientset.ObjectUser().GetSecret(tt.in, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, data, tt.out)
		})
	}
}

var objectUserCreateSecretTest = []struct {
	in1     string
	in2     model.ObjectUserSecretKeyCreateReq
	out     *model.ObjectUserSecretKeyCreateRes
	withErr bool
}{
	{
		in1: "zmvodjnrbmjxagvwcxf5cg==",
		in2: model.ObjectUserSecretKeyCreateReq{
			SecretKey: "aaaaaa",
		},
		out: &model.ObjectUserSecretKeyCreateRes{
			SecretKey:          "aaaaaa",
			KeyTimeStamp:       "2020-03-12 19:45:00.333",
			KeyExpiryTimestamp: "",
		},
	},
	{
		in1:     "non existed UID",
		out:     nil,
		withErr: true,
	},
}

func TestObjectUser_CreateSecret(t *testing.T) {
	for _, tt := range objectUserCreateSecretTest {
		t.Run(tt.in1, func(t *testing.T) {
			data, err := clientset.ObjectUser().CreateSecret(tt.in1, tt.in2, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, data, tt.out)
		})
	}
}

var objectUserDeleteSecretTest = []struct {
	in1     string
	in2     model.ObjectUserSecretKeyDeleteReq
	withErr bool
}{
	{
		in1: "zmvodjnrbmjxagvwcxf5cg==",
		in2: model.ObjectUserSecretKeyDeleteReq{
			SecretKey: "aaaaaa",
		},
	},
	{
		in1: "non existed UID",
		in2: model.ObjectUserSecretKeyDeleteReq{
			SecretKey: "aaaaaa",
		},
		withErr: true,
	},
}

func TestObjectUser_DeleteSecret(t *testing.T) {
	for _, tt := range objectUserDeleteSecretTest {
		t.Run(tt.in1, func(t *testing.T) {
			err := clientset.ObjectUser().DeleteSecret(tt.in1, tt.in2, nil)
			if tt.withErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
