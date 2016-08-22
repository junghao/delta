package main

/*
 *  WARNING: CODE GENERATED AUTOMATICALLY.
 *
 *  To update: edit or add yaml file(s) in the config directory.
 *  Commit these changes and run "go generate" in the main project
 *  directory. Changes to this file should then also be commited.
 *
 *  THIS FILE SHOULD NOT BE EDITED BY HAND.
 *
 */

var preparedBy = "Elisabetta D'Anastasio"

var primaryDatacentre = `ftp.geonet.org.nz`
var urlForMoreInformation = `www.geonet.org.nz`
var extraNotes = `additional information and pictures could be
found at http://magma.geonet.org.nz/delta/app
then search for CGPS mark`

var contactAgency = Agency{
	Agency:                `GNS Science`,
	PreferredAbbreviation: `GNS`,
	MailingAddress: `1 Fairway Drive, Avalon 5010,
PO Box 30-368, Lower Hutt
New Zealand`,
	PrimaryContact: Contact{
		Name:               `GeoNet reception`,
		TelephonePrimary:   `+64 4 570 1444`,
		TelephoneSecondary: ``,
		Fax:                `+64 4 570 4676`,
		Email:              `info@geonet.org.nz`,
	},
	SecondaryContact: Contact{
		Name:               `Elisabetta D'Anastasio`,
		TelephonePrimary:   `+64 4 570 4744`,
		TelephoneSecondary: ``,
		Fax:                ``,
		Email:              `e.danastasio@gns.cri.nz`,
	},
	Notes: ``,
}

var responsibleAgency = Agency{
	Agency:                `Land Information New Zealand`,
	PreferredAbbreviation: `LINZ`,
	MailingAddress:        `155 The Terrace, PO Box 5501, Wellington 6145 New Zealand`,
	PrimaryContact: Contact{
		Name:               `LINZ Reception`,
		TelephonePrimary:   `+64 4 460 0110`,
		TelephoneSecondary: ``,
		Fax:                `+64 4 472 2244`,
		Email:              `positionz@linz.govt.nz`,
	},
	SecondaryContact: Contact{
		Name:               `Paula Gentle`,
		TelephonePrimary:   `+64 4 460 2757`,
		TelephoneSecondary: ``,
		Fax:                ``,
		Email:              `pgentle@linz.govt.nz`,
	},
	Notes: `CGPS site is part of the LINZ PositioNZ Network http://www.linz.govt.nz/positionz`,
}

var countryList = []struct {
	name     string
	lat, lon float64
}{
	{"New Zealand", -40, 174},
	{"Niue", -19, -169.9},
	{"Samoa", -13.8, -172.1},
	{"Tonga", -21.2, -175.2},
}

var antennaGraphs = map[string]string{
	"AOAD/M_T":       "202020202020202020202020202020202020202020202020202d2d2d2d2d0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130323020205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303338300a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033353020204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a202020202020202020202020202020202020203d7c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020205258433d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333831302020202020202020202020202020202020202020202d2d3e0a",
	"ASH700936B_M":   "202020202020202020202020202020202020202020202020202d2d2d2d2d0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130323020205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303338300a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033353020204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a202020202020202020202020202020202020203d7c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333831302020202020202020202020202020202020202020202d2d3e0a",
	"ASH700936C_M":   "202020202020202020202020202020202020202020202020202d2d2d2d2d0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130323020205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303338300a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033353020204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a202020202020202020202020202020202020203d7c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333831302020202020202020202020202020202020202020202d2d3e0a",
	"ASH700936D_M":   "202020202020202020202020202020202020202020202020202d2d2d2d2d0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130323020205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303338300a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033353020204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a202020202020202020202020202020202020203d7c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333831302020202020202020202020202020202020202020202d2d3e0a",
	"ASH700936E":     "202020202020202020202020202020202020202020202b2d2d2d2d2d2d2d2d2d2b0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130303820205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303337380a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033343820204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333739342020202020202020202020202020202020202020202d2d3e0a",
	"ASH701933C_M":   "202020202020202020202020202020202020202020202b2d2d2d2d2d2d2d2d2d2b0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130303820205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303337380a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033343820204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333739342020202020202020202020202020202020202020202d2d3e0a",
	"ASH701945B_M":   "20202020202020202020202020202020202020202020202b2d2d2d2d2d2d2d2b0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130303620205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303337360a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033343620204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333739342020202020202020202020202020202020202020202d2d3e0a",
	"ASH701945E_M":   "20202020202020202020202020202020202020202020202b2d2d2d2d2d2d2d2b0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130303620205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303337360a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033343620204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333739342020202020202020202020202020202020202020202d2d3e0a",
	"LEIAT504":       "20202020202020202020202020202020202020202020202b2d2d2d2d2d2d2d2b20202020202020202020202020202020202020202020202020203c2d2d2020302e313339330a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130313220205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303337380a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033343520204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333739342020202020202020202020202020202020202020202d2d3e0a",
	"TRM22020.00+GP": "20202020202020202020202020202020202020202020202d2d2d2d5e2d2d2d2d0a20202020202020202020202020202020202020202f202020202020202020202020205c0a2b2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2b2020203c2d2d2020302e3035393120205447500a2b2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2b2020203c2d2d2020302e3035353620204247500a20202020202020202020202020202020207c202020202020202020202020202020202020207c0a2020202020202020202020202020202020207c20202020202020202020202020202020207c0a202020202020202020202020202020202020207c2020202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a203c2d2d202020202020202020202020202020202020202020302e3436363820202020202020202020202020202020202020202d2d3e20202020202020202020202020202020204e6f74636865730a3c2d2d20202020202020202020202020202020202020202020302e343832362020202020202020202020202020202020202020202d2d3e20202020202020202020202020202020456467650a",
	"TRM29659.00":    "202020202020202020202020202020202020202020202020202d2d2d2d2d0a2020202020202020202020202020202020202020202f20202020202020202020205c0a20202020202020202020202020202020202020207c202020202020202020202020207c0a20202b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b20202020203c2d2d2020302e3130323020205443520a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a20207c202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020207c0a2b2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2b2020203c2d2d2020302e303338310a2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2020203c2d2d2020302e3033353020204243520a20202020202020202020202020202020202020207c202020202020202020202020207c0a202020202020202020202020202020202020203d7c202020202020202020202020207c0a20202020202020202020202020202020202020202b2d2d2d2d2d2d782d2d2d2d2d2d2b20202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020205258433d4e52500a3c2d2d20202020202020202020202020202020202020202020302e333831302020202020202020202020202020202020202020202d2d3e0a",
	"TRM33429.20+GP": "20202020202020202020202020202020202020202020202d2d2d2d5e2d2d2d2d0a20202020202020202020202020202020202020202f202020202020202020202020205c0a2b2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2b2020203c2d2d2020302e3036333120205447500a2b2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2b2b2020203c2d2d2020302e3035393620204247500a20202020202020202020202020202020207c202020202020202020202020202020202020207c0a20202020202020202020202020202020207c202020202020202020202020202020202020207c0a2020202020202020202020202020202020207c20202020202020202020202020202020207c0a202020202020202020202020202020202020202b2d2d2d2d2d2d2d782d2d2d2d2d2d2d2b202020202020202020202020202020202020202020203c2d2d2020302e3030303020204250413d4152500a20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020204e4f4d3d4e52500a203c2d2d202020202020202020202020202020202020202020302e3436363820202020202020202020202020202020202020202d2d3e20202020202020202020202020202020204e6f74636865730a3c2d2d20202020202020202020202020202020202020202020302e343832362020202020202020202020202020202020202020202d2d3e20202020202020202020202020202020456467650a",
	"TRM41249.00":    "202f202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d205c0a2b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202b0a205c202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d202f0a20202020202020202020202020202020205c2020202020202020202020202020202f0a2020202020202020202020202020202020205c202020202020202020202020202f0a202020202020202020202020202020202020205c2d2d2d2d2d782d2d2d2d2d2f2020202020202020202020202020202020202020202020203c2d2d2020302e30303030202042414d3d4152500a2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020205258433d4e52500a3c2d2d2020202020202020202020202020202020202020302e33333936202020202020202020202020202020202020202d2d3e2020202020202020202020202020202020204e6f74636865730a",
	"TRM55971.00":    "202f202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d205c0a2b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202b0a205c202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d202f0a20202020202020202020202020202020205c2020202020202020202020202020202f0a2020202020202020202020202020202020205c202020202020202020202020202f0a202020202020202020202020202020202020205c2d2d2d2d2d782d2d2d2d2d2f2020202020202020202020202020202020202020202020203c2d2d2020302e30303030202042414d3d4152500a2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020205258433d4e52500a3c2d2d2020202020202020202020202020202020202020302e33333936202020202020202020202020202020202020202d2d3e2020202020202020202020202020202020204e6f74636865730a0a",
	"TRM57971.00":    "202f202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d205c0a2b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202b0a205c202d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d202f0a20202020202020202020202020202020205c2020202020202020202020202020202f0a2020202020202020202020202020202020205c202020202020202020202020202f0a202020202020202020202020202020202020205c2d2d2d2d2d782d2d2d2d2d2f2020202020202020202020202020202020202020202020203c2d2d2020302e30303030202042414d3d4152500a2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020205258433d4e52500a3c2d2d2020202020202020202020202020202020202020302e33333936202020202020202020202020202020202020202d2d3e2020202020202020202020202020202020204e6f74636865730a",
}
