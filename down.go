package main

import (
	"github.com/techoner/gophp/serialize"
	"strings"
	"supply_warehouse/orm"
)

func main() {
	item := orm.ImsSupermanHand2Item{}
	list, err := item.SelectList(db)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		album := v.Album
		marshal, err := serialize.UnMarshal([]byte(album))
		if err != nil {
			panic(err)
		}
		if marshal != nil {
			newPics := make([]string, 0)
			pics := marshal.([]interface{})
			for _, pic := range pics {
				if pic != nil {
					if strings.Contains(pic.(string), "https://w7.dapp100.cn/attachment/") {
						download(pic.(string))
					}
				}
				newPics = append(newPics, strings.ReplaceAll(pic.(string), "https://w7.dapp100.cn/attachment/", ""))
			}
			piscByte, _ := serialize.MarshalSlice(newPics)
			ups := map[string]interface{}{
				"album": string(piscByte),
				"thumb": string(piscByte),
			}
			if err := v.Updates(db, ups); err != nil {
				panic(err)
			}

		}

	}
}
