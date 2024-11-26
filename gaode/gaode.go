package gaode

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/qichengzx/coordtransform"
)

var key string

func Init(GKey string) error {
	key = GKey
	if key == "" {
		return errors.New("no value")
	}
	return nil
}

//105.981762,31.576942
//105.981625350929,31.576960570779

//102.162052454511,29.224182013834

func InChina(lat, lng float64) (string, string, string, string, error) {
	lon1, lat1 := coordtransform.WGS84toGCJ02(lng, lat)
	//纬度是latitude；经度是longitude
	//传入内容规则：经度在前，纬度在后，经纬度间以“,”分割，经纬度小数点后不要超过 6 位。
	loc := fmt.Sprintf("%s,%s", strconv.FormatFloat(lon1, 'f', -6, 32), strconv.FormatFloat(lat1, 'f', -6, 32))
	url := fmt.Sprintf("https://restapi.amap.com/v3/geocode/regeo?key=%s&location=%s&extensions=base", key, loc)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var response struct {
		Status    string `json:"status"`
		Regeocode struct {
			AddressComponent struct {
				Province string `json:"province"`
				City     string `json:"city"`
				District string `json:"district"`
				Township string `json:"township"`
			} `json:"addressComponent"`
		} `json:"regeocode"`
		Info string `json:"info"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", "", "", "", err
	}
	return response.Regeocode.AddressComponent.Province, response.Regeocode.AddressComponent.City, response.Regeocode.AddressComponent.District, response.Regeocode.AddressComponent.Township, nil
}
