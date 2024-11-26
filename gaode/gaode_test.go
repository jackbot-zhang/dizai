package gaode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/qichengzx/coordtransform"
)

func TestName(t *testing.T) {
	err := Init("3d96eda8c20e02f8cc4ea69635e2a510")
	if err != nil {
		t.Fatal(err)
	}
	//41.764566, 52.226635
	//31.210638, 105.8144
	p, c, d, town, err := InChina(31.6411, 102.0722)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p, c, d, town)
}
func TestGetLoc(t *testing.T) {
	//lon, lat := coordtransform.BD09toGCJ02(105.9114, 31.1456)
	//loc := fmt.Sprintf("%s,%s", strconv.FormatFloat(lon, 'f', -6, 32), strconv.FormatFloat(lat, 'f', -6, 32))
	//fmt.Println(loc)

	//105.922013,31.138664
	lon0 := 102.0722
	lat0 := 31.6411
	lon1, lat1 := coordtransform.WGS84toGCJ02(lon0, lat0)
	loc := fmt.Sprintf("%s,%s", strconv.FormatFloat(lon1, 'f', -6, 32), strconv.FormatFloat(lat1, 'f', -6, 32))
	fmt.Println("通用转高德", loc)
	loc1 := fmt.Sprintf("%s,%s", strconv.FormatFloat(lon0, 'f', -6, 32), strconv.FormatFloat(lat0, 'f', -6, 32))
	fmt.Println("原始", loc1)

	lon2, lat2 := coordtransform.BD09toGCJ02(lon0, lat0)
	loc2 := fmt.Sprintf("%s,%s", strconv.FormatFloat(lon2, 'f', -6, 32), strconv.FormatFloat(lat2, 'f', -6, 32))
	fmt.Println("百度转高德", loc2)

}
