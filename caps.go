/*
This utility uses the browsercap.ini file along with a
library to take a user-agent string and return the 
simplified capabilities. These capabilites are stolen
from the browser-capabilities npm project. We map the 
browsers to the following capabilities.

    ECMAScript 2015 (aka ES6).
    'es2015'
    ECMAScript 2016.
    'es2016'
    ECMAScript 2017.
    'es2017'
    ECMAScript 2018.
    'es2018'
    HTTP/2 Server Push.
    'push'
    Service Worker API.
    'serviceworker'
    JavaScript modules, including dynamic import and import.meta.
    'modules'
We must initialize this with a browsercap.ini file that can be downloaded
here: browsercap.org
*/
package go_prpl

import (
	"io/ioutil"
	bgo "github.com/digitalcrab/browscap_go"
	"gopkg.in/yaml.v2"
)


type CapSince struct {
	ES2015 []int `yaml:"es2015"`
	ES2016 []int `yaml:"es2016"`
	ES2017 []int `yaml:"es2017"`
	ES2018 []int `yaml:"es2018"`
	Push []int `yaml:"push"`
	ServiceWorker []int `yaml:"serviceWorker"`
	Modules []int
}

type CapMap struct {
	BrowserNames []string `yaml:"browsers"`
	Caps CapSince `yaml:"caps"`
	OSCaps CapSince `yaml:"oscaps"`
}

type BCaps struct {
	Caps, OSCaps *CapSince
}

type Caps struct {
	ES2015, ES2016, ES2017, ES2018, Push, ServiceWorker, Modules bool
	BrowserType string
}

var bCapMap map[string]*BCaps 

func uaCaps(browser *bgo.Browser, bmap *map[string]*BCaps) (*Caps, error) {

	return nil, nil
}


	func initCapMap(browserCapIni string) (*map[string]*BCaps, error) {
	if err := bgo.InitBrowsCap(browserCapIni, false); err != nil {
		return nil, err
	}
	
	// now we'll parse the config
	dat, err := ioutil.ReadFile("capmap.yaml")
	if (err != nil) {
		return nil, err
	}

	var capMap = []CapMap{}
	err = yaml.Unmarshal(dat, &capMap)
	if (err != nil) {
		return nil, err
	}

	bCapMap = make(map[string]*BCaps)

	
	for _, c := range capMap {
		bcaps := normalizeBcaps(&BCaps{ &c.Caps, &c.OSCaps })

		for _, b := range c.BrowserNames {
			bCapMap[b] = bcaps
		}
	}


	return &bCapMap, nil 
}

const MaxInt = int( (^uint(0)) >> 1 ) // https://stackoverflow.com/a/6878625
/*
Every Caps and OSCaps should have the same number of elements
and all elements should have a value. For empty values in Caps a  default
of MaxInt will be set (nothing should be bigger then that right?)
For empty values in OSCaps a zero will be set as we are going to check with
an and if caps is true and oscaps is false that would be wrong.
*/
func normalizeBcaps(bcaps *BCaps) (*BCaps) {
	if bcaps.Caps == nil {
		bcaps.Caps = &CapSince{}
	}
	if bcaps.OSCaps == nil {
		bcaps.OSCaps = &CapSince{}
	}

	bcaps.Caps = normalizeCapSince(bcaps.Caps, MaxInt)
	bcaps.OSCaps = normalizeCapSince(bcaps.OSCaps, 0)
	
	return bcaps

}

func normalizeCapSince(capSince *CapSince, defaultVal int) (*CapSince) {

	if capSince == nil {
		capSince = &CapSince{}
	}

	capSince.ES2015 = normalizeCap(capSince.ES2015, defaultVal)
	capSince.ES2016 = normalizeCap(capSince.ES2016, defaultVal)
	capSince.ES2017 = normalizeCap(capSince.ES2017, defaultVal)
	capSince.ES2018 = normalizeCap(capSince.ES2018, defaultVal)
	capSince.Push = normalizeCap(capSince.Push, defaultVal)
	capSince.ServiceWorker = normalizeCap(
		capSince.ServiceWorker, defaultVal)
	capSince.Modules = normalizeCap(capSince.Modules, defaultVal)

	return capSince
}

func normalizeCap(cap []int, defaultVal int) ([]int) {
	if cap == nil || len(cap) == 0 {
		cap = []int{ defaultVal, defaultVal }
	} else {
		if len(cap) == 1 {
			cap = []int { cap[0], 0 }
		}
	}

	return cap
}
