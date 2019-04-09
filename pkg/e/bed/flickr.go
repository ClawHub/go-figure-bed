package bed

import "encoding/xml"

type FlickrGetPicResp struct {
	XMLName        xml.Name `xml:"photo"`
	Id             string   `xml:"id,attr"`
	Secret         string   `xml:"secret,attr"`
	Server         string   `xml:"server,attr"`
	Farm           string   `xml:"farm,attr"`
	Dateuploaded   string   `xml:"dateuploaded,attr"`
	Originalsecret string   `xml:"originalsecret,attr"`
	Originalformat string   `xml:"originalformat,attr"`
}
