package fetcher

const fixture = `
{
	"status": "ok",
	"code": 0,
	"data": {
		"team": {
			"id": 1,
			"optaId": 479,
			"name": "Apoel FC",
			"logoUrls": [{
				"size": "56x56",
				"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/1.png"
			}, {
				"size": "164x164",
				"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/1.png"
			}],
			"isNational": false,
			"matches": {
				"last": {
					"scoreaway": "1",
					"scorehome": "2",
					"status": "FullTime",
					"id": 504296,
					"competitionId": 7,
					"seasonId": 1709,
					"stadiumId": 24,
					"matchdayId": 5669744,
					"matchday": {
						"id": 5669744
					},
					"kickoff": "2016-09-15T17:00:00Z",
					"minute": 95,
					"teamhome": {
						"idInternal": 1,
						"id": 479,
						"name": "Apoel FC",
						"colors": {
							"shirtColorHome": "0066CC",
							"shirtColorAway": "FF9966",
							"crestMainColor": "4F2C7D",
							"mainColor": "0066CC"
						},
						"logoUrls": [{
							"size": "56x56",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/1.png"
						}, {
							"size": "164x164",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/1.png"
						}]
					},
					"teamaway": {
						"idInternal": 1874,
						"id": 3751,
						"name": "FC Astana",
						"colors": {
							"shirtColorHome": "",
							"shirtColorAway": "",
							"crestMainColor": "2B2667",
							"mainColor": "2B2667"
						},
						"logoUrls": [{
							"size": "56x56",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/1874.png"
						}, {
							"size": "164x164",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/1874.png"
						}]
					}
				},
				"next": {
					"scoreaway": "-1",
					"scorehome": "-1",
					"status": "PreMatch",
					"id": 504335,
					"competitionId": 7,
					"seasonId": 1709,
					"stadiumId": 29,
					"matchdayId": 5669745,
					"matchday": {
						"id": 5669745
					},
					"kickoff": "2016-09-29T19:05:00Z",
					"minute": 0,
					"teamhome": {
						"idInternal": 24,
						"id": 202,
						"name": "Olympiakos",
						"colors": {
							"shirtColorHome": "FFFFFF",
							"shirtColorAway": "003399",
							"crestMainColor": "FF0000",
							"mainColor": "FFFFFF"
						},
						"logoUrls": [{
							"size": "56x56",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/24.png"
						}, {
							"size": "164x164",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/24.png"
						}]
					},
					"teamaway": {
						"idInternal": 1,
						"id": 479,
						"name": "Apoel FC",
						"colors": {
							"shirtColorHome": "0066CC",
							"shirtColorAway": "FF9966",
							"crestMainColor": "4F2C7D",
							"mainColor": "0066CC"
						},
						"logoUrls": [{
							"size": "56x56",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/1.png"
						}, {
							"size": "164x164",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/1.png"
						}]
					}
				},
				"following": {
					"scoreaway": "-1",
					"scorehome": "-1",
					"status": "PreMatch",
					"id": 504345,
					"competitionId": 7,
					"seasonId": 1709,
					"stadiumId": 335,
					"matchdayId": 5669746,
					"matchday": {
						"id": 5669746
					},
					"kickoff": "2016-10-20T19:05:00Z",
					"minute": 0,
					"teamhome": {
						"idInternal": 347,
						"id": 1963,
						"name": "BSC YB",
						"colors": {
							"shirtColorHome": "FF9900",
							"shirtColorAway": "FFFFFF",
							"crestMainColor": "",
							"mainColor": "FF9900"
						},
						"logoUrls": [{
							"size": "56x56",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/347.png"
						}, {
							"size": "164x164",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/347.png"
						}]
					},
					"teamaway": {
						"idInternal": 1,
						"id": 479,
						"name": "Apoel FC",
						"colors": {
							"shirtColorHome": "0066CC",
							"shirtColorAway": "FF9966",
							"crestMainColor": "4F2C7D",
							"mainColor": "0066CC"
						},
						"logoUrls": [{
							"size": "56x56",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/56\/1.png"
						}, {
							"size": "164x164",
							"url": "https:\/\/images.onefootball.com\/icons\/internal\/164\/1.png"
						}]
					}
				}
			},
			"competitions": [{
				"competitionId": 140
			}, {
				"competitionId": 21
			}, {
				"competitionId": 7
			}],
			"players": [{
				"country": "Portugal",
				"id": "6",
				"firstName": "Nuno",
				"lastName": "Morais",
				"name": "Nuno Morais",
				"position": "Midfielder",
				"number": 26,
				"birthDate": "1984-01-29",
				"age": "32",
				"height": 185,
				"weight": 76,
				"thumbnailSrc": "https:\/\/images.onefootball.com\/default\/default_player.png"
			}, {
				"country": "Cyprus",
				"id": "19",
				"firstName": "Nektarious",
				"lastName": "Alexandrou",
				"name": "Nektarious Alexandrou",
				"position": "Midfielder",
				"number": 11,
				"birthDate": "1983-12-19",
				"age": "32",
				"height": 182,
				"weight": 76,
				"thumbnailSrc": "https:\/\/images.onefootball.com\/player\/98\/98bdd1b3e9ba596ffb0d8c09071a0577.jpg"
			}, {
				"country": "Spain",
				"id": "770",
				"firstName": "Urko",
				"lastName": "Pardo",
				"name": "Urko Pardo",
				"position": "Goalkeeper",
				"number": 78,
				"birthDate": "1983-01-28",
				"age": "33",
				"height": 189,
				"weight": 85,
				"thumbnailSrc": "https:\/\/images.onefootball.com\/player\/36\/36a9143ede9200fff4fbae81db38da60.jpg"
			}],
			"officials": [{
				"countryName": "Spain",
				"id": "49381",
				"firstName": "Thomas",
				"lastName": "Christiansen",
				"country": "ES",
				"position": "Coach"
			}],
			"colors": {
				"shirtColorHome": "0066CC",
				"shirtColorAway": "FF9966",
				"crestMainColor": "4F2C7D",
				"mainColor": "0066CC"
			}
		}
	},
	"message": "Team feed successfully generated. Api Version: 1"
}
`
