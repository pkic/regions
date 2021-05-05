package regions

import (
	"regexp"
	"strings"
)

var (
	// Some values contain meta data that we don't want
	metaDataFormat = regexp.MustCompile(`\s{2,}|\*|\([^\)]*\)|\[[^\]]*\]`)

	// Initially generated with Google Translate with English as the source language
	// - Replaced |) with )
	// - French (fr) [Added 'en']
	// - Polish (pl) [Added 'Prowincja']
	// - Portuguese (pt) [Added 'do']
	// - Romanian (ro) [Added 'a']
	// - Slovenian (sl) [Changed Province to Provinca]
	// - China (zh) [Added 'Sheng', 'Shi', and 'SAR']
	// - Corsican (co) [Removed space in 'l ']
	localPreOrSuffix = []*regexp.Regexp{
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|Region|Capital city|Capital region|City|Territory|of|the)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                              // English (en)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincie|Regio|Hoofdstad|Hoofdstedelijk gewest|stad|Grondgebied|van|de)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                        // Dutch (nl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(provinsie|streek|Hoofstad|streek kapitaal|Stad|gebied|van|die)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                  // Afrikaans (af)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(krahinë|Rajon|Kryeqytet|rajonin e kryeqytetit|qytet|territor|të)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                // Albanian (sq)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ክፍለ ሀገር|ክልል|ዋና ከተማ|ካፒታል ክልል|ከተማ|ግዛት|የ|የ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                        // Amharic (am)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(المحافظة|منطقة|العاصمة|منطقة العاصمة|مدينة|منطقة|من|ال)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                         // Arabic (ar)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(նահանգ|մարզ|Մայրաքաղաք|մայրաքաղաքը մարզ|քաղաք|տարածքը|Հյուրատետր|որ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                            // Armenian (hy)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(vilayət|rayon|Paytaxt|Capital region|şəhər|ərazi|of|the)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                        // Azerbaijani (az)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Probintzia|Region|Hiriburua|Capital eskualdean|hiria|Lurraldea|of|du)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Basque (eu)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(правінцыя|вобласць|Сталіца|сталічны рэгіён|горад|тэрыторыя|з)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                   // Belarusian (be)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(প্রদেশ|অঞ্চল|রাজধানী শহর|ক্যাপিটাল অঞ্চল|শহর|এলাকা|এর|দ্য)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                      // Bengali (Bangla) (bn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincija|regija|Glavni grad|Capital regija|grad|teritorija|od|u)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                              // Bosnian (bs)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(провинция|област|Столица|област Капитал|град|Територия|на|на)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                   // Bulgarian (bg)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(နယ်ပယ်|ဒေသ|မြို့တော်|မြို့တော်ဒေသ|မြို့|နယ်မြေတွေကို|၏|အဆိုပါ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                  // Burmese (my)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(província|regió|Ciutat capital|regió de la capital|ciutat|territori|de|la)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                      // Catalan (ca)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|Region|likulu|Capital dera|City|gawo|wa|ndi)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                           // Chichewa, Chewa, Nyanja (ny)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(省|地区|首都|首都地区|城市|领土|的|这|Sheng|Shi|SAR)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                           // Chinese (zh)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(省|地区|首都|首都地区|城市|领土|的|这|Sheng|Shi|SAR)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                           // Chinese (Simplified) (zh-Hans)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(省|地區|首都|首都地區|城市|領土|的|這|Sheng|Shi|SAR)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                           // Chinese (Traditional) (zh-Hant)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Pruvincia|Region|a cità capitale|righjoni Capital|cità|Territory|di|l')(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                         // Corsican (co)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Pokrajina|Regija|Glavni grad|Capital regija|Grad|Teritorija|od)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                 // Croatian (hr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincie|Kraj|Hlavní město|Capital region|Město|Území|z)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                       // Czech (cs)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provins|Område|Hovedstad|region hovedstaden|by|Territorium|af|det)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                              // Danish (da)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provinco|Regiono|Ĉefurbo|ĉefurbo regiono|Urbo|teritorio|el|la)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                  // Esperanto (eo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(provints|piirkond|Pealinn|Capital piirkonnas|linn|territoorium|kohta)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Estonian (et)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(maakunta|alue|Pääkaupunki|pääkaupunkiseutu|Kaupunki|Alue|of)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Finnish (fi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|Région|Capitale|région de la capitale|Ville|Territoire|de|les|en)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                      // French (fr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(provincia|rexión|capital|rexión da capital|cidade|territorio|de|o)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                              // Galician (gl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Mòr-roinn|Region|prìomh-bhaile|calpa roinn|City|Territory|de|a ')(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                               // Gaelic (Scottish) (gd)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(პროვინციაში|რეგიონი|Დედაქალაქი|დედაქალაქის რეგიონი|ქალაქი|ტერიტორია|საქართველოს)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                // Georgian (ka)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provinz|Region|Hauptstadt|Hauptstadtregion|Stadt|Gebiet|von|das)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                // German (de)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Επαρχία|Περιφέρεια|Πρωτεύουσα|περιοχή της πρωτεύουσας|Πόλη|Εδαφος|του|ο)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                        // Greek (el)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(પ્રાંત|પ્રદેશ|રાજધાની શહેર|કેપિટલ પ્રદેશ|શહેરનું|ટેરિટરી|ના|આ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                  // Gujarati (gu)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Pwovens|rejyon|Kapital vil|rejyon Kapital|City|Teritwa|nan|nan)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                 // Haitian Creole (ht)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(lardin|yankin|babban birnin|babban birnin yankin|City|Territory|na|da)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                          // Hausa (ha)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(מָחוֹז|אזור|עיר בירה|אזור קפיטל|עִיר|שֶׁטַח|שֶׁל|ה)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                             // Hebrew (he)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(प्रांत|क्षेत्र|राजधानी|राजधानी क्षेत्र|Faridabad|क्षेत्र|का)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Hindi (hi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Tartomány|Vidék|Főváros|Főváros régióban|Város|Terület|nak,-nek|az)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Hungarian (hu)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|Region|Höfuðborg|Capital Region|Borg|Territory|af|sem)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                 // Icelandic (is)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|Region|Capital obodo|Capital region|City|Territory|nke|na)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Igbo (ig)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Cúige|Réigiún|Príomhchathair|réigiún caipitil|Cathair|Críoch|de|an)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Irish (ga)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincia|Regione|Capitale|regione della capitale|Città|Territorio|di|il)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                       // Italian (it)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(州|領域|首都|首都圏|市|地域|の|インクルード)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                                      // Japanese (ja)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Langkawi|Region|kutha Capital|wilayah Capital|City|Territory|saka|ing)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                          // Javanese (jv)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ಪ್ರಾಂತ್ಯ|ಪ್ರದೇಶ|ರಾಜಧಾನಿ|ರಾಜಧಾನಿ ಪ್ರದೇಶ|ಸಿಟಿ|ಟೆರಿಟರಿ|ಆಫ್|ದಿ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                   // Kannada (kn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(провинция|аймақ|Астана|капитал облысы|қала|аумақ|туралы|The)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Kazakh (kk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ខេត្ត|តំបន់|រាជធានី|តំបន់រដ្ឋធានី|ទីក្រុង|ដែនដី|នៃ|នេះ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                         // Khmer (km)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Intara|Region|umurwa mukuru|Capital karere|Umugi|Territory|ya|mu)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                               // Kinyarwanda (Rwanda) (rw)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Облус|регион|борбор шаар|Капитал аймак|Сити|территория|боюнча|жана)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Kyrgyz (ky)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(지방|부위|자본 도시|수도권|시티|영토|의|그만큼)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                                    // Korean (ko)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Herêm|Herêm|Paytext|herêmê Capital|Bajar|Herêm|ji|ew)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                           // Kurdish (ku)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ແຂວງ|ພາກພື້ນ|ນະຄອນຫຼວງ|ເຂດນະຄອນຫຼວງ|ເມືອງ|ອານາເຂດຂອງ|ຂອງ|ໄດ້)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                   // Lao (lo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|regio|Oppida prima|regionem capitis|Urbs|Territorium|autem|quod)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                       // Latin (la)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(province|Novads|Galvaspilsēta|Capital reģions|pilsēta|teritorija|no)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                            // Latvian (Lettish) (lv)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(provincija|regionas|Sostinė|Capital Region|miestas|teritorija|apie)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Lithuanian (lt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provënz|Regioun|Haaptstad|Capital Regioun|City|Territoire|vun|der)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                              // Luxembourgish (lb)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(провинција|регионот|Главен град|капитал регион|Сити|територија|на|на)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Macedonian (mk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(-tokony eran'ny fanjakana|Region|Renivohitra|Capital faritra|City|FARITANY|ny|ny)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),               // Malagasy (mg)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(wilayah|rantau|Ibu negeri|Capital Region|City|wilayah|daripada|yang)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                            // Malay (ms)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(പ്രവിശ്യ|പ്രദേശം|തലസ്ഥാന നഗരം|തലസ്ഥാന|നഗരം|ടെറിട്ടറി|എന്ന|The)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                  // Malayalam (ml)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(provinċja|reġjun|Belt kapitali|reġjun kapitali|belt|territorju|ta|il)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Maltese (mt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Porowini|Region|Capital pa|Capital rohe|City|Territory|o|te)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Maori (mi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(प्रांत|प्रदेश|राजधानी|कॅपिटल प्रदेश|सिटी|प्रदेश|च्या|अगोदर निर्देश केलेल्या बाबीसंबंधी बोलताना)(?:\A|\z|\s)[\s\.\-\,\;]*)+`), // Marathi (mr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincie|Regiune|Capitala|regiunea de capital|Oraș|Teritoriu|de)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                               // Moldavian (mo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(аймгийн|бүс нутаг|Нийслэл хот|Нийслэлийн бүс|хот|газар нутаг|нь|The)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                            // Mongolian (mn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(प्रान्त|क्षेत्र|राजधानी|राजधानी क्षेत्र|शहर|इलाका|को|को)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                        // Nepali (ne)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provins|Region|Hovedstad|Capital region|By|Territory|av|de)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                     // Norwegian (no)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provins|Region|Hovedstad|Capital region|By|Territory|av|de)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                     // Norwegian bokmål (nb)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ପ୍ରେଦଶ|ଅଞ୍ଚଳ|ରାଜଧାନୀ ସହର|Capital ଅଞ୍ଚଳ|ସହର|ଭୂଭାଗ|ର|େଯମାେନ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                      // Oriya (or)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ولایت|سیمه|پلازمېنه ښار|پلازمیینه سیمه|ښار|خاوره|د|د)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                           // Pashto, Pushto (ps)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(استان|منطقه|پایتخت|منطقه پایتخت|شهرستان|قلمرو|از)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                               // Persian (Farsi) (fa)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Województwo|Region|Stolica|Region stołeczny|Miasto|Terytorium|z|Prowincja)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                      // Polish (pl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Província|Região|Capital|região da capital|Cidade|Território|de|a|do)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Portuguese (pt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ਸੂਬੇ|ਖੇਤਰ|ਰਾਜਧਾਨੀ|ਰਾਜਧਾਨੀ ਖੇਤਰ|ਸਿਟੀ|ਟੈਰੀਟਰੀ|ਦੇ|ਇਹ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                              // Punjabi (Eastern) (pa)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincie|Regiune|Capitala|regiunea de capital|Oraș|Teritoriu|de|a)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Romanian (ro)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Провинция|Область|Столица|Столичный регион|Город|территория|из)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                 // Russian (ru)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(setate|vāega|laumua|Tupe Faavae itulagi|aʻai|atunuʻu|a|le)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                      // Samoan (sm)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(провинција|регија|Главни град|Цапитал region|град|територија|од)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                // Serbian (sr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|sebakeng|motse-moholo|motse-moholo oa sebaka|City|Territory|ya|ho)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                     // Sesotho (st)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(dunhu|nharaunda|Guta guru|Capital nharaunda|guta|nzvimbo|pamusoro|ari)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                          // Shona (sn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(صوبي|علائقي|گادي جو شھر|گاديء واري علائقي|شهر|سڏجي ٿو|جي|جي)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Sindhi (sd)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(පළාත|කලාපයේ|අග නගරය|අගනුවර කලාපයේ|නගරය|භූමිය|වල|එම)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                             // Sinhalese (si)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincie|kraj|Hlavné mesto|capital región|veľkomesto|územie|z)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                 // Slovak (sk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provinca|regija|Glavno mesto|Capital regija|Kraj|ozemlje|za)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Slovenian (sl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Gobolka|gobolka|magaalada Capital|gobolka Capital|City|Territory|of|ah)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                         // Somali (so)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provincia|Región|Ciudad capital|región de la capital|Ciudad|Territorio|de|la)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                   // Spanish (es)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Propinsi|daerah|kota Capital|wewengkon ibukota|kota|wewengkon|ti|éta)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Sundanese (su)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Mkoa|Mkoa|Mji mkuu|Capital kanda|City|Wilaya|ya|the)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                            // Swahili (Kiswahili) (sw)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provins|Område|Huvudstad|kapital region|Stad|Territorium|av|de)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                 // Swedish (sv)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(lalawigan|Rehiyon|capital city|Capital region|lungsod|lupain|ng|ang)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                            // Tagalog (tl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(вилоят|вилоят|Пойтахт|минтақаи Пойтахт|шаҳр|қаламрави|аз|ба)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                    // Tajik (tg)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(மாகாணம்|பகுதி|தலை நாகரம்|தலைநகர பிராந்தியம்|நகரம்|பிரதேசம்|இன்|தி)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                              // Tamil (ta)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(провинция|өлкә|Башкала|капитал районы|шәһәр|территория|һәм|бу)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                  // Tatar (tt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(ప్రావిన్స్|ప్రాంతం|రాజధాని నగరం|రాజధాని ప్రాంతం|నగరం|భూభాగం|ఆఫ్|ది)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                             // Telugu (te)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(จังหวัด|ภูมิภาค|เมืองหลวง|ภูมิภาคทุน|เมือง|อาณาเขต|ของ)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                         // Thai (th)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Bölge|bölge|Başkent|Sermaye bölge|Kent|bölge|nın-nin)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                           // Turkish (tr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(welaýat|sebit|Paýtagt|Capital sebit|şäher|meýdany|we|mälimlik görkeziji artikl)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                 // Turkmen (tk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(өлкә|رايون|پايتەخت شەھەر|پايتەخت رايون|شەھەر|территория|نىڭ|ишлитилмәйду)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                       // Uyghur (ug)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(провінція|область|Столиця|столичний регіон|місто|територія|з)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                   // Ukrainian (uk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(صوبہ|علاقہ|دارالحکومت|کیپٹل خطے|شہر|علاقہ|کے)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                                   // Urdu (ur)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(viloyat|mintaqa|Poytaxt shahar|Capital viloyati|shahar|hududi|ning|The)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                         // Uzbek (uz)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Tỉnh|Khu vực|Thủ đô|khu vực vốn|thành phố|lãnh thổ|của|các)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                     // Vietnamese (vi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Talaith|rhanbarth|Prifddinas|rhanbarth y Brifddinas|City|Tiriogaeth|o|y)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                        // Welsh (cy)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Provinsje|Regio|Haadstêd|Capital regio|Stêd|Gebiet|fan|de)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                      // Western Frisian (fy)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(kwiPhondo|Region|Isixeko esikhulu|indawo Capital|City|Territory|of|i)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                           // Xhosa (xh)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Province|Region|Olú ìlú|Capital ekun|City|agbegbe|ti|awọn)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                                      // Yoruba (yo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*(?:\A|\z|\s)(Isifundazwe|Isifunda|Inhloko-dolobha|Capital esifundeni|City|Territory|ka|le)(?:\A|\z|\s)[\s\.\-\,\;]*)+`),                   // Zulu (zu)
	}
)

func removeMetaData(name string) string {
	for _, l := range localPreOrSuffix {
		name = l.ReplaceAllLiteralString(name, " ")
	}
	name = metaDataFormat.ReplaceAllLiteralString(name, "")
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
