package db

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func FilterLings(flr FilterLingsRequest) (FilterLingsResponse, error) {
	err := validateFilterLingsRequest(flr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingsResponse{}, err
	}

	lings := make([]FilterLingsResponseLing, len(flr.Lings))
	lMap := make(map[Ling][]NameValuePair)

	// pass group then lings into query args
	qargs := make([]interface{}, len(flr.Lings)+len(flr.LingProperties)+1)
	qargs[0] = flr.Group
	for i, id := range flr.Lings {
		qargs[i+1] = id
	}
	for i, id := range flr.LingProperties {
		qargs[i+len(flr.Lings)+1] = id
	}

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ? AND lings.depth = 0 AND lings.id IN (?` + strings.Repeat(",?", len(flr.Lings)-1) + `)`
	if len(flr.LingProperties) != 0 {
		stmt += ` AND properties.id IN (?` + strings.Repeat(",?", len(flr.LingProperties)-1) + `)`
	}

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingsResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var l Ling
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingsResponse{}, err
		}

		if ps, ok := lMap[l]; ok {
			lMap[l] = append(ps, nv)
		} else {
			lMap[l] = []NameValuePair{nv}
		}
	}

	i := 0

	for l, ns := range lMap {
		if flr.LingPropertiesInclusive && len(ns) < len(flr.LingProperties) {
			continue
		}

		lings[i] = FilterLingsResponseLing{
			Id:                 l.Id,
			Name:               l.Name,
			PropertyValuePairs: ns,
		}
		i++
	}

	return FilterLingsResponse{
		Type:  "filter",
		On:    "lings",
		Lings: lings[:i],
	}, nil
}

func FilterLingProperties(flpr FilterLingPropertiesRequest) (FilterLingPropertiesResponse, error) {
	err := validateFilterLingPropertiesRequest(flpr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingPropertiesResponse{}, err
	}

	properties := make([]FilterLingPropertiesResponseProperty, len(flpr.LingProperties))

	// pass group then ling properties into query args
	qargs := make([]interface{}, len(flpr.LingProperties)+1)
	qargs[0] = flpr.Group
	for i, id := range flpr.LingProperties {
		qargs[i+1] = id
	}

	// SELECT properties
	stmt := `SELECT id, name FROM properties WHERE group_id = ? AND id IN (?` + strings.Repeat(",?", len(flpr.LingProperties)-1) + `)`
	ps, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingPropertiesResponse{}, err
	}
	defer ps.Close()

	i := 0

	for ps.Next() {
		var p FilterLingPropertiesResponseProperty

		err = ps.Scan(&p.Id, &p.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingPropertiesResponse{}, err
		}

		qargs = make([]interface{}, len(flpr.Lings)+2)
		qargs[0] = flpr.Group
		qargs[1] = p.Id
		for i, id := range flpr.Lings {
			qargs[i+2] = id
		}

		// if we need all lings, perform an extra query in case of early return
		if len(flpr.Lings) > 0 && flpr.LingsInclusive {
			stmt = `SELECT COUNT(lings.id) FROM lings_properties INNER JOIN lings ON lings_properties.ling_id = lings.id WHERE lings.depth = 0 AND lings_properties.group_id = ? AND ling_id = ? AND lings_properties.ling_id IN (?` + strings.Repeat(",?", len(flpr.Lings)-1) + `)`
			c := db.QueryRow(stmt, qargs...)

			var count int
			c.Scan(&count)

			if count != len(flpr.Lings) {
				continue
			}
		}

		// build statement dynamically
		stmt = "SELECT lings.id, lings.name, lings_properties.value FROM lings_properties INNER JOIN lings ON lings_properties.ling_id = lings.id WHERE lings.depth = 0 AND lings_properties.group_id = ? AND property_id = ?"
		if len(flpr.Lings) != 0 {
			stmt += ` AND lings_properties.ling_id IN (?` + strings.Repeat(",?", len(flpr.Lings)-1) + `)`
		}

		// SELECT lings
		ls, err := db.Query(stmt, qargs...)
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterLingPropertiesResponse{}, err
		}
		defer ls.Close()

		lvs := make([]NameValuePair, 0)

		for ls.Next() {
			var nv NameValuePair

			err = ls.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return FilterLingPropertiesResponse{}, err
			}

			lvs = append(lvs, nv)
		}

		if len(lvs) == 0 {
			continue
		}

		p.LingValuePairs = lvs
		properties[i] = p
		i++
	}

	return FilterLingPropertiesResponse{
		Type:       "filter",
		On:         "ling_properties",
		Properties: properties[:i],
	}, nil
}

func FilterLinglets(fllr FilterLingletsRequest) (FilterLingletsResponse, error) {
	err := validateFilterLingletsRequest(fllr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingletsResponse{}, err
	}

	lings := make([]FilterLingletsResponseLing, 0)
	lMap := make(map[Ling]map[Linglet][]NameValuePair)

	// pass group then lings into query args
	qargs := make([]interface{}, len(fllr.Linglets)+len(fllr.LingletProperties)+1)
	qargs[0] = fllr.Group
	for i, id := range fllr.Linglets {
		qargs[i+1] = id
	}
	for i, id := range fllr.LingletProperties {
		qargs[i+len(fllr.Linglets)+1] = id
	}

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, linglets.id, linglets.name, properties.id, properties.name, lings_properties.value FROM lings AS linglets INNER JOIN lings_properties ON linglets.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id INNER JOIN lings ON linglets.parent_id=lings.id WHERE lings.group_id = ? AND linglets.depth = 1 AND linglets.id IN (?` + strings.Repeat(",?", len(fllr.Linglets)-1) + `)`
	if len(fllr.LingletProperties) != 0 {
		stmt += ` AND properties.id IN (?` + strings.Repeat(",?", len(fllr.LingletProperties)-1) + `)`
	}

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingletsResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var l Ling
		var ll Linglet
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &ll.Id, &ll.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingletsResponse{}, err
		}

		if llMap, ok := lMap[l]; ok {
			if ps, ok := llMap[ll]; ok {
				llMap[ll] = append(ps, nv)
			} else {
				llMap[ll] = []NameValuePair{nv}
			}

			lMap[l] = llMap
		} else {
			llMap = make(map[Linglet][]NameValuePair)
			llMap[ll] = []NameValuePair{nv}

			lMap[l] = llMap
		}

	}

	for l, llMap := range lMap {
		if len(llMap) == 0 {
			continue
		}

		linglets := make([]FilterLingletsResponseLinglet, len(llMap))

		i := 0

		for ll, ns := range llMap {
			if fllr.LingletPropertiesInclusive && len(ns) < len(fllr.LingletProperties) {
				continue
			}

			linglets[i] = FilterLingletsResponseLinglet{
				Id:                 ll.Id,
				Name:               ll.Name,
				PropertyValuePairs: ns,
			}

			i++
		}

		lings = append(lings, FilterLingletsResponseLing{
			Id:       l.Id,
			Name:     l.Name,
			Linglets: linglets[:i],
		})
	}

	return FilterLingletsResponse{
		Type:  "filter",
		On:    "linglets",
		Lings: lings,
	}, nil
}

func FilterLingletProperties(fllpr FilterLingletPropertiesRequest) (FilterLingletPropertiesResponse, error) {
	err := validateFilterLingletPropertiesRequest(fllpr)
	if err != nil {
		log.Print("Malformed request!")
		return FilterLingletPropertiesResponse{}, err
	}

	properties := make([]FilterLingletPropertiesResponseProperty, len(fllpr.LingletProperties))

	// pass group then ling properties into query args
	qargs := make([]interface{}, len(fllpr.LingletProperties)+1)
	qargs[0] = fllpr.Group
	for i, id := range fllpr.LingletProperties {
		qargs[i+1] = id
	}

	// SELECT properties
	stmt := `SELECT id, name FROM properties WHERE group_id = ? AND id IN (?` + strings.Repeat(",?", len(fllpr.LingletProperties)-1) + `)`
	ps, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return FilterLingletPropertiesResponse{}, err
	}
	defer ps.Close()

	i := 0

	for ps.Next() {
		var p FilterLingletPropertiesResponseProperty

		err = ps.Scan(&p.Id, &p.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return FilterLingletPropertiesResponse{}, err
		}

		qargs = make([]interface{}, len(fllpr.Linglets)+2)
		qargs[0] = fllpr.Group
		qargs[1] = p.Id
		for i, id := range fllpr.Linglets {
			qargs[i+2] = id
		}

		// if we need all linglets, perform an extra query in case of early return
		if len(fllpr.Linglets) > 0 && fllpr.LingletsInclusive {
			stmt = `SELECT COUNT(lings.id) FROM lings_properties INNER JOIN lings ON lings_properties.ling_id = lings.id WHERE lings.depth = 1 AND lings_properties.group_id = ? AND ling_id = ? AND lings_properties.ling_id IN (?` + strings.Repeat(",?", len(fllpr.Linglets)-1) + `)`
			c := db.QueryRow(stmt, qargs...)

			var count int
			c.Scan(&count)

			if count != len(fllpr.Linglets) {
				continue
			}
		}

		// build statement dynamically
		stmt = "SELECT lings.id, lings.name, lings_properties.value FROM lings_properties INNER JOIN lings ON lings_properties.ling_id = lings.id WHERE lings.depth = 1 AND lings_properties.group_id = ? AND property_id = ?"
		if len(fllpr.Linglets) != 0 {
			stmt += ` AND lings_properties.ling_id IN (?` + strings.Repeat(",?", len(fllpr.Linglets)-1) + `)`
		}

		// SELECT linglets
		lls, err := db.Query(stmt, qargs...)
		if err != nil {
			log.Print("Error preparing database request!")
			return FilterLingletPropertiesResponse{}, err
		}
		defer lls.Close()

		llvs := make([]NameValuePair, 0)

		for lls.Next() {
			var nv NameValuePair

			err = lls.Scan(&nv.Id, &nv.Name, &nv.Value)
			if err != nil {
				log.Print("Error executing database request!")
				return FilterLingletPropertiesResponse{}, err
			}

			llvs = append(llvs, nv)
		}

		if len(llvs) == 0 {
			continue
		}

		p.LingletValuePairs = llvs
		properties[i] = p
		i++
	}

	return FilterLingletPropertiesResponse{
		Type:       "filter",
		On:         "linglet_properties",
		Properties: properties[:i],
	}, nil
}
