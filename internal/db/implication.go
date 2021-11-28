package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ImplicationAntecedent(ir ImplicationRequest) (ImplicationResponse, error) {
	err := validateImplicationRequest(ir)
	if err != nil {
		log.Print("Malformed request!")
		return ImplicationResponse{}, err
	}

	inc := make(map[Ling]bool)
	lMap := make(map[Ling][]NameValuePair)
	pMap := make(map[NameValuePair]int)
	is := make([]NameValuePair, 0)

	// pass group then lings into query args
	qargs := make([]interface{}, 1)
	qargs[0] = ir.Group

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ?`

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return ImplicationResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var l Ling
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return ImplicationResponse{}, err
		}

		if nv.Name == ir.Property.Name && nv.Value == ir.Property.Value {
			inc[l] = true
		} else {
			if ps, ok := lMap[l]; ok {
				lMap[l] = append(ps, nv)
			} else {
				lMap[l] = []NameValuePair{nv}
			}
		}
	}

	for l, _ := range inc {
		if ps, ok := lMap[l]; ok {
			for _, p := range ps {
				pMap[p]++
			}
		}
	}

	for p, c := range pMap {
		if c == len(inc) {
			is = append(is, p)
		}
	}

	return ImplicationResponse{
		Type:         "implication",
		Property:     ir.Property,
		Direction:    "antecedent",
		Implications: is,
	}, nil
}

func ImplicationConsequent(ir ImplicationRequest) (ImplicationResponse, error) {
	err := validateImplicationRequest(ir)
	if err != nil {
		log.Print("Malformed request!")
		return ImplicationResponse{}, err
	}

	inc := make(map[Ling]bool)
	lMap := make(map[Ling][]NameValuePair)
	pMap := make(map[NameValuePair]int)
	is := make([]NameValuePair, 0)

	// pass group then lings into query args
	qargs := make([]interface{}, 1)
	qargs[0] = ir.Group

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ?`

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return ImplicationResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var l Ling
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return ImplicationResponse{}, err
		}

		if nv.Name == ir.Property.Name && nv.Value == ir.Property.Value {
			inc[l] = true
		} else {
			if ps, ok := lMap[l]; ok {
				lMap[l] = append(ps, nv)
			} else {
				lMap[l] = []NameValuePair{nv}
			}
		}
	}

	for l, _ := range inc {
		if ps, ok := lMap[l]; ok {
			for _, p := range ps {
				pMap[p]++
			}
		}
	}

	for l, ps := range lMap {
		if _, ok := inc[l]; !ok {
			for _, p := range ps {
				if _, ok := pMap[p]; ok {
					delete(pMap, p)
				}
			}
		}
	}

	for p, _ := range pMap {
		is = append(is, p)
	}

	return ImplicationResponse{
		Type:         "implication",
		Property:     ir.Property,
		Direction:    "consequent",
		Implications: is,
	}, nil
}

func ImplicationBoth(ir ImplicationRequest) (ImplicationResponse, error) {
	err := validateImplicationRequest(ir)
	if err != nil {
		log.Print("Malformed request!")
		return ImplicationResponse{}, err
	}

	inc := make(map[Ling]bool)
	lMap := make(map[Ling][]NameValuePair)
	pMap := make(map[NameValuePair]int)
	is := make([]NameValuePair, 0)

	// pass group then lings into query args
	qargs := make([]interface{}, 1)
	qargs[0] = ir.Group

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ?`

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return ImplicationResponse{}, err
	}
	defer ls.Close()

	for ls.Next() {
		var l Ling
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return ImplicationResponse{}, err
		}

		if nv.Name == ir.Property.Name && nv.Value == ir.Property.Value {
			inc[l] = true
		} else {
			if ps, ok := lMap[l]; ok {
				lMap[l] = append(ps, nv)
			} else {
				lMap[l] = []NameValuePair{nv}
			}
		}
	}

	for l, _ := range inc {
		if ps, ok := lMap[l]; ok {
			for _, p := range ps {
				pMap[p]++
			}
		}
	}

	for l, ps := range lMap {
		if _, ok := inc[l]; !ok {
			for _, p := range ps {
				if _, ok := pMap[p]; ok {
					delete(pMap, p)
				}
			}
		}
	}

	for p, c := range pMap {
		if c == len(inc) {
			is = append(is, p)
		}
	}

	return ImplicationResponse{
		Type:         "implication",
		Property:     ir.Property,
		Direction:    "both",
		Implications: is,
	}, nil
}
