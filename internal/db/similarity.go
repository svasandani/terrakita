package db

import (
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func SimilarityLings(slr SimilarityLingsRequest) (SimilarityLingsResponse, error) {
	err := validateSimilarityLingsRequest(slr)
	if err != nil {
		log.Print("Malformed request!")
		return SimilarityLingsResponse{}, err
	}

	pMap := make(map[NameValuePair][]Ling)
	prMap := make(map[[2]string]int)
	iMap := make(map[string]Ling)
	lings := make([]string, len(slr.Lings))

	// pass group then lings into query args
	qargs := make([]interface{}, len(slr.Lings)+1)
	qargs[0] = slr.Group
	for i, id := range slr.Lings {
		qargs[i+1] = id
	}

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ? AND lings.depth = 0 AND lings.id IN (?` + strings.Repeat(",?", len(slr.Lings)-1) + `)`

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return SimilarityLingsResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l Ling
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return SimilarityLingsResponse{}, err
		}

		if ls, ok := pMap[nv]; ok {
			pMap[nv] = append(ls, l)
		} else {
			pMap[nv] = []Ling{l}
		}

		if _, ok := iMap[l.Id]; !ok {
			lings[i] = l.Name
			iMap[l.Id] = l

			i++
		}
	}

	// oof
	for _, ls := range pMap {
		for i, l := range ls {
			for j := i + 1; j < len(ls); j++ {
				k := ls[j]

				// deterministic order
				var first Ling
				var second Ling
				if l.Id < k.Id {
					first = l
					second = k
				} else {
					first = k
					second = l
				}

				key := [2]string{first.Name, second.Name}
				prMap[key]++
			}
		}
	}

	pairs := make([]SimilarityLingsResponsePair, len(prMap))
	j := 0

	for pr, c := range prMap {
		ps := make([]string, 2)
		copy(ps, pr[:])

		pairs[j] = SimilarityLingsResponsePair{
			Lings:                ps,
			CommonPropertyValues: c,
		}

		j++
	}

	return SimilarityLingsResponse{
		Type:  "similarity",
		On:    "lings",
		Lings: lings[:i],
		Pairs: pairs,
	}, nil
}

func SimilarityLinglets(sllr SimilarityLingletsRequest) (SimilarityLingletsResponse, error) {
	err := validateSimilarityLingletsRequest(sllr)
	if err != nil {
		log.Print("Malformed request!")
		return SimilarityLingletsResponse{}, err
	}

	pMap := make(map[NameValuePair][]Linglet)
	prMap := make(map[[2]string]int)
	iMap := make(map[string]Linglet)
	linglets := make([]string, len(sllr.Linglets))

	// pass group then lings into query args
	qargs := make([]interface{}, len(sllr.Linglets)+1)
	qargs[0] = sllr.Group
	for i, id := range sllr.Linglets {
		qargs[i+1] = id
	}

	// SELECT lings and properties
	stmt := `SELECT lings.id, lings.name, properties.id, properties.name, lings_properties.value FROM lings INNER JOIN lings_properties ON lings.id=lings_properties.ling_id INNER JOIN properties ON lings_properties.property_id=properties.id WHERE lings.group_id = ? AND lings.depth = 1 AND lings.id IN (?` + strings.Repeat(",?", len(sllr.Linglets)-1) + `)`

	ls, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return SimilarityLingletsResponse{}, err
	}
	defer ls.Close()

	i := 0

	for ls.Next() {
		var l Linglet
		var nv NameValuePair

		err = ls.Scan(&l.Id, &l.Name, &nv.Id, &nv.Name, &nv.Value)
		if err != nil {
			log.Print("Error executing database request!")
			return SimilarityLingletsResponse{}, err
		}

		if ls, ok := pMap[nv]; ok {
			pMap[nv] = append(ls, l)
		} else {
			pMap[nv] = []Linglet{l}
		}

		if _, ok := iMap[l.Id]; !ok {
			linglets[i] = l.Name
			iMap[l.Id] = l

			i++
		}
	}

	// oof
	for _, ls := range pMap {
		for i, l := range ls {
			for j := i + 1; j < len(ls); j++ {
				k := ls[j]

				// deterministic order
				var first Linglet
				var second Linglet
				if l.Id < k.Id {
					first = l
					second = k
				} else {
					first = k
					second = l
				}

				key := [2]string{first.Name, second.Name}
				prMap[key]++
			}
		}
	}

	pairs := make([]SimilarityLingletsResponsePair, len(prMap))
	j := 0

	for pr, c := range prMap {
		ps := make([]string, 2)
		copy(ps, pr[:])

		pairs[j] = SimilarityLingletsResponsePair{
			Linglets:             ps,
			CommonPropertyValues: c,
		}

		j++
	}

	return SimilarityLingletsResponse{
		Type:     "similarity",
		On:       "linglets",
		Linglets: linglets[:i],
		Pairs:    pairs,
	}, nil
}
