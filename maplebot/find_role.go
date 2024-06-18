package maplebot

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"fmt"
	. "github.com/CuteReimu/onebot"
	. "github.com/vicanso/go-charts/v2"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed lucid_bkg_mid.jpeg
var bkgFile []byte

var bkg *image.Image

func init() {
	file, err := jpeg.Decode(bytes.NewReader(bkgFile))
	if err != nil {
		slog.Error("解析背景图片失败", "error", err)
	} else {
		bkg = &file
	}
}

type findRoleReturnData struct {
	CharacterData struct {
		CharacterImageURL string  `json:"CharacterImageURL"`
		Class             string  `json:"Class"`
		EXPPercent        float64 `json:"EXPPercent"`
		GraphData         []struct {
			CurrentEXP int64  `json:"CurrentEXP"`
			DateLabel  string `json:"DateLabel"`
			Level      int    `json:"Level"`
		} `json:"GraphData"`
		LegionLevel int    `json:"LegionLevel"`
		Level       int    `json:"Level"`
		Name        string `json:"Name"`
	} `json:"CharacterData"`
}

func findRole(name string) MessageChain {
	resp, err := restyClient.R().Get("https://api.maplestory.gg/v2/public/character/gms/" + name)
	if err != nil {
		slog.Error("请求失败", "error", err)
		return nil
	}
	switch resp.StatusCode() {
	case 404:
		return MessageChain{&Text{Text: name + "已身死道消"}}
	case 200:
	default:
		slog.Error("请求失败", "status", resp.StatusCode())
		return nil
	}
	body := resp.Body()
	var data *findRoleReturnData
	if err = json.Unmarshal(body, &data); err != nil {
		slog.Error("json解析失败", "error", err, "body", body)
		return nil
	}
	imgUrl := data.CharacterData.CharacterImageURL
	rawName := data.CharacterData.Name
	class := TranslateClassName(data.CharacterData.Class)
	level := data.CharacterData.Level
	exp := data.CharacterData.EXPPercent
	levelExp := fmt.Sprintf("%d (%g%%)", level, exp)
	legionLevel := data.CharacterData.LegionLevel

	var messageChain MessageChain
	if len(imgUrl) > 0 {
		resp, err := restyClient.R().Get(imgUrl)
		if err != nil {
			slog.Error("请求失败", "error", err)
		} else {
			messageChain = append(messageChain, &Image{File: "base64://" + base64.StdEncoding.EncodeToString(resp.Body())})
		}
	}

	s := fmt.Sprintf("角色名：%s\n职业：%s\n等级：%s\n联盟：%d\n", rawName, class, levelExp, legionLevel)

	expValues := make([]float64, 0, 14)
	levelValues := make([]float64, 0, 15)
	labels := make([]string, 0, 15)
	a := data.CharacterData.GraphData
	if len(a) > 0 {
		for _, d := range a {
			labels = append(labels, d.DateLabel)
			levelUpExp := levelExpData.GetFloat64(fmt.Sprintf("data.%d", d.Level))
			var expPercent float64
			if levelUpExp > 0 {
				expPercent = float64(d.CurrentEXP) / levelUpExp
			}
			levelValues = append(levelValues, float64(d.Level)+expPercent)
		}
		if len(labels) < 15 {
			t, err := time.Parse("2006-01-02", a[0].DateLabel)
			for len(labels) < 15 {
				if err != nil {
					labels = append([]string{""}, labels...)
				} else {
					t = t.Add(-24 * time.Hour)
					labels = append([]string{t.Format("2006-01-02")}, labels...)
				}
				// maplestory.gg只统计220级以上角色的经验数据，没数据我们就只能认为是220级了
				levelValues = append([]float64{220}, levelValues...)
			}
		}
		// 处理一下，有可能有的数据是0级
		for i := 1; i < len(levelValues); i++ {
			if levelValues[i] == 0 {
				levelValues[i] = levelValues[i-1]
			}
		}
		for i := len(levelValues) - 2; i >= 0; i-- {
			if levelValues[i] == 0 {
				levelValues[i] = levelValues[i+1]
			}
		}
		for i := 1; i < len(levelValues); i++ {
			level0, level1 := int(levelValues[i-1]), int(levelValues[i])
			expPercent0, expPercent1 := levelValues[i-1]-float64(level0), levelValues[i]-float64(level1)
			var totalExp float64
			for j := level0; j < level1; j++ {
				totalExp += levelExpData.GetFloat64(fmt.Sprintf("data.%d", j))
			}
			totalExp -= expPercent0 * levelExpData.GetFloat64(fmt.Sprintf("data.%d", level0))
			totalExp += expPercent1 * levelExpData.GetFloat64(fmt.Sprintf("data.%d", level1))
			expValues = append(expValues, max(math.Round(totalExp), 0))
		}
		labels = labels[1:]
		for i := range labels {
			labels[i] = labels[i][5:]
		}
		levelValues = levelValues[1:]
		maxValue := slices.Max(expValues)
		digits := int(math.Floor(math.Log10(maxValue))) + 1
		factor := math.Pow(10, float64(digits-1))
		highest := math.Floor(maxValue / factor)
		if highest < 2 {
			factor /= 5
		} else if highest < 5 {
			factor /= 2
		}
		maxValue = math.Ceil(maxValue/factor) * factor
		divideCount := int(maxValue / factor)
		var yAxisOptions []YAxisOption
		yAxisOptions = append(yAxisOptions, YAxisOption{Min: NewFloatPoint(0), Max: &maxValue, DivideCount: divideCount, Unit: 1})
		for _, diff := range []float64{0.1, 0.2, 0.5, 1, 2, 5, 10, 20, 50, 100} {
			minLevel, maxLevel := math.Floor(slices.Min(levelValues)/diff)*diff, math.Ceil(slices.Max(levelValues)/diff)*diff
			remainCount := divideCount - int(math.Round((maxLevel-minLevel)/diff))
			if remainCount >= 0 {
				if remainCount > 0 {
					maxLevel += diff * float64(remainCount-remainCount/2)
					minLevel -= diff * float64(remainCount/2)
				}
				if minLevel <= 220 { // 为了图表好看，最低显示220级
					maxLevel -= minLevel - 220
					minLevel = 220
				}
				yAxisOptions = append(yAxisOptions, YAxisOption{Min: NewFloatPoint(minLevel), Max: NewFloatPoint(maxLevel), DivideCount: divideCount, Unit: 1})
				break
			}
		}
		if slices.ContainsFunc(expValues, func(f float64) bool { return f != 0 }) {
			if levelEnd := strings.Index(levelExp, "("); levelEnd >= 0 {
				if totalExp := float64(levelExpData.GetInt64("data." + strings.TrimSpace(levelExp[:levelEnd]))); totalExp > 0 {
					if expPercentStart := strings.Index(levelExp, "("); expPercentStart >= 0 {
						if expPercentEnd := strings.Index(levelExp, "%"); expPercentEnd >= 0 {
							if expPercent, err := strconv.ParseFloat(levelExp[expPercentStart+1:expPercentEnd], 64); err == nil {
								var sumExp, n float64
								for _, v := range expValues {
									if v != 0 || n > 0 {
										sumExp += v
										n++
									}
								}
								aveExp := sumExp / n
								days := int(math.Ceil((totalExp - totalExp/100.0*expPercent) / aveExp))
								s += fmt.Sprintf("预计还有%d天升级\n", days)
							}
						}
					}
				}
			}
			var seriesList []Series
			seriesList = append(seriesList, NewSeriesFromValues(expValues, ChartTypeBar))
			if len(levelValues) == len(expValues) {
				seriesList = append(seriesList, NewSeriesFromValues(levelValues, ChartTypeLine))
				seriesList[1].AxisIndex = 1
			}
			p, err := Render(
				ChartOption{SeriesList: seriesList},
				ThemeOptionFunc(ThemeDark),
				PaddingOptionFunc(Box{Top: 30, Left: 10, Right: 10, Bottom: 10}),
				XAxisDataOptionFunc(labels),
				//MarkLineOptionFunc(0, SeriesMarkDataTypeAverage),
				YAxisOptionFunc(yAxisOptions...),
				func(opt *ChartOption) {
					opt.XAxis.TextRotation = -math.Pi / 4
					opt.XAxis.LabelOffset = Box{Top: -5, Left: 7}
					opt.XAxis.FontSize = 7.5
					opt.ValueFormatter = func(f float64) string {
						switch {
						case f == 220.0:
							return "\u3000" // 不显示220级的坐标，用全角空格代替
						case f < 1000.0:
							return fmt.Sprintf("%g", math.Round(f*1000)/1000)
						case f < 1000000.0:
							return fmt.Sprintf("%gK", math.Round(f)/1000)
						case f < 1000000000.0:
							return fmt.Sprintf("%gM", math.Round(f)/1000000)
						case f < 1000000000000.0:
							return fmt.Sprintf("%gB", math.Round(f)/1000000000)
						default:
							return fmt.Sprintf("%gT", math.Round(f)/1000000000000)
						}
					}
				},
			)
			if err != nil {
				slog.Error("render chart failed", "error", err)
			} else if buf, err := p.Bytes(); err != nil {
				slog.Error("render chart failed", "error", err)
			} else {
				if bkg != nil {
					img, err := png.Decode(bytes.NewReader(buf))
					if err != nil {
						slog.Error("解析图片失败", "error", err)
					} else {
						mask := &image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 56}}
						newImg := image.NewRGBA(img.Bounds())
						draw.Draw(newImg, newImg.Bounds(), img, image.Point{}, draw.Src)
						draw.DrawMask(newImg, newImg.Bounds(), *bkg, image.Point{}, mask, image.Point{}, draw.Over)
						buf2 := &bytes.Buffer{}
						if err = png.Encode(buf2, newImg); err != nil {
							slog.Error("生成图片失败", "error", err)
						} else {
							buf = buf2.Bytes()
						}
					}
				}
				messageChain = append(messageChain, &Text{Text: s}, &Image{File: "base64://" + base64.StdEncoding.EncodeToString(buf)})
			}
		} else {
			messageChain = append(messageChain, &Text{Text: s + "近日无经验变化"})
		}
	} else {
		messageChain = append(messageChain, &Text{Text: s + "近日无经验变化"})
	}
	return messageChain
}
