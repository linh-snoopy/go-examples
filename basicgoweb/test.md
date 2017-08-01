# Index:
* [Create a new template](#create-a-template)
* [Get detail of specific template](#get-detail-of-template)


### Create a template
```javascript
POST /api/template
{
	"formatVersion": 1,
	"teamIdentifier": "L7JL5WW78J",
	"organizationName": "Mercedes-Benz",
	"passTypeIdentifier": "pass.eu.missmp.testing",
	"subtemplates": [
		{
			"description": "San Francisco to London",
			"logoText": "aaaaaaaaa",
			"backgroundColor": "rgb(0, 0, 0)",
			"foregroundColor": "rgb(255, 255, 255)",
			"labelColor": "rgb(100, 100, 100)",
			"userInfo": {
				"code": "M3309247"
			},
			"layoutId": 1,
			"locales": {
				"en": {
					"primary_label_tag-name": "Event",
					"primary_value_tag-name": "The Hectic Glow in concert",
					"secondary_label_doors-open": "Doors open",
					"secondary_value_doors-open": "2013-08-10T19:30-06:00",
					"secondary_label_seating-section": "Seating section",
					"secondary_value_seating-section": "5",
					"back_label_phone": "Telephone",
					"back_value_phone": "+49 7031 6640",
					"back_label_address": "Contact",
					"back_value_address": "Betrieb Böblingen\nWolf-Hirth Straße 28\n71034 Böblingen",
					"header_label_event-name": "Track my checked bags",
					"header_value_event-name": "http://www.example.com/track-bags/XYZ123"
				},
				"de": {
					"primary_label_tag-name": "Event",
					"primary_value_tag-name": "The Hectic Glow in concert",
					"secondary_label_doors-open": "Doors open",
					"secondary_value_doors-open": "2013-08-10T19:30-06:00",
					"secondary_label_seating-section": "Seating section",
					"secondary_value_seating-section": "5",
					"back_label_phone": "Telefon",
					"back_value_phone": "+49 7031 6640",
					"back_label_address": "Kontakt",
					"back_value_address": "Betrieb Böblingen\nWolf-Hirth Straße 28\n71034 Böblingen",
					"header_label_event-name": "Veranstaltungen",
					"header_value_event-name": "11. und 12. März 2017\nErlebnis Outdoor Kongresshalle Böblingen\nPremiumpartner Tourismusverband Tannheimer Tal - Präsentation SUV-Flotte und V-Klasse Marco Polo"
				}
			},
			"barcodes" : [
		        {
		            "message" : "ABCD 123 EFGH 456 IJKL 789 MNOP",
		            "format" : "PKBarcodeFormatPDF417",
		            "messageEncoding" : "iso-8859-1"
		        }
		    ]
		},
		{
			"description": "San Francisco to London",
			"logoText": "bbbbbbbbbbbbbbb",
			"backgroundColor": "rgb(0, 0, 0)",
			"foregroundColor": "rgb(255, 255, 255)",
			"labelColor": "rgb(100, 100, 100)",
			"userInfo": {
				"code": "M3309247"
			},
			"layoutId": 1,
			"locales": {
				"en": {
					"primary_label_tag-name": "Event",
					"primary_value_tag-name": "The Hectic Glow in concert",
					"secondary_label_doors-open": "Doors open",
					"secondary_value_doors-open": "2013-08-10T19:30-06:00",
					"secondary_label_seating-section": "Seating section",
					"secondary_value_seating-section": "5",
					"back_label_phone": "Telephone",
					"back_value_phone": "+49 7031 6640",
					"back_label_address": "Contact",
					"back_value_address": "Betrieb Böblingen\nWolf-Hirth Straße 28\n71034 Böblingen",
					"header_label_event-name": "Track my checked bags",
					"header_value_event-name": "http://www.example.com/track-bags/XYZ123"
				},
				"de": {
					"primary_label_tag-name": "Event",
					"primary_value_tag-name": "The Hectic Glow in concert",
					"secondary_label_doors-open": "Doors open",
					"secondary_value_doors-open": "2013-08-10T19:30-06:00",
					"secondary_label_seating-section": "Seating section",
					"secondary_value_seating-section": "5",
					"back_label_phone": "Telefon",
					"back_value_phone": "+49 7031 6640",
					"back_label_address": "Kontakt",
					"back_value_address": "Betrieb Böblingen\nWolf-Hirth Straße 28\n71034 Böblingen",
					"header_label_event-name": "Veranstaltungen",
					"header_value_event-name": "11. und 12. März 2017\nErlebnis Outdoor Kongresshalle Böblingen\nPremiumpartner Tourismusverband Tannheimer Tal - Präsentation SUV-Flotte und V-Klasse Marco Polo"
				}
			},
			"barcodes" : [
		        {
		            "message" : "ABCD 123 EFGH 456 IJKL 789 MNOP",
		            "format" : "PKBarcodeFormatPDF417",
		            "messageEncoding" : "iso-8859-1"
		        }
		    ]
		},
		{
			"description": "San Francisco to London",
			"logoText": "cccccccccccccc",
			"backgroundColor": "rgb(0, 0, 0)",
			"foregroundColor": "rgb(255, 255, 255)",
			"labelColor": "rgb(100, 100, 100)",
			"userInfo": {
				"code": "M3309247"
			},
			"layoutId": 1,
			"locales": {
				"en": {
					"primary_label_tag-name": "Event",
					"primary_value_tag-name": "The Hectic Glow in concert",
					"secondary_label_doors-open": "Doors open",
					"secondary_value_doors-open": "2013-08-10T19:30-06:00",
					"secondary_label_seating-section": "Seating section",
					"secondary_value_seating-section": "5",
					"back_label_phone": "Telephone",
					"back_value_phone": "+49 7031 6640",
					"back_label_address": "Contact",
					"back_value_address": "Betrieb Böblingen\nWolf-Hirth Straße 28\n71034 Böblingen",
					"header_label_event-name": "Track my checked bags",
					"header_value_event-name": "http://www.example.com/track-bags/XYZ123"
				},
				"de": {
					"primary_label_tag-name": "Event",
					"primary_value_tag-name": "The Hectic Glow in concert",
					"secondary_label_doors-open": "Doors open",
					"secondary_value_doors-open": "2013-08-10T19:30-06:00",
					"secondary_label_seating-section": "Seating section",
					"secondary_value_seating-section": "5",
					"back_label_phone": "Telefon",
					"back_value_phone": "+49 7031 6640",
					"back_label_address": "Kontakt",
					"back_value_address": "Betrieb Böblingen\nWolf-Hirth Straße 28\n71034 Böblingen",
					"header_label_event-name": "Veranstaltungen",
					"header_value_event-name": "11. und 12. März 2017\nErlebnis Outdoor Kongresshalle Böblingen\nPremiumpartner Tourismusverband Tannheimer Tal - Präsentation SUV-Flotte und V-Klasse Marco Polo"
				}
			}
		}
	]
}
```
---

### Get detail of template
```javascript
GET /api/template/{templateId}/subtemplate/{subtemplateId}
Example: GET /api/template/vtuLbNJYUARH/subtemplate/285
Response:
{
    "id": 285,
    "passTypeId": "pass.eu.missmp.mercedes.testing",
    "templateId": "vtuLbNJYUARH",
    "createdAt": "2017-07-27T07:18:12.313254Z",
    "updatedAt": "2017-07-27T07:18:12.313254Z",
    "passData": {
        "formatVersion": 1,
        "passTypeIdentifier": "pass.eu.missmp.mercedes.testing",
        "serialNumber": "",
        "teamIdentifier": "L7JL5WW78J",
        "webServiceURL": "https://1138315b.ngrok.io/passkit/",
        "authenticationToken": "",
        "organizationName": "Mercedes-Benz",
        "description": "San Francisco to London",
        "logoText": "Version 1",
        "backgroundColor": "rgb(0, 0, 0)",
        "foregroundColor": "rgb(255, 255, 255)",
        "labelColor": "rgb(100, 100, 100)",
        "associatedStoreIdentifiers": [
            1029372196
        ],
        "associatedPlayIdentifiers": [
            "com.daimler.mm.android"
        ],
        "userInfo": {
            "code": "M3309247"
        },
        "eventTicket": {
            "primaryFields": [
                {
                    "key": "tag-name",
                    "label": "primary_label_tag-name",
                    "value": "primary_value_tag-name"
                }
            ],
            "secondaryFields": [
                {
                    "key": "doors-open",
                    "label": "secondary_label_doors-open",
                    "value": "secondary_value_doors-open",
                    "dateStyle": "PKDateStyleMedium",
                    "timeStyle": "PKDateStyleShort"
                },
                {
                    "key": "seating-section",
                    "label": "secondary_label_seating-section",
                    "value": "secondary_value_seating-section",
                    "textAlignment": "PKTextAlignmentRight",
                    "numberStyle": "PKNumberStyleSpellOut"
                }
            ],
            "backFields": [
                {
                    "key": "phone",
                    "label": "back_label_phone",
                    "value": "back_value_phone",
                    "dataDetectorTypes": [
                        "PKDataDetectorTypePhoneNumber"
                    ]
                },
                {
                    "key": "address",
                    "label": "back_label_address",
                    "value": "back_value_address",
                    "dataDetectorTypes": [
                        "PKDataDetectorTypeAddress"
                    ]
                },
                {
                    "key": "URL",
                    "label": "back_label_URL",
                    "value": "back_value_URL",
                    "dataDetectorTypes": [
                        "PKDataDetectorTypeLink"
                    ]
                }
            ],
            "headerFields": [
                {
                    "key": "event-name",
                    "label": "header_label_event-name",
                    "value": "header_value_event-name"
                }
            ]
        }
    },
    "isActive": true,
    "startAt": "2017-07-27T07:18:12.313254Z",
    "endAt": "2017-08-03T07:18:12.448757Z",
    "updateTag": 1501139892,
    "layout": 3,
    "passStyle": "eventTicket",
    "images": {
        "icon": "mercedes/passes/testing/vtuLbNJYUARH/0/icon.png",
        "logo": "mercedes/passes/testing/vtuLbNJYUARH/0/logo.png",
        "strip": "mercedes/passes/testing/vtuLbNJYUARH/0/M5086666_r.jpg"
    }
}
```
