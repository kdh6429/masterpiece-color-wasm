package myprocessor

import (
	"encoding/base64"
	"image"
	"log"
	"strings"

	"image/jpeg"
	"image/png"

	"github.com/EdlinOrg/prominentcolor"
)

func createMap(colors string, state bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"colors":  colors,
		"state":   state,
		"message": message,
	}
}

func Do(base64string string) interface{} {
	var img image.Image = nil

	i := strings.Index(base64string, ";base64")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64string[i+8:]))
	switch strings.TrimSuffix(base64string[5:i], ";base64") {
	case "image/png":
		image, err := png.Decode(reader)
		if err != nil {
			return createMap("", false, "Cannot Decode png image "+err.Error())
		}
		img = image
	case "image/jpeg":
		image, err := jpeg.Decode(reader)
		if err != nil {
			return createMap("", false, "Cannot Decode jpeg image "+err.Error())
		}
		img = image
	}
	if img == nil {
		return createMap("", false, "img object cannot be nil")
	}
	//cols, err := prominentcolor.KmeansWithArgs(prominentcolor.ArgumentNoCropping, img)

	//KmeansWithAll(k, img, bitarr[i], resizeSize, bgmasks)

	rtnClors := ""
	resizeSize := uint(prominentcolor.DefaultSize)
	bgmasks := prominentcolor.GetDefaultMasks()
	for i := 2; i <= 4; i++ {
		cols, err := prominentcolor.KmeansWithAll(
			i,
			img,
			prominentcolor.ArgumentNoCropping,
			resizeSize,
			bgmasks,
		)
		if err != nil {
			log.Println(err)
			continue
			//return Result{nil, false, ""}
		}
		tmpColor := ""
		for _, col := range cols {
			if tmpColor == "" {
				tmpColor += ("#" + col.AsString())
			} else {
				tmpColor += "," + ("#" + col.AsString())
			}
		}
		rtnClors += "|" + tmpColor
	}
	return createMap(rtnClors, true, "")
}
