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
	lings := make([]string, len(clr.Lings))

	// pass group then lings into query args
	qargs := make([]interface{}, len(clr.Lings)+1)
	qargs[0] = clr.Group
	for i, id := range clr.Lings {
		qargs[i+1] = id
	}

	// SELECT lings
	stmt := `SELECT id, name FROM lings WHERE group_id = ? AND depth = 0 AND id IN (?` + strings.Repeat(",?", len(clr.Lings)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return CompareLingsResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l Ling

		err = ls.Scan(&l.Id, &l.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return CompareLingsResponse{}, err
		}

		lings[i] = l.Name

		// build statement dynamically
		stmt = "SELECT properties.id, properties.name, lings_properties.value FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ?"

		// SELECT properties
		ps, err := db.Query(stmt, clr.Group, l.Id)
		if err != nil {
			log.Print("Error preparing database request!")
			return CompareLingsResponse{}, err
		}
		defer ps.Close()

		for ps.Next() {
			var nv NameValuePair

			err = ps.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return CompareLingsResponse{}, err
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

		i++
	}

	common := make([]NameValuePair, len(vMap))
	i = 0

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
		On:       lings,
		Common:   common,
		Distinct: distinct,
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
	lings := make([]string, len(cllr.Linglets))

	// pass group then lings into query args
	qargs := make([]interface{}, len(cllr.Linglets)+1)
	qargs[0] = cllr.Group
	for i, id := range cllr.Linglets {
		qargs[i+1] = id
	}

	// SELECT lings
	stmt := `SELECT id, name FROM lings WHERE group_id = ? AND depth = 1 AND id IN (?` + strings.Repeat(",?", len(cllr.Linglets)-1) + `)`
	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return CompareLingletsResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l Ling

		err = ls.Scan(&l.Id, &l.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return CompareLingletsResponse{}, err
		}

		lings[i] = l.Name

		// build statement dynamically
		stmt = "SELECT properties.id, properties.name, lings_properties.value FROM lings_properties INNER JOIN properties ON lings_properties.property_id = properties.id WHERE lings_properties.group_id = ? AND ling_id = ?"

		// SELECT properties
		ps, err := db.Query(stmt, cllr.Group, l.Id)
		if err != nil {
			log.Print("Error preparing database request!")
			return CompareLingletsResponse{}, err
		}
		defer ps.Close()

		for ps.Next() {
			var nv NameValuePair

			err = ps.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return CompareLingletsResponse{}, err
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

		i++
	}

	common := make([]NameValuePair, len(vMap))
	i = 0

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
		On:       lings,
		Common:   common,
		Distinct: distinct,
	}, nil
}
