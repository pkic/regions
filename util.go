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
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|Region|Capital city|Capital region|City|Territory|of|the)\b[\s\.\-\,\;]*)+`),                              // English (en)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincie|Regio|Hoofdstad|Hoofdstedelijk gewest|stad|Grondgebied|van|de)\b[\s\.\-\,\;]*)+`),                        // Dutch (nl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(provinsie|streek|Hoofstad|streek kapitaal|Stad|gebied|van|die)\b[\s\.\-\,\;]*)+`),                                  // Afrikaans (af)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(krahinë|Rajon|Kryeqytet|rajonin e kryeqytetit|qytet|territor|të)\b[\s\.\-\,\;]*)+`),                                // Albanian (sq)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ክፍለ ሀገር|ክልል|ዋና ከተማ|ካፒታል ክልል|ከተማ|ግዛት|የ|የ)\b[\s\.\-\,\;]*)+`),                                                        // Amharic (am)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(المحافظة|منطقة|العاصمة|منطقة العاصمة|مدينة|منطقة|من|ال)\b[\s\.\-\,\;]*)+`),                                         // Arabic (ar)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(նահանգ|մարզ|Մայրաքաղաք|մայրաքաղաքը մարզ|քաղաք|տարածքը|Հյուրատետր|որ)\b[\s\.\-\,\;]*)+`),                            // Armenian (hy)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(vilayət|rayon|Paytaxt|Capital region|şəhər|ərazi|of|the)\b[\s\.\-\,\;]*)+`),                                        // Azerbaijani (az)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Probintzia|Region|Hiriburua|Capital eskualdean|hiria|Lurraldea|of|du)\b[\s\.\-\,\;]*)+`),                           // Basque (eu)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(правінцыя|вобласць|Сталіца|сталічны рэгіён|горад|тэрыторыя|з)\b[\s\.\-\,\;]*)+`),                                   // Belarusian (be)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(প্রদেশ|অঞ্চল|রাজধানী শহর|ক্যাপিটাল অঞ্চল|শহর|এলাকা|এর|দ্য)\b[\s\.\-\,\;]*)+`),                                      // Bengali (Bangla) (bn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincija|regija|Glavni grad|Capital regija|grad|teritorija|od|u)\b[\s\.\-\,\;]*)+`),                              // Bosnian (bs)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(провинция|област|Столица|област Капитал|град|Територия|на|на)\b[\s\.\-\,\;]*)+`),                                   // Bulgarian (bg)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(နယ်ပယ်|ဒေသ|မြို့တော်|မြို့တော်ဒေသ|မြို့|နယ်မြေတွေကို|၏|အဆိုပါ)\b[\s\.\-\,\;]*)+`),                                  // Burmese (my)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(província|regió|Ciutat capital|regió de la capital|ciutat|territori|de|la)\b[\s\.\-\,\;]*)+`),                      // Catalan (ca)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|Region|likulu|Capital dera|City|gawo|wa|ndi)\b[\s\.\-\,\;]*)+`),                                           // Chichewa, Chewa, Nyanja (ny)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(省|地区|首都|首都地区|城市|领土|的|这|Sheng|Shi|SAR)\b[\s\.\-\,\;]*)+`),                                                           // Chinese (zh)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(省|地区|首都|首都地区|城市|领土|的|这|Sheng|Shi|SAR)\b[\s\.\-\,\;]*)+`),                                                           // Chinese (Simplified) (zh-Hans)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(省|地區|首都|首都地區|城市|領土|的|這|Sheng|Shi|SAR)\b[\s\.\-\,\;]*)+`),                                                           // Chinese (Traditional) (zh-Hant)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Pruvincia|Region|a cità capitale|righjoni Capital|cità|Territory|di|l')\b[\s\.\-\,\;]*)+`),                         // Corsican (co)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Pokrajina|Regija|Glavni grad|Capital regija|Grad|Teritorija|od)\b[\s\.\-\,\;]*)+`),                                 // Croatian (hr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincie|Kraj|Hlavní město|Capital region|Město|Území|z)\b[\s\.\-\,\;]*)+`),                                       // Czech (cs)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provins|Område|Hovedstad|region hovedstaden|by|Territorium|af|det)\b[\s\.\-\,\;]*)+`),                              // Danish (da)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provinco|Regiono|Ĉefurbo|ĉefurbo regiono|Urbo|teritorio|el|la)\b[\s\.\-\,\;]*)+`),                                  // Esperanto (eo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(provints|piirkond|Pealinn|Capital piirkonnas|linn|territoorium|kohta)\b[\s\.\-\,\;]*)+`),                           // Estonian (et)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(maakunta|alue|Pääkaupunki|pääkaupunkiseutu|Kaupunki|Alue|of)\b[\s\.\-\,\;]*)+`),                                    // Finnish (fi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|Région|Capitale|région de la capitale|Ville|Territoire|de|les|en)\b[\s\.\-\,\;]*)+`),                      // French (fr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(provincia|rexión|capital|rexión da capital|cidade|territorio|de|o)\b[\s\.\-\,\;]*)+`),                              // Galician (gl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Mòr-roinn|Region|prìomh-bhaile|calpa roinn|City|Territory|de|a ')\b[\s\.\-\,\;]*)+`),                               // Gaelic (Scottish) (gd)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(პროვინციაში|რეგიონი|Დედაქალაქი|დედაქალაქის რეგიონი|ქალაქი|ტერიტორია|საქართველოს)\b[\s\.\-\,\;]*)+`),                // Georgian (ka)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provinz|Region|Hauptstadt|Hauptstadtregion|Stadt|Gebiet|von|das)\b[\s\.\-\,\;]*)+`),                                // German (de)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Επαρχία|Περιφέρεια|Πρωτεύουσα|περιοχή της πρωτεύουσας|Πόλη|Εδαφος|του|ο)\b[\s\.\-\,\;]*)+`),                        // Greek (el)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(પ્રાંત|પ્રદેશ|રાજધાની શહેર|કેપિટલ પ્રદેશ|શહેરનું|ટેરિટરી|ના|આ)\b[\s\.\-\,\;]*)+`),                                  // Gujarati (gu)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Pwovens|rejyon|Kapital vil|rejyon Kapital|City|Teritwa|nan|nan)\b[\s\.\-\,\;]*)+`),                                 // Haitian Creole (ht)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(lardin|yankin|babban birnin|babban birnin yankin|City|Territory|na|da)\b[\s\.\-\,\;]*)+`),                          // Hausa (ha)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(מָחוֹז|אזור|עיר בירה|אזור קפיטל|עִיר|שֶׁטַח|שֶׁל|ה)\b[\s\.\-\,\;]*)+`),                                             // Hebrew (he)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(प्रांत|क्षेत्र|राजधानी|राजधानी क्षेत्र|Faridabad|क्षेत्र|का)\b[\s\.\-\,\;]*)+`),                                    // Hindi (hi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Tartomány|Vidék|Főváros|Főváros régióban|Város|Terület|nak,-nek|az)\b[\s\.\-\,\;]*)+`),                             // Hungarian (hu)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|Region|Höfuðborg|Capital Region|Borg|Territory|af|sem)\b[\s\.\-\,\;]*)+`),                                 // Icelandic (is)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|Region|Capital obodo|Capital region|City|Territory|nke|na)\b[\s\.\-\,\;]*)+`),                             // Igbo (ig)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Cúige|Réigiún|Príomhchathair|réigiún caipitil|Cathair|Críoch|de|an)\b[\s\.\-\,\;]*)+`),                             // Irish (ga)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincia|Regione|Capitale|regione della capitale|Città|Territorio|di|il)\b[\s\.\-\,\;]*)+`),                       // Italian (it)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(州|領域|首都|首都圏|市|地域|の|インクルード)\b[\s\.\-\,\;]*)+`),                                                                      // Japanese (ja)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Langkawi|Region|kutha Capital|wilayah Capital|City|Territory|saka|ing)\b[\s\.\-\,\;]*)+`),                          // Javanese (jv)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ಪ್ರಾಂತ್ಯ|ಪ್ರದೇಶ|ರಾಜಧಾನಿ|ರಾಜಧಾನಿ ಪ್ರದೇಶ|ಸಿಟಿ|ಟೆರಿಟರಿ|ಆಫ್|ದಿ)\b[\s\.\-\,\;]*)+`),                                   // Kannada (kn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(провинция|аймақ|Астана|капитал облысы|қала|аумақ|туралы|The)\b[\s\.\-\,\;]*)+`),                                    // Kazakh (kk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ខេត្ត|តំបន់|រាជធានី|តំបន់រដ្ឋធានី|ទីក្រុង|ដែនដី|នៃ|នេះ)\b[\s\.\-\,\;]*)+`),                                         // Khmer (km)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Intara|Region|umurwa mukuru|Capital karere|Umugi|Territory|ya|mu)\b[\s\.\-\,\;]*)+`),                               // Kinyarwanda (Rwanda) (rw)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Облус|регион|борбор шаар|Капитал аймак|Сити|территория|боюнча|жана)\b[\s\.\-\,\;]*)+`),                             // Kyrgyz (ky)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(지방|부위|자본 도시|수도권|시티|영토|의|그만큼)\b[\s\.\-\,\;]*)+`),                                                                    // Korean (ko)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Herêm|Herêm|Paytext|herêmê Capital|Bajar|Herêm|ji|ew)\b[\s\.\-\,\;]*)+`),                                           // Kurdish (ku)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ແຂວງ|ພາກພື້ນ|ນະຄອນຫຼວງ|ເຂດນະຄອນຫຼວງ|ເມືອງ|ອານາເຂດຂອງ|ຂອງ|ໄດ້)\b[\s\.\-\,\;]*)+`),                                   // Lao (lo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|regio|Oppida prima|regionem capitis|Urbs|Territorium|autem|quod)\b[\s\.\-\,\;]*)+`),                       // Latin (la)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(province|Novads|Galvaspilsēta|Capital reģions|pilsēta|teritorija|no)\b[\s\.\-\,\;]*)+`),                            // Latvian (Lettish) (lv)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(provincija|regionas|Sostinė|Capital Region|miestas|teritorija|apie)\b[\s\.\-\,\;]*)+`),                             // Lithuanian (lt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provënz|Regioun|Haaptstad|Capital Regioun|City|Territoire|vun|der)\b[\s\.\-\,\;]*)+`),                              // Luxembourgish (lb)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(провинција|регионот|Главен град|капитал регион|Сити|територија|на|на)\b[\s\.\-\,\;]*)+`),                           // Macedonian (mk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(-tokony eran'ny fanjakana|Region|Renivohitra|Capital faritra|City|FARITANY|ny|ny)\b[\s\.\-\,\;]*)+`),               // Malagasy (mg)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(wilayah|rantau|Ibu negeri|Capital Region|City|wilayah|daripada|yang)\b[\s\.\-\,\;]*)+`),                            // Malay (ms)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(പ്രവിശ്യ|പ്രദേശം|തലസ്ഥാന നഗരം|തലസ്ഥാന|നഗരം|ടെറിട്ടറി|എന്ന|The)\b[\s\.\-\,\;]*)+`),                                  // Malayalam (ml)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(provinċja|reġjun|Belt kapitali|reġjun kapitali|belt|territorju|ta|il)\b[\s\.\-\,\;]*)+`),                           // Maltese (mt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Porowini|Region|Capital pa|Capital rohe|City|Territory|o|te)\b[\s\.\-\,\;]*)+`),                                    // Maori (mi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(प्रांत|प्रदेश|राजधानी|कॅपिटल प्रदेश|सिटी|प्रदेश|च्या|अगोदर निर्देश केलेल्या बाबीसंबंधी बोलताना)\b[\s\.\-\,\;]*)+`), // Marathi (mr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincie|Regiune|Capitala|regiunea de capital|Oraș|Teritoriu|de)\b[\s\.\-\,\;]*)+`),                               // Moldavian (mo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(аймгийн|бүс нутаг|Нийслэл хот|Нийслэлийн бүс|хот|газар нутаг|нь|The)\b[\s\.\-\,\;]*)+`),                            // Mongolian (mn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(प्रान्त|क्षेत्र|राजधानी|राजधानी क्षेत्र|शहर|इलाका|को|को)\b[\s\.\-\,\;]*)+`),                                        // Nepali (ne)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provins|Region|Hovedstad|Capital region|By|Territory|av|de)\b[\s\.\-\,\;]*)+`),                                     // Norwegian (no)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provins|Region|Hovedstad|Capital region|By|Territory|av|de)\b[\s\.\-\,\;]*)+`),                                     // Norwegian bokmål (nb)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ପ୍ରେଦଶ|ଅଞ୍ଚଳ|ରାଜଧାନୀ ସହର|Capital ଅଞ୍ଚଳ|ସହର|ଭୂଭାଗ|ର|େଯମାେନ)\b[\s\.\-\,\;]*)+`),                                      // Oriya (or)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ولایت|سیمه|پلازمېنه ښار|پلازمیینه سیمه|ښار|خاوره|د|د)\b[\s\.\-\,\;]*)+`),                                           // Pashto, Pushto (ps)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(استان|منطقه|پایتخت|منطقه پایتخت|شهرستان|قلمرو|از)\b[\s\.\-\,\;]*)+`),                                               // Persian (Farsi) (fa)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Województwo|Region|Stolica|Region stołeczny|Miasto|Terytorium|z|Prowincja)\b[\s\.\-\,\;]*)+`),                      // Polish (pl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Província|Região|Capital|região da capital|Cidade|Território|de|a|do)\b[\s\.\-\,\;]*)+`),                           // Portuguese (pt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ਸੂਬੇ|ਖੇਤਰ|ਰਾਜਧਾਨੀ|ਰਾਜਧਾਨੀ ਖੇਤਰ|ਸਿਟੀ|ਟੈਰੀਟਰੀ|ਦੇ|ਇਹ)\b[\s\.\-\,\;]*)+`),                                              // Punjabi (Eastern) (pa)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincie|Regiune|Capitala|regiunea de capital|Oraș|Teritoriu|de|a)\b[\s\.\-\,\;]*)+`),                             // Romanian (ro)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Провинция|Область|Столица|Столичный регион|Город|территория|из)\b[\s\.\-\,\;]*)+`),                                 // Russian (ru)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(setate|vāega|laumua|Tupe Faavae itulagi|aʻai|atunuʻu|a|le)\b[\s\.\-\,\;]*)+`),                                      // Samoan (sm)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(провинција|регија|Главни град|Цапитал region|град|територија|од)\b[\s\.\-\,\;]*)+`),                                // Serbian (sr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|sebakeng|motse-moholo|motse-moholo oa sebaka|City|Territory|ya|ho)\b[\s\.\-\,\;]*)+`),                     // Sesotho (st)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(dunhu|nharaunda|Guta guru|Capital nharaunda|guta|nzvimbo|pamusoro|ari)\b[\s\.\-\,\;]*)+`),                          // Shona (sn)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(صوبي|علائقي|گادي جو شھر|گاديء واري علائقي|شهر|سڏجي ٿو|جي|جي)\b[\s\.\-\,\;]*)+`),                                    // Sindhi (sd)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(පළාත|කලාපයේ|අග නගරය|අගනුවර කලාපයේ|නගරය|භූමිය|වල|එම)\b[\s\.\-\,\;]*)+`),                                             // Sinhalese (si)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincie|kraj|Hlavné mesto|capital región|veľkomesto|územie|z)\b[\s\.\-\,\;]*)+`),                                 // Slovak (sk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provinca|regija|Glavno mesto|Capital regija|Kraj|ozemlje|za)\b[\s\.\-\,\;]*)+`),                                    // Slovenian (sl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Gobolka|gobolka|magaalada Capital|gobolka Capital|City|Territory|of|ah)\b[\s\.\-\,\;]*)+`),                         // Somali (so)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provincia|Región|Ciudad capital|región de la capital|Ciudad|Territorio|de|la)\b[\s\.\-\,\;]*)+`),                   // Spanish (es)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Propinsi|daerah|kota Capital|wewengkon ibukota|kota|wewengkon|ti|éta)\b[\s\.\-\,\;]*)+`),                           // Sundanese (su)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Mkoa|Mkoa|Mji mkuu|Capital kanda|City|Wilaya|ya|the)\b[\s\.\-\,\;]*)+`),                                            // Swahili (Kiswahili) (sw)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provins|Område|Huvudstad|kapital region|Stad|Territorium|av|de)\b[\s\.\-\,\;]*)+`),                                 // Swedish (sv)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(lalawigan|Rehiyon|capital city|Capital region|lungsod|lupain|ng|ang)\b[\s\.\-\,\;]*)+`),                            // Tagalog (tl)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(вилоят|вилоят|Пойтахт|минтақаи Пойтахт|шаҳр|қаламрави|аз|ба)\b[\s\.\-\,\;]*)+`),                                    // Tajik (tg)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(மாகாணம்|பகுதி|தலை நாகரம்|தலைநகர பிராந்தியம்|நகரம்|பிரதேசம்|இன்|தி)\b[\s\.\-\,\;]*)+`),                              // Tamil (ta)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(провинция|өлкә|Башкала|капитал районы|шәһәр|территория|һәм|бу)\b[\s\.\-\,\;]*)+`),                                  // Tatar (tt)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(ప్రావిన్స్|ప్రాంతం|రాజధాని నగరం|రాజధాని ప్రాంతం|నగరం|భూభాగం|ఆఫ్|ది)\b[\s\.\-\,\;]*)+`),                             // Telugu (te)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(จังหวัด|ภูมิภาค|เมืองหลวง|ภูมิภาคทุน|เมือง|อาณาเขต|ของ)\b[\s\.\-\,\;]*)+`),                                         // Thai (th)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Bölge|bölge|Başkent|Sermaye bölge|Kent|bölge|nın-nin)\b[\s\.\-\,\;]*)+`),                                           // Turkish (tr)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(welaýat|sebit|Paýtagt|Capital sebit|şäher|meýdany|we|mälimlik görkeziji artikl)\b[\s\.\-\,\;]*)+`),                 // Turkmen (tk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(өлкә|رايون|پايتەخت شەھەر|پايتەخت رايون|شەھەر|территория|نىڭ|ишлитилмәйду)\b[\s\.\-\,\;]*)+`),                       // Uyghur (ug)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(провінція|область|Столиця|столичний регіон|місто|територія|з)\b[\s\.\-\,\;]*)+`),                                   // Ukrainian (uk)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(صوبہ|علاقہ|دارالحکومت|کیپٹل خطے|شہر|علاقہ|کے)\b[\s\.\-\,\;]*)+`),                                                   // Urdu (ur)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(viloyat|mintaqa|Poytaxt shahar|Capital viloyati|shahar|hududi|ning|The)\b[\s\.\-\,\;]*)+`),                         // Uzbek (uz)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Tỉnh|Khu vực|Thủ đô|khu vực vốn|thành phố|lãnh thổ|của|các)\b[\s\.\-\,\;]*)+`),                                     // Vietnamese (vi)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Talaith|rhanbarth|Prifddinas|rhanbarth y Brifddinas|City|Tiriogaeth|o|y)\b[\s\.\-\,\;]*)+`),                        // Welsh (cy)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Provinsje|Regio|Haadstêd|Capital regio|Stêd|Gebiet|fan|de)\b[\s\.\-\,\;]*)+`),                                      // Western Frisian (fy)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(kwiPhondo|Region|Isixeko esikhulu|indawo Capital|City|Territory|of|i)\b[\s\.\-\,\;]*)+`),                           // Xhosa (xh)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Province|Region|Olú ìlú|Capital ekun|City|agbegbe|ti|awọn)\b[\s\.\-\,\;]*)+`),                                      // Yoruba (yo)
		regexp.MustCompile(`(?i)([\s\.\-\,\;]*\b(Isifundazwe|Isifunda|Inhloko-dolobha|Capital esifundeni|City|Territory|ka|le)\b[\s\.\-\,\;]*)+`),                   // Zulu (zu)
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
