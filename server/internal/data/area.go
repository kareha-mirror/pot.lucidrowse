package data

import (
	"context"
	"fmt"
	"strings"
)

var regionNames = map[string]string{
	"capital": "王都地方",
	"west":    "西方地方",
	"north":   "北方地方",
	"east":    "東方地方",
	"south":   "南方地方",
	"coast":   "海岸地方",
	"islands": "島嶼地方",
}

type AreaItem struct {
	RegionCode string
	AreaCode   string
	Name       string
}

var areaSeeds = []AreaItem{
	{"capital", "city", "王都"},
	{"capital", "outskirts", "近郊"},
	{"west", "road", "街道"},
	{"west", "hills", "丘陵"},
	{"north", "mountains", "山岳"},
	{"north", "highlands", "高原"},
	{"east", "forest", "森林"},
	{"east", "wetlands", "湖沼"},
	{"south", "plains", "平野"},
	{"south", "farmland", "農村"},
	{"coast", "harbor", "港湾"},
	{"coast", "shore", "海岸"},
	{"islands", "main", "大島"},
	{"islands", "outer", "周辺諸島"},
}

func SeedAreas() error {
	ctx := context.Background()

	for _, area := range areaSeeds {
		_, err := db.Exec(ctx, `
			INSERT INTO areas (region_code, area_code, name)
			VALUES ($1, $2, $3)
			ON CONFLICT (region_code, area_code) DO UPDATE
			SET name = EXCLUDED.name
			`, area.RegionCode, area.AreaCode, area.Name,
		)
		if err != nil {
			return fmt.Errorf(
				"seed area %q-%q: %w", area.RegionCode, area.AreaCode, err,
			)
		}
	}

	return nil
}

func AreaList() ([]AreaItem, error) {
	rows, err := db.Query(context.Background(), `
		SELECT region_code, area_code, name
		FROM areas
		ORDER BY id
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areas []AreaItem

	for rows.Next() {
		var a AreaItem

		err := rows.Scan(
			&a.RegionCode,
			&a.AreaCode,
			&a.Name,
		)
		if err != nil {
			return nil, err
		}

		areas = append(areas, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return areas, nil
}

func DescribeAreas() (string, error) {
	areas, err := AreaList()
	if err != nil {
		return "", err
	}

	b := strings.Builder{}

	for _, area := range areas {
		b.WriteString(fmt.Sprintf("code: %s-%s  name: %s %s\n", area.RegionCode, area.AreaCode, regionNames[area.RegionCode], area.Name))
	}

	return b.String(), nil
}
