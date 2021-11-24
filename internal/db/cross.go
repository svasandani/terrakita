package db

import (
	"log"
	"reflect"
	"strings"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
)

func CrossLingProperties(clpr CrossLingPropertiesRequest) (CrossLingPropertiesResponse, error) {
	err := validateCrossLingPropertiesRequest(clpr)
	if err != nil {
		log.Print("Malformed request!")
		return CrossLingPropertiesResponse{}, err
	}

	pnMap := make(map[string]int)
	lMap := make(map[string][]NameValuePair)
	pMap := make(map[[6]NameValuePair][]string)
	properties := make([]string, len(clpr.LingProperties))

	// pass group then ling properties into query args
	qargs := make([]interface{}, len(clpr.LingProperties)+1)
	qargs[0] = clpr.Group
	for i, id := range clpr.LingProperties {
		qargs[i+1] = id
	}

	// SELECT properties
	stmt := `SELECT properties.id, properties.name, lings_properties.value, lings.name FROM lings_properties INNER JOIN properties ON lings_properties.property_id=properties.id INNER JOIN lings ON lings_properties.ling_id=lings.id WHERE lings.group_id = ? AND lings.depth = 0 AND properties.id IN (?` + strings.Repeat(",?", len(clpr.LingProperties)-1) + `)`
	ps, err := db.Query(stmt, qargs...)
	if err != nil {
		log.Print("Error preparing database request!")
		return CrossLingPropertiesResponse{}, err
	}
	defer ps.Close()

	i := 0

	for ps.Next() {
		var nv NameValuePair
		var l Ling

		err = ps.Scan(&nv.Id, &nv.Name, &nv.Value, &l.Name)
		if err != nil {
			log.Print("Error executing database request!")
			return CrossLingPropertiesResponse{}, err
		}

		if _, ok := pnMap[nv.Name]; !ok {
			properties[i] = nv.Name
			pnMap[nv.Name] = 1
			i++
		}

		if n, ok := lMap[l.Name]; ok {
			lMap[l.Name] = append(n, nv)
		} else {
			ns := make([]NameValuePair, 1, 6)
			ns[0] = nv
			lMap[l.Name] = ns
		}
	}

	j := 0

	for l, n := range(lMap) {
		if len(n) == len(clpr.LingProperties) {
			// ew
			s := (*reflect.SliceHeader)(unsafe.Pointer(&n))
			nArr := *(*[6]NameValuePair)(unsafe.Pointer(s.Data))

			if ls, ok := pMap[nArr]; ok {
				pMap[nArr] = append(ls, l)
			} else {
				pMap[nArr] = []string{l}
				j++
			}
		}
	}

	pc := make([]CrossLingPropertiesResponsePropertyCombinations, j)

	k := 0

	for n, ls := range(pMap) {
		ns := make([]NameValuePair, len(clpr.LingProperties))
		copy(ns, n[:len(clpr.LingProperties)])
		pc[k] = CrossLingPropertiesResponsePropertyCombinations{
			PropertyValuePairs: ns,
			Lings: ls,
		}

		k++
	}

	return CrossLingPropertiesResponse{
		Type:     "cross",
		On:       "ling_properties",
		LingProperties:    properties,
		PropertyCombinations: pc,
	}, nil
}