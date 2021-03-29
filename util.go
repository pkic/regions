package regions

import (
	"regexp"
	"strings"
)

var (
	// Some values contain meta data that we don't want
	metaDataFormat = regexp.MustCompile(`\*|\([^\)]*\)|\[[^\]]*\]`)

	// Initially generated with Google Translate with English as the source language
	localPreOrSuffix = []*regexp.Regexp{
		regexp.MustCompile(`(?i)([-,\s]*(Province|Region|Capital city|Capital region)[-,]*)(\s(of|the))?`),                                 // English (en)
		regexp.MustCompile(`(?i)([-,\s]*(Provincie|Regio|Hoofdstad|Hoofdstedelijk gewest)[-,]*)(\s(van|de))?`),                             // Dutch (nl)
		regexp.MustCompile(`(?i)([-,\s]*(provinsie|streek|Hoofstad|streek kapitaal)[-,]*)(\s(van|die))?`),                                  // Afrikaans (af)
		regexp.MustCompile(`(?i)([-,\s]*(krahinë|Rajon|Kryeqytet|rajonin e kryeqytetit)[-,]*)(\s(të))?`),                                   // Albanian (sq)
		regexp.MustCompile(`(?i)([-,\s]*(ክፍለ ሀገር|ክልል|ዋና ከተማ|ካፒታል ክልል)[-,]*)(\s(የ|የ))?`),                                                    // Amharic (am)
		regexp.MustCompile(`(?i)([-,\s]*(مقاطعة|منطقة|العاصمة|منطقة العاصمة)[-,]*)(\s(من|ال))?`),                                           // Arabic (ar)
		regexp.MustCompile(`(?i)([-,\s]*(նահանգ|մարզ|Մայրաքաղաք|մայրաքաղաքը մարզ)[-,]*)(\s(Հյուրատետր|որ))?`),                              // Armenian (hy)
		regexp.MustCompile(`(?i)([-,\s]*(vilayət|rayon|Paytaxt|Capital region)[-,]*)(\s(of|the))?`),                                        // Azerbaijani (az)
		regexp.MustCompile(`(?i)([-,\s]*(Probintzia|Region|Hiriburua|Capital eskualdean)[-,]*)(\s(of|du))?`),                               // Basque (eu)
		regexp.MustCompile(`(?i)([-,\s]*(правінцыя|вобласць|Сталіца|сталічны рэгіён)[-,]*)(\s(з))?`),                                       // Belarusian (be)
		regexp.MustCompile(`(?i)([-,\s]*(প্রদেশ|অঞ্চল|রাজধানী শহর|ক্যাপিটাল অঞ্চল)[-,]*)(\s(এর|দ্য))?`),                                    // Bengali (Bangla) (bn)
		regexp.MustCompile(`(?i)([-,\s]*(Provincija|regija|Glavni grad|Capital regija)[-,]*)(\s(od|u))?`),                                  // Bosnian (bs)
		regexp.MustCompile(`(?i)([-,\s]*(провинция|област|Столица|област Капитал)[-,]*)(\s(на|на))?`),                                      // Bulgarian (bg)
		regexp.MustCompile(`(?i)([-,\s]*(နယ်ပယ်|ဒေသ|မြို့တော်|မြို့တော်ဒေသ)[-,]*)(\s(၏|အဆိုပါ))?`),                                         // Burmese (my)
		regexp.MustCompile(`(?i)([-,\s]*(província|regió|Ciutat capital|regió de la capital)[-,]*)(\s(de|la))?`),                           // Catalan (ca)
		regexp.MustCompile(`(?i)([-,\s]*(Province|Region|likulu|Capital dera)[-,]*)(\s(wa|ndi))?`),                                         // Chichewa, Chewa, Nyanja (ny)
		regexp.MustCompile(`(?i)([-,\s]*(省|地区|首都|首都地区)[-,]*)(\s(的|这))?`),                                                                   // Chinese (zh)
		regexp.MustCompile(`(?i)([-,\s]*(省|地区|首都|首都地区)[-,]*)(\s(的|这))?`),                                                                   // Chinese (Simplified) (zh-Hans)
		regexp.MustCompile(`(?i)([-,\s]*(省|地區|首都|首都地區)[-,]*)(\s(的|這))?`),                                                                   // Chinese (Traditional) (zh-Hant)
		regexp.MustCompile(`(?i)([-,\s]*(Pruvincia|Region|a cità capitale|righjoni Capital)[-,]*)(\s(di|l '))?`),                           // Corsican (co)
		regexp.MustCompile(`(?i)([-,\s]*(Pokrajina|Regija|Glavni grad|Capital regija)[-,]*)(\s(od))?`),                                     // Croatian (hr)
		regexp.MustCompile(`(?i)([-,\s]*(Provincie|Kraj|Hlavní město|Capital region)[-,]*)(\s(z))?`),                                       // Czech (cs)
		regexp.MustCompile(`(?i)([-,\s]*(Provins|Område|Hovedstad|region hovedstaden)[-,]*)(\s(af|det))?`),                                 // Danish (da)
		regexp.MustCompile(`(?i)([-,\s]*(Provinco|Regiono|Ĉefurbo|ĉefurbo regiono)[-,]*)(\s(el|la))?`),                                     // Esperanto (eo)
		regexp.MustCompile(`(?i)([-,\s]*(provints|piirkond|Pealinn|Capital piirkonnas)[-,]*)(\s(kohta))?`),                                 // Estonian (et)
		regexp.MustCompile(`(?i)([-,\s]*(maakunta|alue|Pääkaupunki|pääkaupunkiseutu)[-,]*)(\s(of))?`),                                      // Finnish (fi)
		regexp.MustCompile(`(?i)([-,\s]*(Province|Région|Capitale|région de la capitale)[-,]*)(\s(de|le))?`),                               // French (fr)
		regexp.MustCompile(`(?i)([-,\s]*(provincia|rexión|capital|rexión da capital)[-,]*)(\s(de|o))?`),                                    // Galician (gl)
		regexp.MustCompile(`(?i)([-,\s]*(Mòr-roinn|Region|prìomh-bhaile|calpa roinn)[-,]*)(\s(de|a '))?`),                                  // Gaelic (Scottish) (gd)
		regexp.MustCompile(`(?i)([-,\s]*(პროვინციაში|რეგიონი|Დედაქალაქი|დედაქალაქის რეგიონი)[-,]*)(\s(საქართველოს))?`),                     // Georgian (ka)
		regexp.MustCompile(`(?i)([-,\s]*(Provinz|Region|Hauptstadt|Hauptstadtregion)[-,]*)(\s(von|das))?`),                                 // German (de)
		regexp.MustCompile(`(?i)([-,\s]*(Επαρχία|Περιφέρεια|Πρωτεύουσα|περιοχή της πρωτεύουσας)[-,]*)(\s(του|ο))?`),                        // Greek (el)
		regexp.MustCompile(`(?i)([-,\s]*(પ્રાંત|પ્રદેશ|રાજધાની શહેર|કેપિટલ પ્રદેશ)[-,]*)(\s(ના|આ))?`),                                      // Gujarati (gu)
		regexp.MustCompile(`(?i)([-,\s]*(Pwovens|rejyon|Kapital vil|rejyon Kapital)[-,]*)(\s(nan|nan))?`),                                  // Haitian Creole (ht)
		regexp.MustCompile(`(?i)([-,\s]*(lardin|yankin|babban birnin|babban birnin yankin)[-,]*)(\s(na|da))?`),                             // Hausa (ha)
		regexp.MustCompile(`(?i)([-,\s]*(מָחוֹז|אזור|עיר בירה|אזור קפיטל)[-,]*)(\s(שֶׁל|ה))?`),                                             // Hebrew (he)
		regexp.MustCompile(`(?i)([-,\s]*(प्रांत|क्षेत्र|राजधानी|राजधानी क्षेत्र)[-,]*)(\s(का))?`),                                          // Hindi (hi)
		regexp.MustCompile(`(?i)([-,\s]*(Tartomány|Vidék|Főváros|Főváros régióban)[-,]*)(\s(nak,-nek|a))?`),                                // Hungarian (hu)
		regexp.MustCompile(`(?i)([-,\s]*(Province|Region|Höfuðborg|Capital Region)[-,]*)(\s(af|sem))?`),                                    // Icelandic (is)
		regexp.MustCompile(`(?i)([-,\s]*(Province|Region|Capital obodo|Capital region)[-,]*)(\s(nke|na))?`),                                // Igbo (ig)
		regexp.MustCompile(`(?i)([-,\s]*(Cúige|Réigiún|Príomhchathair|réigiún caipitil)[-,]*)(\s(de|an))?`),                                // Irish (ga)
		regexp.MustCompile(`(?i)([-,\s]*(Provincia|Regione|Capitale|regione della capitale)[-,]*)(\s(di|il))?`),                            // Italian (it)
		regexp.MustCompile(`(?i)([-,\s]*(州|領域|首都|首都圏)[-,]*)(\s(の|インクルード))?`),                                                               // Japanese (ja)
		regexp.MustCompile(`(?i)([-,\s]*(Langkawi|Region|kutha Capital|wilayah Capital)[-,]*)(\s(saka|ing))?`),                             // Javanese (jv)
		regexp.MustCompile(`(?i)([-,\s]*(ಪ್ರಾಂತ್ಯ|ಪ್ರದೇಶ|ರಾಜಧಾನಿ|ರಾಜಧಾನಿ ಪ್ರದೇಶ)[-,]*)(\s(ಆಫ್|ದಿ))?`),                                    // Kannada (kn)
		regexp.MustCompile(`(?i)([-,\s]*(провинция|аймақ|Астана|капитал облысы)[-,]*)(\s(туралы|The))?`),                                   // Kazakh (kk)
		regexp.MustCompile(`(?i)([-,\s]*(ខេត្ត|តំបន់|រាជធានី|តំបន់រដ្ឋធានី)[-,]*)(\s(នៃ|នេះ))?`),                                           // Khmer (km)
		regexp.MustCompile(`(?i)([-,\s]*(Intara|Region|umurwa mukuru|Capital karere)[-,]*)(\s(ya|mu))?`),                                   // Kinyarwanda (Rwanda) (rw)
		regexp.MustCompile(`(?i)([-,\s]*(Облус|регион|борбор шаар|Капитал аймак)[-,]*)(\s(боюнча|жана))?`),                                 // Kyrgyz (ky)
		regexp.MustCompile(`(?i)([-,\s]*(지방|부위|자본 도시|수도권)[-,]*)(\s(의|그만큼))?`),                                                              // Korean (ko)
		regexp.MustCompile(`(?i)([-,\s]*(Herêm|Herêm|Paytext|herêmê Capital)[-,]*)(\s(ji|ew))?`),                                           // Kurdish (ku)
		regexp.MustCompile(`(?i)([-,\s]*(ແຂວງ|ພາກພື້ນ|ນະຄອນຫຼວງ|ເຂດນະຄອນຫຼວງ)[-,]*)(\s(ຂອງ|ໄດ້))?`),                                        // Lao (lo)
		regexp.MustCompile(`(?i)([-,\s]*(Province|regio|Oppida prima|regionem capitis)[-,]*)(\s(autem|quod))?`),                            // Latin (la)
		regexp.MustCompile(`(?i)([-,\s]*(province|Novads|Galvaspilsēta|Capital reģions)[-,]*)(\s(no))?`),                                   // Latvian (Lettish) (lv)
		regexp.MustCompile(`(?i)([-,\s]*(provincija|regionas|Sostinė|Capital Region)[-,]*)(\s(apie))?`),                                    // Lithuanian (lt)
		regexp.MustCompile(`(?i)([-,\s]*(Provënz|Regioun|Haaptstad|Capital Regioun)[-,]*)(\s(vun|der))?`),                                  // Luxembourgish (lb)
		regexp.MustCompile(`(?i)([-,\s]*(провинција|регионот|Главен град|капитал регион)[-,]*)(\s(на|на))?`),                               // Macedonian (mk)
		regexp.MustCompile(`(?i)([-,\s]*(-tokony eran'ny fanjakana|Region|Renivohitra|Capital faritra)[-,]*)(\s(ny|ny))?`),                 // Malagasy (mg)
		regexp.MustCompile(`(?i)([-,\s]*(wilayah|rantau|Ibu negeri|Capital Region)[-,]*)(\s(daripada|yang))?`),                             // Malay (ms)
		regexp.MustCompile(`(?i)([-,\s]*(പ്രവിശ്യ|പ്രദേശം|തലസ്ഥാന നഗരം|തലസ്ഥാന)[-,]*)(\s(എന്ന|The))?`),                                     // Malayalam (ml)
		regexp.MustCompile(`(?i)([-,\s]*(provinċja|reġjun|Belt kapitali|reġjun kapitali)[-,]*)(\s(ta|il))?`),                               // Maltese (mt)
		regexp.MustCompile(`(?i)([-,\s]*(Porowini|Region|Capital pa|Capital rohe)[-,]*)(\s(o|te))?`),                                       // Maori (mi)
		regexp.MustCompile(`(?i)([-,\s]*(प्रांत|प्रदेश|राजधानी|कॅपिटल प्रदेश)[-,]*)(\s(च्या|अगोदर निर्देश केलेल्या बाबीसंबंधी बोलताना))?`), // Marathi (mr)
		regexp.MustCompile(`(?i)([-,\s]*(Provincie|Regiune|Capitala|regiunea de capital)[-,]*)(\s(de))?`),                                  // Moldavian (mo)
		regexp.MustCompile(`(?i)([-,\s]*(аймгийн|бүс нутаг|Нийслэл хот|Нийслэлийн бүс)[-,]*)(\s(нь|The))?`),                                // Mongolian (mn)
		regexp.MustCompile(`(?i)([-,\s]*(प्रान्त|क्षेत्र|राजधानी|राजधानी क्षेत्र)[-,]*)(\s(को|को))?`),                                      // Nepali (ne)
		regexp.MustCompile(`(?i)([-,\s]*(Provins|Region|Hovedstad|Capital region)[-,]*)(\s(av|de))?`),                                      // Norwegian (no)
		regexp.MustCompile(`(?i)([-,\s]*(Provins|Region|Hovedstad|Capital region)[-,]*)(\s(av|de))?`),                                      // Norwegian bokmål (nb)
		regexp.MustCompile(`(?i)([-,\s]*(ପ୍ରେଦଶ|ଅଞ୍ଚଳ|ରାଜଧାନୀ ସହର|Capital ଅଞ୍ଚଳ)[-,]*)(\s(ର|େଯମାେନ))?`),                                    // Oriya (or)
		regexp.MustCompile(`(?i)([-,\s]*(ولایت|سیمه|پلازمېنه ښار|پلازمیینه سیمه)[-,]*)(\s(د|د))?`),                                         // Pashto, Pushto (ps)
		regexp.MustCompile(`(?i)([-,\s]*(استان|منطقه|پایتخت|منطقه پایتخت)[-,]*)(\s(از))?`),                                                 // Persian (Farsi) (fa)
		regexp.MustCompile(`(?i)([-,\s]*(Województwo|Region|Stolica|Region stołeczny)[-,]*)(\s(z))?`),                                      // Polish (pl)
		regexp.MustCompile(`(?i)([-,\s]*(Província|Região|Capital|região da capital)[-,]*)(\s(de|a))?`),                                    // Portuguese (pt)
		regexp.MustCompile(`(?i)([-,\s]*(ਸੂਬੇ|ਖੇਤਰ|ਰਾਜਧਾਨੀ|ਰਾਜਧਾਨੀ ਖੇਤਰ)[-,]*)(\s(ਦੇ|ਇਹ))?`),                                               // Punjabi (Eastern) (pa)
		regexp.MustCompile(`(?i)([-,\s]*(Provincie|Regiune|Capitala|regiunea de capital)[-,]*)(\s(de))?`),                                  // Romanian (ro)
		regexp.MustCompile(`(?i)([-,\s]*(Провинция|Область, край|Столица|Столичный регион)[-,]*)(\s(из))?`),                                // Russian (ru)
		regexp.MustCompile(`(?i)([-,\s]*(setate|vāega|laumua|Tupe Faavae itulagi)[-,]*)(\s(a|le))?`),                                       // Samoan (sm)
		regexp.MustCompile(`(?i)([-,\s]*(провинција|регија|Главни град|Цапитал region)[-,]*)(\s(од))?`),                                    // Serbian (sr)
		regexp.MustCompile(`(?i)([-,\s]*(Province|sebakeng|motse-moholo|motse-moholo oa sebaka)[-,]*)(\s(ya|ho))?`),                        // Sesotho (st)
		regexp.MustCompile(`(?i)([-,\s]*(dunhu|nharaunda|Guta guru|Capital nharaunda)[-,]*)(\s(pamusoro|ari))?`),                           // Shona (sn)
		regexp.MustCompile(`(?i)([-,\s]*(صوبي|علائقي|گادي جو شھر|گاديء واري علائقي)[-,]*)(\s(جي|جي))?`),                                    // Sindhi (sd)
		regexp.MustCompile(`(?i)([-,\s]*(පළාත|කලාපයේ|අග නගරය|අගනුවර කලාපයේ)[-,]*)(\s(වල|එම))?`),                                            // Sinhalese (si)
		regexp.MustCompile(`(?i)([-,\s]*(Provincie|kraj|Hlavné mesto|capital región)[-,]*)(\s(z))?`),                                       // Slovak (sk)
		regexp.MustCompile(`(?i)([-,\s]*(Province|regija|Glavno mesto|Capital regija)[-,]*)(\s(za))?`),                                     // Slovenian (sl)
		regexp.MustCompile(`(?i)([-,\s]*(Gobolka|gobolka|magaalada Capital|gobolka Capital)[-,]*)(\s(of|ah))?`),                            // Somali (so)
		regexp.MustCompile(`(?i)([-,\s]*(Provincia|Región|Ciudad capital|región de la capital)[-,]*)(\s(de|la))?`),                         // Spanish (es)
		regexp.MustCompile(`(?i)([-,\s]*(Propinsi|daerah|kota Capital|wewengkon ibukota)[-,]*)(\s(ti|éta))?`),                              // Sundanese (su)
		regexp.MustCompile(`(?i)([-,\s]*(Mkoa|Mkoa|Mji mkuu|Capital kanda)[-,]*)(\s(ya|the))?`),                                            // Swahili (Kiswahili) (sw)
		regexp.MustCompile(`(?i)([-,\s]*(Provins|Område|Huvudstad|kapital region)[-,]*)(\s(av|de))?`),                                      // Swedish (sv)
		regexp.MustCompile(`(?i)([-,\s]*(lalawigan|Rehiyon|capital city|Capital region)[-,]*)(\s(ng|ang))?`),                               // Tagalog (tl)
		regexp.MustCompile(`(?i)([-,\s]*(вилоят|вилоят|Пойтахт|минтақаи Пойтахт)[-,]*)(\s(аз|ба))?`),                                       // Tajik (tg)
		regexp.MustCompile(`(?i)([-,\s]*(மாகாணம்|பகுதி|தலை நாகரம்|தலைநகர பிராந்தியம்)[-,]*)(\s(இன்|தி))?`),                                 // Tamil (ta)
		regexp.MustCompile(`(?i)([-,\s]*(провинция|өлкә|Башкала|капитал районы)[-,]*)(\s(һәм|бу))?`),                                       // Tatar (tt)
		regexp.MustCompile(`(?i)([-,\s]*(ప్రావిన్స్|ప్రాంతం|రాజధాని నగరం|రాజధాని ప్రాంతం)[-,]*)(\s(ఆఫ్|ది))?`),                             // Telugu (te)
		regexp.MustCompile(`(?i)([-,\s]*(จังหวัด|ภูมิภาค|เมืองหลวง|ภูมิภาคทุน)[-,]*)(\s(ของ))?`),                                           // Thai (th)
		regexp.MustCompile(`(?i)([-,\s]*(Bölge|bölge|Başkent|Sermaye bölge)[-,]*)(\s(nın-nin))?`),                                          // Turkish (tr)
		regexp.MustCompile(`(?i)([-,\s]*(welaýat|sebit|Paýtagt|Capital sebit)[-,]*)(\s(we|mälimlik görkeziji artikl))?`),                   // Turkmen (tk)
		regexp.MustCompile(`(?i)([-,\s]*(өлкә|رايون|پايتەخت شەھەر|پايتەخت رايون)[-,]*)(\s(نىڭ|ишлитилмәйду))?`),                            // Uyghur (ug)
		regexp.MustCompile(`(?i)([-,\s]*(провінція|область|Столиця|столичний регіон)[-,]*)(\s(з))?`),                                       // Ukrainian (uk)
		regexp.MustCompile(`(?i)([-,\s]*(صوبہ|علاقہ|دارالحکومت|کیپٹل خطے)[-,]*)(\s(کے))?`),                                                 // Urdu (ur)
		regexp.MustCompile(`(?i)([-,\s]*(viloyat|mintaqa|Poytaxt shahar|Capital viloyati)[-,]*)(\s(ning|The))?`),                           // Uzbek (uz)
		regexp.MustCompile(`(?i)([-,\s]*(Tỉnh|Khu vực|Thủ đô|khu vực vốn)[-,]*)(\s(của|các))?`),                                            // Vietnamese (vi)
		regexp.MustCompile(`(?i)([-,\s]*(Talaith|rhanbarth|Prifddinas|rhanbarth y Brifddinas)[-,]*)(\s(o|y))?`),                            // Welsh (cy)
		regexp.MustCompile(`(?i)([-,\s]*(Provinsje|Regio|Haadstêd|Capital regio)[-,]*)(\s(fan|de))?`),                                      // Western Frisian (fy)
		regexp.MustCompile(`(?i)([-,\s]*(kwiPhondo|Region|Isixeko esikhulu|indawo Capital)[-,]*)(\s(of|i))?`),                              // Xhosa (xh)
		regexp.MustCompile(`(?i)([-,\s]*(Province|Region|Olú ìlú|Capital ekun)[-,]*)(\s(ti|awọn))?`),                                       // Yoruba (yo)
		regexp.MustCompile(`(?i)([-,\s]*(Isifundazwe|Isifunda|Inhloko-dolobha|Capital esifundeni)[-,]*)(\s(ka|le))?`),                      // Zulu (zu)	}
	}
)

func removeMetaData(name string) string {
	name = metaDataFormat.ReplaceAllLiteralString(name, "")
	for _, l := range localPreOrSuffix {
		name = l.ReplaceAllLiteralString(name, "")
	}
	return strings.TrimSpace(name)
}

func stringInSlice(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, str) {
			return true
		}
	}
	return false
}
