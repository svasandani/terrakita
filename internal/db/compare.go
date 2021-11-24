package db

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func CompareLings(clr CompareLingsRequest) (CompareLingsResponse, error) {
	err := validateCompareLingsRequest(clr)
	if err != nil {
		log.Print("Malformed request!")
		return CompareLingsResponse{}, err
	}

	vMap := make(map[string]string)          // property values
	nMap := make(map[string]string)          // property names
	cMap := make(map[string]int)             // common property counts
	dMap := make(map[string][]NameValuePair) // data values
	lMap := make(map[string]bool)
	lings := make([]string, 0)

	// pass group then lings into query args
	qargs := make([]interface{}, len(clr.Lings)+1)
	qargs[0] = clr.Group
	for i, id := range clr.Lings {
		qargs[i+1] = id
	}

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ? AND lings.depth = 0 AND lings.id IN (?` + strings.Repeat(",?", len(clr.Lings)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return CompareLingsResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var l Ling
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return CompareLingsResponse{}, err
		}

		if _, ok := lMap[l.Id]; !ok {
			lings = append(lings, l.Name)
			lMap[l.Id] = true
		}

		if d, ok := dMap[nv.Id]; ok {
			// we've seen this property before
			dMap[nv.Id] = append(d, NameValuePair{
				Id:    l.Id,
				Name:  l.Name,
				Value: nv.Value,
			})

			if v, ok := vMap[nv.Id]; ok {
				// property value is common so far
				if v == nv.Value {
					// property value is still common
					cMap[nv.Id]++
				} else {
					// property is no longer common
					delete(vMap, nv.Id)
					delete(cMap, nv.Id)
				}
			}
		} else {
			// first time seeing this property
			dMap[nv.Id] = []NameValuePair{{
				Id:    l.Id,
				Name:  l.Name,
				Value: nv.Value,
			}}
			nMap[nv.Id] = nv.Name
			vMap[nv.Id] = nv.Value
			cMap[nv.Id] = 1
		}
	}

	common := make([]NameValuePair, len(vMap))
	i := 0

	for id, v := range vMap {
		if cMap[id] != len(clr.Lings) {
			continue
		}

		common[i] = NameValuePair{
			Id:    id,
			Name:  nMap[id],
			Value: v,
		}

		i++
	}

	distinct := make([]CompareLingsResponseProperty, len(dMap)-i)
	j := 0

	for id, d := range dMap {
		if _, ok := vMap[id]; ok && cMap[id] == len(clr.Lings) {
			continue
		}

		distinct[j] = CompareLingsResponseProperty{
			Id:             id,
			Name:           nMap[id],
			LingValuePairs: d,
		}

		j++
	}

	return CompareLingsResponse{
		Type:     "compare",
		On:       "lings",
		Lings:    lings,
		Common:   common[:i],
		Distinct: distinct[:j],
	}, nil
}

func CompareLinglets(cllr CompareLingletsRequest) (CompareLingletsResponse, error) {
	err := validateCompareLingletsRequest(cllr)
	if err != nil {
		log.Print("Malformed request!")
		return CompareLingletsResponse{}, err
	}

	vMap := make(map[string]string)          // property values
	nMap := make(map[string]string)          // property names
	cMap := make(map[string]int)             // common property counts
	dMap := make(map[string][]NameValuePair) // data values
	llMap := make(map[string]bool)
	linglets := make([]string, 0)

	// pass group then lings into query args
	qargs := make([]interface{}, len(cllr.Linglets)+1)
	qargs[0] = cllr.Group
	for i, id := range cllr.Linglets {
		qargs[i+1] = id
	}

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ? AND lings.depth = 1 AND lings.id IN (?` + strings.Repeat(",?", len(cllr.Linglets)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return CompareLingletsResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var ll Linglet
		var nv NameValuePair

		err = ls.Scan(&ll.Id, &ll.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return CompareLingletsResponse{}, err
		}

		if _, ok := llMap[ll.Id]; !ok {
			linglets = append(linglets, ll.Name)
			llMap[ll.Id] = true
		}

		if d, ok := dMap[nv.Id]; ok {
			// we've seen this property before
			dMap[nv.Id] = append(d, NameValuePair{
				Id:    ll.Id,
				Name:  ll.Name,
				Value: nv.Value,
			})

			if v, ok := vMap[nv.Id]; ok {
				// property value is common so far
				if v == nv.Value {
					// property value is still common
					cMap[nv.Id]++
				} else {
					// property is no longer common
					delete(vMap, nv.Id)
					delete(cMap, nv.Id)
				}
			}
		} else {
			// first time seeing this property
			dMap[nv.Id] = []NameValuePair{{
				Id:    ll.Id,
				Name:  ll.Name,
				Value: nv.Value,
			}}
			nMap[nv.Id] = nv.Name
			vMap[nv.Id] = nv.Value
			cMap[nv.Id] = 1
		}
	}

	common := make([]NameValuePair, len(vMap))
	i := 0

	for id, v := range vMap {
		if cMap[id] != len(cllr.Linglets) {
			continue
		}

		common[i] = NameValuePair{
			Id:    id,
			Name:  nMap[id],
			Value: v,
		}

		i++
	}

	distinct := make([]CompareLingletsResponseProperty, len(dMap)-i)
	j := 0

	for id, d := range dMap {
		if _, ok := vMap[id]; ok && cMap[id] == len(cllr.Linglets) {
			continue
		}

		distinct[j] = CompareLingletsResponseProperty{
			Id:                id,
			Name:              nMap[id],
			LingletValuePairs: d,
		}

		j++
	}

	return CompareLingletsResponse{
		Type:     "compare",
		On:       "linglets",
		Linglets: linglets,
		Common:   common[:i],
		Distinct: distinct[:j],
	}, nil
}
